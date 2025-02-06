package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// List of packages to install
	packages := []string{
		"flatpak",
		"vim",
		"python",
		"go",
		"dolphin",
		"kate",
		"bluefish",
		"rawtherapee",
		"konsole",
		"vlc",
		"clang",
		"gcc",
	}

	fmt.Println("Updating package database and system...")
	var updateCmd *exec.Cmd

	if _, err := exec.LookPath("dnf"); err == nil {
		updateCmd = exec.Command("sudo", "dnf", "update", "-y")
	} else {
		fmt.Println("dnf not found, trying yum...")
		updateCmd = exec.Command("sudo", "yum", "update", "-y")
	}

	updateCmd.Stdout = os.Stdout
	updateCmd.Stderr = os.Stderr

	if err := updateCmd.Run(); err != nil {
		fmt.Printf("Error updating system: %v\n", err)
		return
	}

	fmt.Println("\nStatus Report:")

	// Check and install each package
	for _, pkg := range packages {
		fmt.Printf("Checking status of %s...\n", pkg)
		var checkCmd *exec.Cmd
		if _, err := exec.LookPath("dnf"); err == nil {
			checkCmd = exec.Command("dnf", "list", "installed", pkg)
		} else {
			checkCmd = exec.Command("yum", "list", "installed", pkg)
		}

		output, err := checkCmd.Output()
		if err == nil && strings.Contains(string(output), pkg) {
			extractVersion(output, pkg)
			continue
		}

		fmt.Printf("%s not installed, attempting installation...\n", pkg)
		var installCmd *exec.Cmd
		if _, err := exec.LookPath("dnf"); err == nil {
			installCmd = exec.Command("sudo", "dnf", "install", "-y", pkg)
		} else {
			installCmd = exec.Command("sudo", "yum", "install", "-y", pkg)
		}

		installCmd.Stdout = os.Stdout
		installCmd.Stderr = os.Stderr

		if err := installCmd.Run(); err != nil {
			fmt.Printf("Error installing %s: %v\n", pkg, err)
		} else {
			fmt.Printf("%s installed successfully.\n", pkg)
		}
	}
}

func extractVersion(output []byte, pkg string) {
	scanner := bufio.NewScanner(strings.NewReader(string(output)))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, pkg) {
			fields := strings.Fields(line)
			if len(fields) >= 2 {
				fmt.Printf("%s is installed, version: %s\n", fields[0], fields[1])
				return
			}
		}
	}
	fmt.Printf("%s installed but version information not found.\n", pkg)
}

