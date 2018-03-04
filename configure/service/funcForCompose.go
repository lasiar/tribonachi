package service

import (
	"fmt"
	"for_job/configure/readConifg"
	"log"
	"os"
	"strings"
)

func SetValue() string {
	var path string
	fmt.Print("where save db? (enter path) ->")
	_, err := fmt.Scanln(&path)
	if err != nil {
		log.Println(err)
	}
	pathExists, err := existsPath(path)
	if err != nil {
		log.Println(err)
	}
	if !pathExists {
		fmt.Println("Dir not found, create or change dir")
		fmt.Print("Create dir? y/n -> ")
		if AskForConfirmation() {
			err := os.MkdirAll(path, 0777)
			if err != nil {
				log.Println(err)
				return SetValue()
			}
			return path
		} else {
			return SetValue()
		}
	}
	return path
}

func SetPortDocer() string {
	portOut := SetPort()
	portIn := setPortInDocker()
	return portOut + ":" + portIn
}

func setPortInDocker() string {
	var c readConifg.ConfigHttp
	err := c.ReadConfig()
	if err != nil {
		log.Println("err read config http: ", err, "because need create it, (./configure -http-config)")
		os.Exit(1)
		return setPortInDocker()
	}
	return strings.Replace(c.Port, ":", "", -1)
}

func existsPath(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}