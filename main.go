package main

import (
	api "basic-crm-server/API"
	mtd "basic-crm-server/MTD"
	"fmt"
	"log"
	"net/http"
	"time"
)

var cacheHelper = mtd.CacheHelper{}
var fileHelper = mtd.FileHelper{}
var httpHelper = mtd.HttpHelper{}
var sysHelper = mtd.SysHelper{}
var tcpHelper = mtd.TcpHelper{}
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
		udpHelper.Broadcast(port, ip+":"+fileHelper.CheckConf().TcpPort)
		time.Sleep(time.Second)
	}
}

// 系统日志f
func systemLog() {
	for {
		if !fileHelper.FileExist(fileHelper.LogDir()) {
			fileHelper.DirMake(fileHelper.LogDir())
		}
		time.Sleep(time.Second)
	}
}

func main() {
	ips := lineSearch()
	go loopBroadcast(ips[len(ips)-1], fileHelper.CheckConf().UdpPort)
	go systemLog()

	mux := http.NewServeMux()
	routes(mux)
	server := &http.Server{
		Addr:         ":" + fileHelper.CheckConf().TcpPort,
		WriteTimeout: time.Second * 5, //设置写超时
		ReadTimeout:  time.Second * 5, //设置读超时
		Handler:      mux,
	}
	log.Println("Http server on port:" + fileHelper.CheckConf().TcpPort)
	log.Fatal(server.ListenAndServe())
}

func routes(mux *http.ServeMux) {
	var httpHelper = mtd.HttpHelper{}
	mux.HandleFunc("/test", httpHelper.Middleware(api.Test))
	mux.HandleFunc("/admin/sign/in", httpHelper.Middleware(api.AdminSignIn))
	mux.HandleFunc("/admin/sign/out", httpHelper.Middleware(api.AdminSignOut))
	mux.HandleFunc("/admin/update", httpHelper.Middleware(api.AdminUpdate))
	mux.HandleFunc("/admin/list", httpHelper.Middleware(api.AdminList))
}
