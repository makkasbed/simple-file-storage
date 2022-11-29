package cmd

import "github.com/spf13/cobra"

var delCmd = &cobra.Command{
	Use:   "del",
	Short: "Delete a bucket",
	Long:  "This command can be used to delete a bucket",
	Run:   deleteBucket,
}

func deleteBucket(cmd *cobra.Command, args []string) {

}

func init() {
	RootCmd.AddCommand(delCmd)
}
