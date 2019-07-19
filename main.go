package main

import (
	"flag"
	"fmt"
	"github.com/go-ini/ini"
	U "github.com/neurons-platform/gotools/utils"
	C "neurons_robot/client"
	CF "neurons_robot/conf"
	"os"
	"time"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "f", "", "配置文件")
	flag.Parse()
	if configFile == "" {
		flag.Usage()
		os.Exit(2)
	}
}

func init_config() {
	// start
	cfg, err := ini.InsensitiveLoad(configFile)
	U.Exit(err, "解析配置文件出错")

	clientConf, err := cfg.GetSection("client")
	U.Exit(err, "没有找到client配置")
	clientNum, err := clientConf.GetKey("num")
	U.Exit(err, "没有找到client num")
	CF.ApNumber = U.Str2Int(clientNum.Value())

	groupConf, err := cfg.GetSection("group")

	U.Exit(err, "没有找到group配置")
	monitorGroup, err := groupConf.GetKey("monitor_group")
	U.Exit(err, "没有找到监控群号")
	CF.MonitorGroup = monitorGroup.Value()

	crontabConf, err := cfg.GetSection("crontab")
	U.Exit(err, "没有找到crontab配置")
	enable, err := crontabConf.GetKey("enable")
	CF.EnableCrontab, err = enable.Bool()
	U.Exit(err, "没有找到 enable")
}

func start_http_client() {
	go func() {
		httpClient := &C.HttpClient{}
		go httpClient.Init()
		time.Sleep(5 * time.Second)
		fmt.Println("start http recv")
		go httpClient.Recv()
	}()
}

func main() {
	f, err := os.OpenFile("./log/log.txt", os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	U.InitLog(f, f, f)

	init_config()
	start_http_client()

	select {}
}
