package utils

const RestoreExample = "mysql-bkup restore --dbname database --file \n" +
	"bkup restore --dbname database --storage s3 --path /custom-path --file db_20231219_022941.sql.gz"
const BackupExample = "mongodb-bkup backup --dbname database --disable-compression\n" +
	"mongodb-bkup backup --dbname database --storage s3 --path /custom-path --disable-compression"

const MainExample = "mongodb-bkup backup --dbname database --disable-compression\n" +
	"mongodb-bkup backup --dbname database --storage s3 --path /custom-path\n" +
	"mongodb-bkup restore --dbname database --file "
