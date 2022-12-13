/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/helicarrierstudio/untitled-evm/evm"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "start the evm and run the code",
	Long: `run recieves a bytecode and executes it in the EVM.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run called")
		// convert strings to bytes
		var code []byte
		for _, arg := range args {
			code = append(code, []byte(arg)...)
		}
		stack, result, status := evm.Run(code)
		fmt.Println("stack:", stack)
		fmt.Println("result:", result)
		fmt.Println("status:", status)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
