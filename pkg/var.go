package pkg

const s3MountPath string = "/s3mnt"
const s3fsPasswdFile string = "/etc/passwd-s3fs"

var (
	storage            = "local"
	file               = ""
	s3Path             = "/mongodb-bkup"
	dbName             = ""
	dbHost             = ""
	dbPort             = "27017"
	dbPassword         = ""
	dbUserName         = ""
	executionMode      = "default"
	storagePath        = "/backup"
	disableCompression = false
)
