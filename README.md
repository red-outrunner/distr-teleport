# distr-teleport
Package Installation Setup Scripts

This project provides setup scripts to install essential packages for various Linux distributions, including Void Linux, RHEL, and others. These scripts ensure a streamlined installation process with automatic package checks and version reporting.

Requirements

Root privileges (sudo access)

Network connectivity for package installation

Setup Instructions

Void Linux Setup (setup-void)

Ensure the script is executable:

chmod +x ./setup-void

Run the script:

./setup-void

This script uses xbps-install to update the system and install the required packages. The status report will indicate the installation status and version of each package.

RHEL Setup (setup-rhel)

Ensure the script is executable:

chmod +x ./setup-rhel

Run the script:

./setup-rhel

This script prioritizes the dnf package manager for RHEL systems, with a fallback to yum if dnf is unavailable. It updates the system and installs the required packages, providing a status report for each.

Status Reporting

Each setup script checks the current installation status of each package and outputs:

Installed: Displays the package name and version.

Not Installed: The script attempts to install the package and reports success or failure.

Packages Installed

flatpak

vim

python

go

dolphin

kate

bluefish

rawtherapee

konsole

vlc

clang

gcc

Error Handling

If any package fails to install, the error will be reported in the console output.

If you wanna tweak it and have other packages edit the go file and build
go build -o setup-rhel/setup-void/... ./main.go
viola

Notes

Ensure your package repositories are properly configured for Void Linux and RHEL systems.

Run the scripts with care to avoid package conflicts.

