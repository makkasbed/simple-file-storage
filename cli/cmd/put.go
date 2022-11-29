package cmd

import "github.com/spf13/cobra"

var putCmd = &cobra.Command{
	Use:   "put",
	Short: "Upload a file to bucket",
	Long:  "This command can be used to list all buckets",
	Run:   putFile,
}

func putFile(cmd *cobra.Command, args []string) {

}

func init() {
	RootCmd.AddCommand(putCmd)
}
