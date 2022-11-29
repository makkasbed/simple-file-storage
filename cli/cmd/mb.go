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

var (
	name        string
	region      string
	description string
	bucket_type string
	versioning  string
)

func makeBucket(cmd *cobra.Command, args []string) {
	buckets := []core.Bucket{}
	for _, x := range args {
		buckets = append(buckets, core.Bucket{Name: x})
	}
	fmt.Println(buckets)
}

func init() {
	RootCmd.AddCommand(mbCmd)

	mbCmd.Flags().StringVarP(&name, "name", "n", "sfs bucket", "Name(e.g: sfs bucket)")
	mbCmd.Flags().StringVarP(&region, "region", "r", "af-west", "Region(e.g af-west")
	mbCmd.Flags().StringVarP(&description, "description", "d", "", "Bucket description")
	mbCmd.Flags().StringVarP(&bucket_type, "bucket_type", "b", "", "Bucket Type")
	mbCmd.Flags().StringVarP(&versioning, "version", "v", "", "Bucket Versioning")
}
