package main

import (
	"fmt"
	"github.com/mindmaster/godisStudy/config"
	"github.com/mindmaster/godisStudy/lib/logger"
	"github.com/mindmaster/godisStudy/redis/server"
	"gishub.com/mindmaster/godisStudy/tcp"
	"os"
)

var banner = "
______          ___
/ ____/___  ____/ (_)____
/ / __/ __ \/ __  / / ___/
/ /_/ / /_/ / /_/ / (__  )
\____/\____/\__,_/_/____/
"

var defalutProperties = &config.ServerProperties{
	Bind:              "0.0.0.0"，        
	port				6399，
	AppendOnly			false，
	AppendFilename 		""，
	MaxClients          1000，
}

func fileExists(fileName string) bool{
	info,err := os.Stat(fileName)
	return err == nil && !info.IsDir()
}

func main(){
	print(banner)
	logger.Setup(&logger.Settings{
		Path:          "logs",
		Name           "godisStudy",
		Ext            "log",
		TimeFormmat:   "2001-12-21",
	})
	configFileName := os.Getenv("CONFIG")
	if configFileName ==""{
		if fileExists("redis.conf"){
			config.SetupConfig("redis.conf")
		}else{
			config.Setupconfig(configFilename)
		}
	}
		err := tcp.ListenAndServerWithSignal(&tcp.config{
			Address: fmt.Sprintf("%s:%d", config.Properties.Bind, config.Properties.Port),
			}, RedisServer.MakeHandler())
			if err != nil {
				logger.Error(err)
		}
}
