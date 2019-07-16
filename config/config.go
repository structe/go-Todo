package config

var (
	CooKieMaxAge int
	ServerPort   string
	TokenSecret  string
	// mysql
	MysqlUser     string
	MysqlPassword string
	MysqlDatabase string
	MysqlPort     string
)

func InitConfig() {
	CooKieMaxAge = 60
	ServerPort = "8000"
	TokenSecret = "meet"
	MysqlUser = "root"
	MysqlPassword = ""
	MysqlDatabase = ""
	MysqlPort = ""
}
