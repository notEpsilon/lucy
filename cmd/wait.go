/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/notEpsilon/lucy/pkg/constants"
	"github.com/notEpsilon/lucy/pkg/server"
	"github.com/spf13/cobra"
)

// waitCmd represents the wait command
var waitCmd = &cobra.Command{
	Use:   "wait",
	Short: "Starts the server to wait for incoming connections and receive files",
	Run: func(cmd *cobra.Command, args []string) {
		bpi, err := cmd.Flags().GetInt32("bytesPerIteration")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		outputPath, err := cmd.Flags().GetString("output-file")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		server.Start(outputPath, int(bpi))
	},
}

func init() {
	rootCmd.AddCommand(waitCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// waitCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	waitCmd.Flags().Int32P("bytesPerIteration", "b", 10*constants.KiB, "How many bytes to receive per iteration, must match the sender (default 10 KB)")
	waitCmd.Flags().StringP("output-file", "o", "lucy_output_file.zip", "Output file name including file extension")
}
