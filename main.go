package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	// 日志文件
	fLog, err := os.Create("svnlog.log")
	if err != nil {
		log.Fatal("create log file failed.")
	}
	defer fLog.Close()
	log.SetOutput(fLog)

	// 读配置
	ptrConfig := NewConfig()
	if err := ptrConfig.Load(); err != nil {
		log.Fatalf("读取配置失败:%v", err)
	}

	// 执行svn命令
	cmd := exec.Command("svn", "log", "-r", ptrConfig.GetVersion(), ptrConfig.SvnPath)
	out, err := cmd.Output()
	if err != nil {
		log.Fatalf("执行命令失败:%v", err)
	}
	svnData := GBKToUTF8(string(out))

	// 命令格式化
	sliSvnLog, err := Format(svnData)
	if err != nil {
		log.Fatalf("解析日志失败:%v", err)
	}

	// 输出文件
	//if err = OutputCSV(ptrConfig.Output, sliSvnLog); err != nil {
	//	log.Fatalf("输出日志失败:%v", err)
	//}
	if err = OutputExcel(ptrConfig.Output, sliSvnLog); err != nil {
		log.Fatalf("输出日志失败:%v", err)
	}

	log.Println("执行命令成功,输出文件:", ptrConfig.Output)
}
