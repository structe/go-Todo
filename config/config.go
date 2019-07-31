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

	Sqlurl string
)

func InitConfig() {
	CooKieMaxAge = 60
	ServerPort = "8000"
	TokenSecret = "meet"
	MysqlUser = "root"
	MysqlPassword = "open"
	MysqlPort = "tcp(192.168.1.150:3500)"
	MysqlDatabase = "meeting"
	Sqlurl = "root:open@tcp(192.168.1.150:3500)/meeting?charset=utf8&parseTime=True&loc=Local"
	//Sqlurl = "root:open@tcp(192.168.1.150:3500)/meeting?charset=utf8&parseTime=True&loc=Local"
}
