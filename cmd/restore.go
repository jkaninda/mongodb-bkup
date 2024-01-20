package cmd

import (
	"github.com/jkaninda/mongodb-bkup/pkg"
	"github.com/spf13/cobra"
)

var RestoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "Restore database operation",
	Run: func(cmd *cobra.Command, args []string) {
		pkg.StartRestore(cmd)

	},
}

func init() {
	//Restore
	RestoreCmd.PersistentFlags().StringP("file", "f", "", "File name of database")

}
