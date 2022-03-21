package cmd

import (
	"fmt"

	"github.com/ThorstenHans/mask/pkg/mask"
	"github.com/spf13/cobra"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:     "set",
	Short:   "Modifies the mask char",
	Example: `mask char set #`,
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		newChar := args[0]
		if len(newChar) != 1 {
			fmt.Println("Consider using a single replacement char. Otherwise formatted outputs could be messed up")
		}
		m := mask.LoadMasks()
		m.MaskChar = newChar
		m.Save()
	},
}

func init() {
	charCmd.AddCommand(setCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
