package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
    // Print prompt for folder/file path
    fmt.Println("Please provide a folder or file path:")

    // Read user input
    var path string
    fmt.Scanln(&path)

    // Check if path is provided
    if path == "" {
        fmt.Println("No folder or file path provided.")
        return
    }

    // Prompt for program name
    fmt.Print("Enter program name to open/run the file/folder: ")
    var program string
    fmt.Scanln(&program)

    // Execute the chosen program with the path as argument
    cmd := exec.Command(program, path)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    err := cmd.Run()
    if err != nil {
        fmt.Printf("Error executing command: %v\n", err)
    }
}
