package main

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type SvnVersion struct {
	From int `yaml:"from"`
	To   int `yaml:"to"`
}

func NewSvnVersion() *SvnVersion {
	return &SvnVersion{}
}

type Config struct {
	SvnPath    string      `yaml:"svnPath"`
	SvnVersion *SvnVersion `yaml:"svnVersion"`
	Output     string      `yaml:"output"`
}

func NewConfig() *Config {
	return &Config{
		SvnVersion: NewSvnVersion(),
	}
}

func (this *Config) Load() error {
	f, e := ioutil.ReadFile("svnlog.yaml")
	if e != nil {
		return errors.New(fmt.Sprintf("读取配置文件错误:%v", e))
	}

	e = yaml.Unmarshal(f, this)
	if e != nil {
		return errors.New(fmt.Sprintf("解析yaml错误:%v", e))
	}

	return nil
}

func (this *Config) GetVersion() string {
	return fmt.Sprintf("r%v:r%v", this.SvnVersion.From, this.SvnVersion.To)
}
