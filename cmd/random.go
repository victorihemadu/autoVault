/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Get a random dad joke",
	Long: `This command fetches a random dad joke from the icanhazdadjoke api`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("random called")
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)
}

type Joke struct {
	ID string `json:"id"`
	Joke string `json:"joke"`
	Status int `json:"status"`
}
