package main

import (
	"log"

	"github.com/ShubhamVerma1811/x/cmd"
	"github.com/ShubhamVerma1811/x/database"
)

func main() {

	_, err := database.SetupDB()

	if err != nil {
		log.Fatal(err)
	}

	rootCmd := cmd.SetupCommands()

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}

}
