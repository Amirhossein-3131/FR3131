package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"sync"

	"summary/alert"
)

func main() {
	// Receive target domain from command line
	target := flag.String("t", "", "Target domain to scan")
	flag.Parse()

	if *target == "" {
		fmt.Println("Please specify a target with -t flag")
		os.Exit(1)
	}

	fmt.Println("[+] Starting reconnaissance tools...")
	alert.SendTelegramMessage(fmt.Sprintf("[FR] Starting reconnaissance on: %s", *target))

	// Run the reconnaissance tools
	Calltools(*target)

	// Combine and process results
	fmt.Println("[+] Combining and processing results...")
	alert.SendTelegramMessage("[FR] Combining and processing results")

	err := ExecuteShellCommands()
	if err != nil {
		fmt.Println("[-] Error in ExecuteShellCommands:", err)
		alert.SendTelegramMessage(fmt.Sprintf("[FR] Error in ExecuteShellCommands: %s", err))
		os.Exit(1)
	}

	fmt.Println("[+] Reconnaissance completed successfully!")
	alert.SendTelegramMessage("[FR] Reconnaissance completed successfully!")
}

func Calltools(target string) {
	var wg sync.WaitGroup

	runCommand := func(command string, outputFile string, args ...string) {
		defer wg.Done()

		fmt.Printf("[+] Running %s...\n", command)
		alert.SendTelegramMessage(fmt.Sprintf("[FR] Running %s on %s", command, target))

		file, err := os.Create(outputFile)
		if err != nil {
			fmt.Printf("[-] Error creating %s file: %s\n", outputFile, err)
			alert.SendTelegramMessage(fmt.Sprintf("[FR] Error creating %s file: %s", outputFile, err))
			return
		}
		defer file.Close()

		cmd := exec.Command(command, args...)
		cmd.Stdout = file
		cmd.Stderr = os.Stderr

		err = cmd.Run()
		if err != nil {
			fmt.Printf("[-] Error running %s: %s\n", command, err)
			alert.SendTelegramMessage(fmt.Sprintf("[FR] Error running %s: %s", command, err))
		} else {
			fmt.Printf("[+] %s completed successfully\n", command)
			alert.SendTelegramMessage(fmt.Sprintf("[FR] %s completed successfully", command))
		}
	}

	wg.Add(4) // Number of tools without gau

	go runCommand("assetfinder", "assetfinder-output.txt", "--subs-only", target)
	// gau removed
	go runCommand("subfinder", "subfinder-output.txt", "-d", target)
	go runCommand("python3", "sublist3r-output.txt", "/opt/Sublist3r/sublist3r.py", "-d", target)
	go runCommand("findomain", "findomain-output.txt", "--target", target, "--quiet", "--unique-output", "findomain-output.txt")

	wg.Wait()

	fmt.Println("[+] All tools executed, processing results...")
	alert.SendTelegramMessage("[FR] All tools executed, processing results")

	logResults := func(toolName, filename string) {
		cmd := exec.Command("bash", "-c", fmt.Sprintf("wc -l < %s", filename))
		output, err := cmd.Output()
		if err != nil {
			fmt.Printf("[-] Error counting lines for %s: %s\n", toolName, err)
			alert.SendTelegramMessage(fmt.Sprintf("[FR] Error counting lines for %s: %s", toolName, err))
			return
		}

		count := string(output)
		fmt.Printf("[+] %s found %s domains\n", toolName, count)
		alert.SendTelegramMessage(fmt.Sprintf("[FR] %s found %s domains", toolName, count))
	}

	logResults("assetfinder", "assetfinder-output.txt")
	// gau removed
	logResults("subfinder", "subfinder-output.txt")
	logResults("sublist3r", "sublist3r-output.txt")
	logResults("findomain", "findomain-output.txt")

	alert.SendTelegramMessage("[FR] Tools section completed")
}

func ExecuteShellCommands() error {
	commands := []struct {
		cmd         string
		description string
	}{
		{
			`cat sublist3r-output.txt | sed 's/\x1b\[[0-9;]*m//g' | sed '/^\[.*\]/d' | grep -E "(ir|com|org|net|io|edu|gov|uk|de|jp|cn)" > sublist3r-temp.txt`,
			"Processing sublist3r results",
		},
		{
			"sort assetfinder-output.txt findomain-output.txt subfinder-output.txt sublist3r-temp.txt | uniq > all-sub-finder.txt",
			"Combining all results",
		},
		{
			"rm -f assetfinder-output.txt subfinder-output.txt sublist3r-temp.txt sublist3r-output.txt findomain-output.txt",
			"Cleaning up temporary files",
		},
	}

	for _, c := range commands {
		fmt.Printf("[+] %s...\n", c.description)
		alert.SendTelegramMessage(fmt.Sprintf("[FR] %s", c.description))

		if err := execute(c.cmd); err != nil {
			fmt.Printf("[-] Error executing command '%s': %s\n", c.description, err)
			alert.SendTelegramMessage(fmt.Sprintf("[FR] Error executing command '%s': %s", c.description, err))
			return err
		}

		fmt.Printf("[+] %s completed\n", c.description)
		alert.SendTelegramMessage(fmt.Sprintf("[FR] %s completed", c.description))
	}

	cmd := exec.Command("bash", "-c", "wc -l < all-sub-finder.txt")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("[-] Error counting final results:", err)
		alert.SendTelegramMessage(fmt.Sprintf("[FR] Error counting final results: %s", err))
	} else {
		count := string(output)
		fmt.Printf("[+] Final results: %s unique domains found\n", count)
		alert.SendTelegramMessage(fmt.Sprintf("[FR] Final results: %s unique domains found", count))
	}

	return nil
}

func execute(commandStr string) error {
	cmd := exec.Command("bash", "-c", commandStr)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
