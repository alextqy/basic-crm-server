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

// 系统日志
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

	// admin ===================================================================
	mux.HandleFunc("/admin/sign/in", httpHelper.Middleware(api.AdminSignIn))
	mux.HandleFunc("/admin/sign/out", httpHelper.Middleware(api.AdminSignOut))
	mux.HandleFunc("/admin/new", httpHelper.Middleware(api.AdminNew))
	mux.HandleFunc("/admin/list", httpHelper.Middleware(api.AdminList))
	mux.HandleFunc("/admin/all", httpHelper.Middleware(api.AdminAll))
	mux.HandleFunc("/admin/data", httpHelper.Middleware(api.AdminData))
	mux.HandleFunc("/admin/del", httpHelper.Middleware(api.AdminDel))
	mux.HandleFunc("/admin/status", httpHelper.Middleware(api.AdminStatus))

	// company ===================================================================
	mux.HandleFunc("/company/new", httpHelper.Middleware(api.CompanyNew))
	mux.HandleFunc("/company/list", httpHelper.Middleware(api.CompanyList))
	mux.HandleFunc("/company/all", httpHelper.Middleware(api.CompanyAll))
	mux.HandleFunc("/company/data", httpHelper.Middleware(api.CompanyData))
	mux.HandleFunc("/company/del", httpHelper.Middleware(api.CompanyDel))

	// customer ===================================================================
	mux.HandleFunc("/customer/new", httpHelper.Middleware(api.CustomerNew))
	mux.HandleFunc("/customer/list", httpHelper.Middleware(api.CustomerList))
	mux.HandleFunc("/customer/all", httpHelper.Middleware(api.CustomerAll))
	mux.HandleFunc("/customer/data", httpHelper.Middleware(api.CustomerData))
	mux.HandleFunc("/customer/del", httpHelper.Middleware(api.CustomerDel))

	// manager ===================================================================
	mux.HandleFunc("/manager/new", httpHelper.Middleware(api.ManagerNew))
	mux.HandleFunc("/manager/list", httpHelper.Middleware(api.ManagerList))
	mux.HandleFunc("/manager/all", httpHelper.Middleware(api.ManagerAll))
	mux.HandleFunc("/manager/data", httpHelper.Middleware(api.ManagerData))
	mux.HandleFunc("/manager/del", httpHelper.Middleware(api.ManagerDel))
	mux.HandleFunc("/manager/status", httpHelper.Middleware(api.ManagerStatus))

	mux.HandleFunc("/manager/sign/in", httpHelper.Middleware(api.ManagerSignIn))
	mux.HandleFunc("/manager/sign/out", httpHelper.Middleware(api.ManagerSignOut))
	mux.HandleFunc("/manager/update", httpHelper.Middleware(api.ManagerUpdate))

	// Group ===================================================================
	mux.HandleFunc("/group/new", httpHelper.Middleware(api.GroupNew))
	mux.HandleFunc("/group/list", httpHelper.Middleware(api.GroupList))
	mux.HandleFunc("/group/all", httpHelper.Middleware(api.GroupAll))
	mux.HandleFunc("/group/data", httpHelper.Middleware(api.GroupData))
	mux.HandleFunc("/group/del", httpHelper.Middleware(api.GroupDel))

	// SalesPlan ===================================================================
	mux.HandleFunc("/sales/plan/new", httpHelper.Middleware(api.SalesPlanNew))
	mux.HandleFunc("/sales/plan/list", httpHelper.Middleware(api.SalesPlanList))
	mux.HandleFunc("/sales/plan/all", httpHelper.Middleware(api.SalesPlanAll))
	mux.HandleFunc("/sales/plan/data", httpHelper.Middleware(api.SalesPlanData))
	mux.HandleFunc("/sales/plan/del", httpHelper.Middleware(api.SalesPlanDel))

	// SalesTarget ===================================================================
	mux.HandleFunc("/sales/target/new", httpHelper.Middleware(api.SalesTargetNew))
	mux.HandleFunc("/sales/target/list", httpHelper.Middleware(api.SalesTargetList))
	mux.HandleFunc("/sales/target/all", httpHelper.Middleware(api.SalesTargetAll))
	mux.HandleFunc("/sales/target/data", httpHelper.Middleware(api.SalesTargetData))
	mux.HandleFunc("/sales/target/del", httpHelper.Middleware(api.SalesTargetDel))
}
