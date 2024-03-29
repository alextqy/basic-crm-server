package main

import (
	api "basic-crm-server/API"
	mtd "basic-crm-server/MTD"
	"fmt"
	"log"
	"net/http"
	"time"
)

// 线路检索
func lineSearch() []string {
	fmt.Println("Local IP address:")
	_, e, ips := mtd.LocalIP()
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
		mtd.Broadcast(port, ip+":"+mtd.CheckConf().TcpPort)
		time.Sleep(time.Second)
	}
}

// 系统日志
func systemLog() {
	for {
		if !mtd.FileExist(mtd.LogDir()) {
			mtd.DirMake(mtd.LogDir())
		}
		time.Sleep(time.Second)
	}
}

func main() {
	ips := lineSearch()
	go loopBroadcast(ips[len(ips)-1], mtd.CheckConf().UdpPort)
	go systemLog()

	mux := http.NewServeMux()
	routes(mux)
	server := &http.Server{
		Addr:         ":" + mtd.CheckConf().TcpPort,
		WriteTimeout: time.Second * 5, //设置写超时
		ReadTimeout:  time.Second * 5, //设置读超时
		Handler:      mux,
	}
	log.Println("Http server on port:" + mtd.CheckConf().TcpPort)
	log.Fatal(server.ListenAndServe())
}

func routes(mux *http.ServeMux) {
	mux.HandleFunc("/sign/in", api.SignIn)
}
