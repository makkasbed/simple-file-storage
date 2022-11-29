package cmd

import "github.com/spf13/cobra"

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove a file from a bucket",
	Long:  "This is used to remove a file from a bucket",
	Run:   rmFile,
}

func rmFile(cmd *cobra.Command, args []string) {

}

func init() {
	RootCmd.AddCommand(rmCmd)
}
