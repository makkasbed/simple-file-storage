package cmd

import "github.com/spf13/cobra"

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "Lists all buckets",
	Long:  "This command can be used to list all buckets",
	Run:   listBuckets,
}

func listBuckets(cmd *cobra.Command, args []string) {

}

func init() {
	RootCmd.AddCommand(lsCmd)
}
