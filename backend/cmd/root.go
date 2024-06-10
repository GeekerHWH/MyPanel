package cmd

import (
	"os/user"

	"github.com/spf13/cobra"
)

func init() {}

var RootCmd = &cobra.Command{
	Use:   "1panel",
	Short: "1Panel ，一款现代化的 Linux 面板",
	RunE: func(cmd *cobra.Command, args []string) error {
		// server.Start()
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
