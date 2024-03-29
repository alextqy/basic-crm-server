package main

import (
	api "basic-crm-server/API"
	mtd "basic-crm-server/MTD"
	"fmt"
	"log"
	"net/http"
	"time"
)

var fileHelper = mtd.FileHelper{}
var sysHelper = mtd.SysHelper{}
var udpHelper = mtd.UdpHelper{}

// 线路检索
func lineSearch() []string {
	fmt.Println("Local IP address:")
	_, e, ips := sysHelper.LocalIP()
	if e != "" {
		panic(e)
	}
	for i := 0; i < len(ips); i++ {
		fmt.Println(ips[i])
	}
	return ips
}

// 开启内网广播
func loopBroadcast(ip string, port string) {
	for {
		udpHelper.Broadcast(port, ip+":"+sysHelper.CheckConf().TcpPort)
		time.Sleep(time.Second)
	}
}

// 系统日志
func systemLog() {
	for {
		if !fileHelper.FileExist(sysHelper.LogDir()) {
			fileHelper.DirMake(sysHelper.LogDir())
		}
		time.Sleep(time.Second)
	}
}

func main() {
	ips := lineSearch()
	go loopBroadcast(ips[len(ips)-1], sysHelper.CheckConf().UdpPort)
	go systemLog()

	mux := http.NewServeMux()
	routes(mux)
	server := &http.Server{
		Addr:         ":" + sysHelper.CheckConf().TcpPort,
		WriteTimeout: time.Second * 5, //设置写超时
		ReadTimeout:  time.Second * 5, //设置读超时
		Handler:      mux,
	}
	log.Println("Http server on port:" + sysHelper.CheckConf().TcpPort)
	log.Fatal(server.ListenAndServe())
}

func routes(mux *http.ServeMux) {
	mux.HandleFunc("/sign/in", api.SignIn)
}
