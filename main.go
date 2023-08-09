package main

import "cryptkit/commands"

func main() {
	rootCmd := commands.RootCmd
	rootCmd.AddCommand(commands.VersionCmd,
		commands.NewEncryptCmd(),
		commands.NewDecryptCmd(),
	)
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
