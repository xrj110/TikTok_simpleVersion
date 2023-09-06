package tools

import (
	"fmt"
	"github.com/spf13/viper"
	// 其他导入语句...
)

func ReadConfig() {
	viper.SetConfigFile("tools/config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error reading config file:", err)
		return
	}

	// 从配置文件中获取值
	model := viper.GetInt("model")

	mysqlServerName := viper.GetString("mysql_server_name")
	mysqlPort := viper.GetString("mysql_port")
	redisServerName := viper.GetString("redis_server_name")
	redisPort := viper.GetString("redis_port")
	serverIP := viper.GetString("server_ip")
	serverPort := viper.GetString("server_port")
	if model == 1 {
		DocRedisInit(redisServerName, redisPort)
		DocSqlInit(mysqlServerName, mysqlPort)
		fmt.Println("docker model")
	} else {
		RedisInit()
		SqlInit()
		fmt.Println("local model")
	}
	CreateTable()
	ServerSetting(serverIP, serverPort)

}
