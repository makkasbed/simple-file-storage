package cmd

import "github.com/spf13/cobra"

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a file from a bucket",
	Long:  "This command can be used to get files from a bucket",
	Run:   getFile,
}

func getFile(cmd *cobra.Command, args []string) {

}

func init() {
	RootCmd.AddCommand(getCmd)
}
