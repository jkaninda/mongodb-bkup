package cmd

import (
	"github.com/jkaninda/mongodb-bkup/pkg"
	"github.com/spf13/cobra"
)

var BackupCmd = &cobra.Command{
	Use:   "backup ",
	Short: "Backup database operation",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			pkg.StartBackup(cmd)
		} else {
		}
	},
}

func init() {
	//Backup
	BackupCmd.PersistentFlags().StringP("mode", "m", "default", "Set execution mode. default or scheduled")
	BackupCmd.PersistentFlags().StringP("period", "", "0 1 * * *", "Set schedule period time")
	BackupCmd.PersistentFlags().BoolP("disable-compression", "", false, "Disable backup compression")

}
