package cmd

import (
	"fmt"

	"github.com/ThorstenHans/mask/pkg/mask"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:     "get",
	Short:   "Prints the current char used for masking",
	Aliases: []string{"show"},
	Run: func(cmd *cobra.Command, args []string) {
		m := mask.LoadMasks()
		fmt.Printf("Current mask chart is %s\n", m.MaskChar)
	},
}

func init() {
	charCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
