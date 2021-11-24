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
	l1 := strings.Index(data, "\r\n")
	if l1 < 0 {
		return errors.New("l1 < 0, no \\r\\n found")
	}
	this.comment = strings.TrimSpace(data[l1:])
	this.comment = strings.ReplaceAll(this.comment, "\r\n", " ")
	data = data[:l1]

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

	sliSvnLog := make([]*SvnLog, 0)
	for {
		for f := strings.Index(svnData, "-"); f == 0; f = strings.Index(svnData, "-") {
			svnData = svnData[1:]
		}
		t := strings.Index(svnData, "---")
		if t < 0 {
			//处理最后一个数据不需要处理
			break
		} else {
			data := svnData[:t]
			data = strings.TrimSpace(data)
			data = strings.ReplaceAll(data, "\r\n\r\n", "\r\n")
			ptrSvnLog := NewSvnLog()
			if err := ptrSvnLog.Load(data); err != nil {
				return sliSvnLog, err
			}
			ptrSvnLog.Test()
			sliSvnLog = append(sliSvnLog, ptrSvnLog)
			//fmt.Println("loop===============")
			//fmt.Println(data)
			svnData = svnData[t:]
		}
	}

	return sliSvnLog, nil
}
