package setting

import (
	"time"
	"fmt"
	"gopkg.in/ini.v1"
)

const (
	TokenExpireDuration = time.Hour * 2
)

var MySecret = []byte("bajfjafjakjgkam")

var (
	AppMode  string
	HttpPort string

	Db       string
	DbHost   string
	DbPort   string
	DbUser   string
	DbSecret string
	DbName   string
)

func InitSetting() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误")
	}

	loadServer(file)
	loadDatabase(file)
}

func loadServer(file *ini.File) {
	AppMode=getValue(file,"server","AppMode","debug")
	HttpPort=getValue(file,"server","HttpPort",":3000")
}
func loadDatabase(file *ini.File) {
	Db=getValue(file,"database","Db","mysql")
	DbHost=getValue(file,"database","DbHost","localhost")
	DbPort=getValue(file,"database","DbPort","3306")
	DbSecret=getValue(file,"database","Db","")
	DbName=getValue(file,"database","Db","goHello")
	
}

func getValue(file *ini.File,section,key,def string)string{
	return file.Section(section).Key(key).MustString(def)
}