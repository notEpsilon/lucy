/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/notEpsilon/lucy/pkg/client"
	"github.com/notEpsilon/lucy/pkg/constants"
	"github.com/spf13/cobra"
)

// sendCmd represents the send command
var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "Send a file to the receiver device",
	Run: func(cmd *cobra.Command, args []string) {
		bpi, err := cmd.Flags().GetInt32("bytesPerIteration")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		file, err := cmd.Flags().GetString("file")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		host, err := cmd.Flags().GetString("host")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		client.Send(file, int(bpi), host)
	},
}

func init() {
	rootCmd.AddCommand(sendCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sendCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sendCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	sendCmd.Flags().Int32P("bytesPerIteration", "b", 10*constants.KiB, "How many bytes to send per iteration (default 10 KB)")
	sendCmd.Flags().StringP("file", "f", "", "The file name to send including extension")
	sendCmd.Flags().String("host", "", "The LAN host address to send the file to")
}
