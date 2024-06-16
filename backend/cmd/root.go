package cmd

import (
	"os/user"

	"github.com/spf13/cobra"
	"web-appstore.io/server"
)

func init() {}

var RootCmd = &cobra.Command{
	Use:   "MyPanel",
	Short: "MyPanel, helps people manage their Linux system in a decent way",
	RunE: func(cmd *cobra.Command, args []string) error {
		server.Start()
		return nil
	},
}

func isRoot() bool {
	currentUser, err := user.Current()
	if err != nil {
		return false
	}
	return currentUser.Uid == "0"
}
