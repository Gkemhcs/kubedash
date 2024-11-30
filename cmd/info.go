package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Print additional information about the tool",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`
 _  ___   _ ___ ___   ___   _   ___ _  _ 
 | |/ / | | | _ ) __| |   \ /_\ / __| || |
 | ' <| |_| | _ \ _|  | |) / _ \\__ \ __ |
 |_|\_\\___/|___/___| |___/_/ \_\___/_||_|
                                                                                                                                                   
                                                                                                                 
		`)
		fmt.Println("This CLI tool is a demonstration of Go, Cobra, and tview.")
	},
}
