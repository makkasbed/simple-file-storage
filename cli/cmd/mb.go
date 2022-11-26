package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"logiclabent.com/sfs-cli/core"
)

var mbCmd = &cobra.Command{
	Use:   "mb",
	Short: "Make buckets with this bucket",
	Long:  "This command enables the creation of a bucket",
	Run:   makeBucket,
}

func makeBucket(cmd *cobra.Command, args []string) {
	buckets := []core.Bucket{}
	for _, x := range args {
		buckets = append(buckets, core.Bucket{Name: x})
	}
	fmt.Println(buckets)
}

func init() {
	RootCmd.AddCommand(mbCmd)
}
