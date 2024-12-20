package cmd

import (
	"fmt"
	"os"

	//   "github.com/rivo/tview"
	"github.com/spf13/cobra"
	// "github.com/Gkemhcs/kubedash/internal/ui"

	"github.com/Gkemhcs/kubedash/internal/ui"
)

var contextFlag string

var rootCmd = &cobra.Command{
	Use:   "kubedash",
	Short: "A CLI tool with tview and cobra integration dislatys",
	Run: func(cmd *cobra.Command, args []string) {

		kubeUI := ui.AppUI{}
		err := kubeUI.InitDashboard()
		if err != nil {
			os.Exit(1)
		}
		if err = kubeUI.AppConfig.App.SetRoot(kubeUI.AppConfig.Pages, true).Run(); err != nil {
			fmt.Fprintf(os.Stderr, "Error running application: %v\n", err)
			os.Exit(1)

		}

	},
}

// Execute executes the cli
func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error executing command: %v\n", err)
		os.Exit(1)
		return err
	}
	return nil
}

func init() {
	rootCmd.Flags().StringVarP(&contextFlag, "context", "c", "", "The context to use for connecting to Kubernetes")
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(infoCmd)
}
