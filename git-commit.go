package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// Terminal color codes
const (
	Bold  = "\033[1m"
	Red   = "\033[31m"
	Green = "\033[32m"
	Reset = "\033[0m"
)

// Print error message with red color and a cross symbol
func printError(msg string, output string) {
	fmt.Println()
	fmt.Println(Bold + Red + "✗" + Reset + " " + msg)
	fmt.Println("--------------------------------")
	if output != "" {
		fmt.Print(Bold + Red)
		fmt.Print(output)
		fmt.Print(Reset)
		fmt.Println()
	}
	fmt.Println("--------------------------------")
}

// Print success message with green color and a checkmark symbol
func printSuccess(msg string, output string) {
	fmt.Println()
	fmt.Println(Bold + Green + "✓" + Reset + " " + msg)
	fmt.Println("--------------------------------")
	if output != "" {
		fmt.Print(output)
		fmt.Println()
	}
	fmt.Println("--------------------------------")
}

// Check if we're inside a git repository
func mustBeGitRepo() {
	cmd := exec.Command(
		"git",
		"rev-parse",
		"--is-inside-work-tree",
	)
	if err := cmd.Run(); err != nil {
		printError("Not a git repository", "Please run this command inside a git repository.")
		os.Exit(1)
	}
}

// Prompt user to select release type
func selectReleaseType() string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Select Release Type:")
	fmt.Println("1) major")
	fmt.Println("2) minor")
	fmt.Println("3) patch")

	for {
		fmt.Print("Enter choice (1-3): ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			return "major"
		case "2":
			return "minor"
		case "3":
			return "patch"
		default:
			fmt.Println("Invalid choice. Please select 1, 2, or 3.")
		}
	}
}

// Main function to handle user input and perform git commit
func main() {
	// Ensure we're in a git repository
	mustBeGitRepo()

	reader := bufio.NewReader(os.Stdin)

	// Task ID (numbers only)
	fmt.Print("Enter Task ID (number only): ")
	taskIDInput, _ := reader.ReadString('\n')
	taskIDInput = strings.TrimSpace(taskIDInput)

	taskID, err := strconv.Atoi(taskIDInput)
	if err != nil {
		printError("Invalid Task ID", "Only numeric Task IDs are allowed.")
		os.Exit(1)
	}

	// Commit message
	fmt.Print("Enter Commit Message: ")
	commitMsg, _ := reader.ReadString('\n')
	commitMsg = strings.TrimSpace(commitMsg)

	if commitMsg == "" {
		printError("Invalid Commit Message", "Commit message cannot be empty.")
		os.Exit(1)
	}

	// Release type
	releaseType := selectReleaseType()

	finalMessage := fmt.Sprintf(
		"%s - %s refs #%d",
		commitMsg,
		releaseType,
		taskID,
	)

	// git add .
	// addCmd := exec.Command("git", "add", ".")
	// if output, err := addCmd.CombinedOutput(); err != nil {
	// 	printError("git add failed", strings.TrimSpace(string(output)))
	// 	os.Exit(1)
	// }

	// git commit
	commitCmd := exec.Command("git", "commit", "-m", finalMessage)
	output, err := commitCmd.CombinedOutput()
	outputStr := strings.TrimSpace(string(output))

	if err != nil {
		printError("Git commit failed", outputStr)
		os.Exit(1)
	}

	// Success
	printSuccess("Git commit successful", outputStr)
}
