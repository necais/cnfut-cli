/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/oktayalizada/cnfut/middleware"
	"github.com/oktayalizada/cnfut/pkg"

	"github.com/spf13/cobra"
)

// copyCmd represents the copy command
var copyCmd = &cobra.Command{
	Use:   "copy",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		srcType, _ := cmd.Flags().GetString("srct")
		destType, _ := cmd.Flags().GetString("destt")
		src, _ := cmd.Flags().GetString("src")
		dest, _ := cmd.Flags().GetString("dest")
		concurrent, _ := cmd.Flags().GetBool("ccrt")

		srcDest := pkg.SourceDestination{
			SourceType:      srcType,
			DestinationType: destType,
			Destination:     dest,
			Source:          src,
			Concurrent:      concurrent,
		}
		fmt.Printf("Middleware checking %+v\n", srcDest)
		if middleware.CheckInput(srcDest) {
			pkg.Copy(srcDest)
		} else {
			fmt.Println("bad request")
		}
		// Storing task in backend calling my-todos REST API
	},
}

func init() {
	rootCmd.AddCommand(copyCmd)
	copyCmd.Flags().StringP("srct", "", "", "specify source type")
	copyCmd.Flags().StringP("src", "", "", "specify source")
	copyCmd.Flags().StringP("destt", "", "", "specify destination type")
	copyCmd.Flags().StringP("dest", "", "", "specify destination")
	copyCmd.Flags().BoolP("ccrt", "c", false, "Should be executed concurrently")
	err := copyCmd.MarkFlagRequired("srct")
	if err != nil {
		return
	}
	err = copyCmd.MarkFlagRequired("src")
	if err != nil {
		return
	}
	err = copyCmd.MarkFlagRequired("destt")
	if err != nil {
		return

	}
	err = copyCmd.MarkFlagRequired("dest")
	if err != nil {
		return
	}

}
