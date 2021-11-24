package main

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

type SvnLog struct {
	version  string
	author   string
	datetime string
	comment  string
}

func NewSvnLog() *SvnLog {
	return &SvnLog{}
}

func (this *SvnLog) Load(data string) error {
	log.Println("[load]", data)
	l1 := strings.Index(data, "\r\n")
	if l1 < 0 {
		l1 = strings.Index(data, "\n")
	}

	// 有comment
	if l1 > 0 {
		this.comment = strings.TrimSpace(data[l1:])
		this.comment = strings.ReplaceAll(this.comment, "\r\n", " ")
		this.comment = strings.ReplaceAll(this.comment, "\n", " ")
		data = data[:l1]
	} else {
		this.comment = "无备注"
	}

	// 头数据
	sliData := strings.Split(data, "|")
	if len(sliData) != 4 {
		return errors.New(fmt.Sprintf("invalid format log:[%v]", data))
	}
	this.version = strings.TrimSpace(sliData[0])[1:]
	this.author = strings.TrimSpace(sliData[1])
	this.datetime = strings.TrimSpace(sliData[2])
	this.datetime = strings.ReplaceAll(this.datetime, ",", "")

	return nil
}

func (this *SvnLog) ToCsv() string {
	return fmt.Sprintf("%v,%v,%v,%v\r\n", this.comment, this.author, this.datetime, this.version)
}

func (this *SvnLog) Test() {
	log.Printf("%+v", *this)
}

func Format(svnData string) ([]*SvnLog, error) {
	//log.Println(svnData)
	f1 := strings.Index(svnData, "\n")
	if f1 < 0 {
		return nil, nil
	}
	sepator := svnData[:f1]
	l := len(sepator)

	sliSvnLog := make([]*SvnLog, 0)
	for {
		for f := strings.Index(svnData, sepator); f == 0; f = strings.Index(svnData, sepator) {
			svnData = svnData[l:]
		}
		t := strings.Index(svnData, sepator)
		if t < 0 {
			//处理最后一个数据不需要处理
			break
		} else {
			data := svnData[:t]
			data = strings.TrimSpace(data)
			data = strings.ReplaceAll(data, "\r\n\r\n", "\r\n")
			data = strings.ReplaceAll(data, "\n\n", "\n")
			ptrSvnLog := NewSvnLog()
			if err := ptrSvnLog.Load(data); err != nil {
				return sliSvnLog, err
			}
			//ptrSvnLog.Test()
			sliSvnLog = append(sliSvnLog, ptrSvnLog)
			//fmt.Println("loop===============")
			//fmt.Println(data)
			svnData = svnData[t:]
		}
	}

	return sliSvnLog, nil
}
