package commands

import "github.com/spf13/cobra"

var (
	cfgFile string
)

func init() {
	//RootCmd.PersistentFlags().StringVar(&cfgFile, "f", "configs/config.toml", "config file (default is ./configs/config.toml)")
}

var RootCmd = &cobra.Command{
	Use:   "cryptkit",
	Short: "cryptkit : kit of crypt",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) (err error) {
		if cmd.Name() == VersionCmd.Name() {
			return nil
		}

		return nil
	},
}
