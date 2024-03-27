package main

import (
	lib "basic-crm-server/LIB"
	"fmt"
	"log"
	"net/http"
	"time"
)

// 线路检索
func lineSearch() []string {
	fmt.Println("Local IP address:")
	_, e, ips := lib.LocalIP()
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
		lib.Broadcast(port, ip+":"+lib.CheckConf().TcpPort)
		time.Sleep(time.Second)
	}
}

// 系统日志
func systemLog() {
	for {
		if !lib.FileExist(lib.LogDir()) {
			lib.DirMake(lib.LogDir())
		}
		time.Sleep(time.Second)
	}
}

func main() {
	ips := lineSearch()
	go loopBroadcast(ips[len(ips)-1], lib.CheckConf().UdpPort)
	go systemLog()

	mux := http.NewServeMux()
	routes(mux)
	server := &http.Server{
		Addr:         ":" + lib.CheckConf().TcpPort,
		WriteTimeout: time.Second * 5, //设置写超时
		ReadTimeout:  time.Second * 5, //设置读超时
		Handler:      mux,
	}
	log.Println("Http server on port:" + lib.CheckConf().TcpPort)
	log.Fatal(server.ListenAndServe())
}

func routes(mux *http.ServeMux) {}
