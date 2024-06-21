package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	fmt.Print("Enter the path you want to create the Project In: ")
	var path string
	fmt.Scanln(&path)

	// Change directory to the specified path
	err := os.Chdir(path)
	if err != nil {
		fmt.Println("could not change directory:", err)
		return
	}

	var reactType string
	for {
		fmt.Print("Would you like to start a Remix or Vite app? (Enter Vite or Remix): ")
		fmt.Scanln(&reactType)

		if reactType == "Remix" {
			// Create Remix App (example command)
			cmd := exec.Command("ls", "./")
			out, err := cmd.CombinedOutput()
			if err != nil {
				fmt.Println("could not run command:", err)
			} else {
				fmt.Println("Output:", string(out))
			}
			break

		} else if reactType == "Vite" {
			// Create Vite App
			fmt.Println("You Picked Vite")

			var projectName string
			fmt.Print("Enter a project name: ")
			fmt.Scanln(&projectName)
			fmt.Printf("You picked %s\n", projectName)

			cmd := exec.Command("npm", "create", "vite@latest", projectName, "--", "--template", "react-swc-ts")
			out, err := cmd.CombinedOutput()
			if err != nil {
				fmt.Println("could not run command:", err)
			} else {
				fmt.Println("Output:", string(out))
			}
			break

		} else {
			fmt.Println("Incorrect Input, please enter 'Vite' or 'Remix'.")
			main()
		}
	}
}
