package cmd

import (
	"fmt"
	"os"

	"github.com/ThorstenHans/mask/pkg/mask"
	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var Version = "v0.0.1"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "mask",
	Version: Version,
	Short:   "mask - A CLI to mask strings in STDIN",
	Long: `With mask you can replace sensitive text fragments before being printed to STDOUT.
    
You can add as many masks as you want and specify the char used for masking`,
	Example: `# Add "bar" to replace it
mask add bar

# pipe output through mask
echo "foo bar baz" | mask
> foo *** baz
    `,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	Run: func(cmd *cobra.Command, args []string) {
		info, err := os.Stdin.Stat()
		if err != nil {
			fmt.Printf("Error while accessing stdin: %s", err)
			os.Exit(1)
		}
		if (info.Mode() & os.ModeCharDevice) == 0 {
			m := mask.LoadMasks()

			mw := mask.NewMaskedWriter(m, os.Stdin, os.Stdout)
			mw.Write()
			os.Exit(0)
		}
		fmt.Println("mask works with data piped to the command e.g.: `ls -al | mask`")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// Find home directory.
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	// Search config in home directory with name ".mask" (without extension).
	viper.AddConfigPath(home)
	viper.SetConfigType("yaml")
	viper.SetConfigName(".mask")

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			fmt.Printf("Error processing config file %s\n", err)
		}
	}
}
