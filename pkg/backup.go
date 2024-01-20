// Package pkg /*
/*
Copyright © 2024 Jonas Kaninda  <jonaskaninda.gmail.com>
*/
package pkg

import (
	"fmt"
	"github.com/jkaninda/mongodb-bkup/utils"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

func StartBackup(cmd *cobra.Command) {
	_, _ = cmd.Flags().GetString("operation")

	//Set env
	utils.SetEnv("STORAGE_PATH", storagePath)
	utils.GetEnv(cmd, "dbname", "DB_NAME")
	utils.GetEnv(cmd, "port", "DB_PORT")
	utils.GetEnv(cmd, "period", "SCHEDULE_PERIOD")

	//Get flag value and set env
	s3Path = utils.GetEnv(cmd, "path", "S3_PATH")
	storage = utils.GetEnv(cmd, "storage", "STORAGE")
	file = utils.GetEnv(cmd, "file", "FILE_NAME")
	disableCompression, _ = cmd.Flags().GetBool("disable-compression")
	executionMode, _ = cmd.Flags().GetString("mode")

	if executionMode == "default" {
		if storage == "s3" {
			utils.Info("Backup database to s3 storage")
			s3Backup(disableCompression, s3Path)
		} else {
			utils.Info("Backup database to local storage")
			BackupDatabase(disableCompression)

		}
	} else if executionMode == "scheduled" {
		scheduledMode()
	} else {
		utils.Fatal("Error, unknown execution mode!")
	}

}

// Run in scheduled mode
func scheduledMode() {

	fmt.Println()
	fmt.Println("**********************************")
	fmt.Println("     Starting MongoDB Bkup...       ")
	fmt.Println("***********************************")
	utils.Info("Running in Scheduled mode")
	utils.Info("Log file in /var/log/mongodb-bkup.log")
	utils.Info("Execution period ", os.Getenv("SCHEDULE_PERIOD"))

	//Test database connexion
	utils.TestDatabaseConnection()

	utils.Info("Creating backup job...")
	CreateCrontabScript(disableCompression, storage)

	//Start Supervisor
	supervisordCmd := exec.Command("supervisord", "-c", "/etc/supervisor/supervisord.conf")
	if err := supervisordCmd.Run(); err != nil {
		utils.Fatalf("Error starting supervisord: %v\n", err)
	}
}

// BackupDatabase backup database
func BackupDatabase(disableCompression bool) {
	dbHost = os.Getenv("DB_HOST")
	//dbPassword := os.Getenv("DB_PASSWORD")
	//dbUserName := os.Getenv("DB_USERNAME")
	dbName = os.Getenv("DB_NAME")
	dbPort = os.Getenv("DB_PORT")
	storagePath = os.Getenv("STORAGE_PATH")
	//TODO: Backup logic

}

func s3Backup(disableCompression bool, s3Path string) {
	// Backup Database to S3 storage
	MountS3Storage(s3Path)
	BackupDatabase(disableCompression)
}
