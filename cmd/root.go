/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "jsonls",
	Short: "Directory listing in JSON format",
	Long:  `Provides a JSON representation of the directory listing.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		if help, _ := cmd.Flags().GetBool("help"); help {
			cmd.Help()
			return
		}
		ds, err := os.ReadDir(".")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
		}
		entries := make([]string, 0, len(ds))
		for _, d := range ds {
			if isdir, _ := cmd.Flags().GetBool("dirs"); isdir && !d.IsDir() {
				continue
			}
			entries = append(entries, d.Name())
		}
		b, _ := json.Marshal(entries)
		os.Stdout.Write(b)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.jsonls.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("help", "h", false, "Display help")
	rootCmd.Flags().BoolP("dirs", "d", false, "List directories only")
}
