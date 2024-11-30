package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
)

var version = "0.0.1"
var author = "Gkemhcs"

var versionCmd = &cobra.Command{
    Use:   "version",
    Short: "Print the version number and author",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Fprintln(cmd.OutOrStdout(), `
        _  ___   _ ___ ___   ___   _   ___ _  _ 
        | |/ / | | | _ ) __| |   \ /_\ / __| || |
        | ' <| |_| | _ \ _|  | |) / _ \\__ \ __ |
        |_|\_\\___/|___/___| |___/_/ \_\___/_||_|
                                                                                                                                 
               `)
               fmt.Fprintf(cmd.OutOrStdout(), "kubedash version %s by %s\n", version, author)
		   },
}
