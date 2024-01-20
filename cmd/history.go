package cmd

import (
	"github.com/jkaninda/mongodb-bkup/utils"
	"github.com/spf13/cobra"
)

var HistoryCmd = &cobra.Command{
	Use:   "history",
	Short: "Show the history of backup",
	Run: func(cmd *cobra.Command, args []string) {
		utils.ShowHistory()
	},
}
