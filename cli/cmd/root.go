package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "sfs",
	Short: "SFS is a simple file storage service similar to AWS s3",
	Long:  `SFS will help you to create buckets, manage buckets, upload files and delete files`,
}

func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
