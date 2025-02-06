package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// List of packages to install
	packages := []string{
		"flatpak",
		"vim",
		"micro",
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

	// Update package database and system
	updateCmd := exec.Command("sudo", "xbps-install", "-Syu")
	updateCmd.Stdout = os.Stdout
	updateCmd.Stderr = os.Stderr

	fmt.Println("Updating package database and system...")
	if err := updateCmd.Run(); err != nil {
		fmt.Printf("Error updating system: %v\n", err)
		return
	}

	// Install each package
	for _, pkg := range packages {
		fmt.Printf("Installing %s...\n", pkg)
		installCmd := exec.Command("sudo", "xbps-install", "-y", pkg)
		installCmd.Stdout = os.Stdout
		installCmd.Stderr = os.Stderr

		if err := installCmd.Run(); err != nil {
			fmt.Printf("Error installing %s: %v\n", pkg, err)
		} else {
			fmt.Printf("%s installed successfully.\n", pkg)
		}
	}
}

