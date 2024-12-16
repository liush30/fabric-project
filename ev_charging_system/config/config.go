package config

import (
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

var ChargeConfig Config

type Config struct {
	ConsoleMode bool   `yaml:"console_mode"` //控制台模式
	LogLevel    string `yaml:"logLevel"`     //日志级别
	ShowSql     bool   `yaml:"showsql"`      //是否显示sql语句
	NodeEnv     string `yaml:"nodeEnv"`
	DBInfo      Mysql  `yaml:"mysql"` //mysql配置文件
	WebInfo     *Web   `yaml:"web"`
	JWTInfo     *JWT   `yaml:"JWT"`
}

// mysql配置文件
type Mysql struct {
	Host   string `yaml:"host"`
	Port   string `yaml:"port"`
	User   string `yaml:"user"`
	Pwd    string `yaml:"passwd"`
	DbName string `yaml:"dbname"`
}

type Web struct {
	Port string `yaml:"port"`
}

type JWT struct {
	JwtSingKey     string `yaml:"jwtSingKey"`     //jwt 签名的秘钥
	ExpirationTime int64  `yaml:"expirationtime"` //过期时间
}

func LoadConfig(path string) {
	file, err := os.Open(path)
	if err != nil {
		panic("Open " + path + " file error,err :" + err.Error())
	}
	defer file.Close()

	filebyte, err := io.ReadAll(file)
	if err != nil {
		panic("Read file error,err : " + err.Error())
	}

	err = yaml.Unmarshal(filebyte, &ChargeConfig)
	if err != nil {
		panic("Unmarshal error,err :" + err.Error())
	}

}

func DatabaseStr() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		ChargeConfig.DBInfo.User, ChargeConfig.DBInfo.Pwd, ChargeConfig.DBInfo.Host, ChargeConfig.DBInfo.Port, ChargeConfig.DBInfo.DbName)
}

func IsConsoleMode() bool {
	return ChargeConfig.ConsoleMode
}
