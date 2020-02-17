//
// Copyright (C) 2019 yech <yech1990@gmail.com>
//
// Distributed under terms of the MIT license.
// ssqr

package main

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/mdp/qrterminal"
)

type Server struct {
	Method   string `json:"method"`
	Password string `json:"password"`
	Server   string `json:"server"`
	Port     int    `json:"server_port"`
}

func main() {
	if len(os.Args) != 2 {
		usage()
		os.Exit(1)
	}
	filePath := os.Args[1]
	// read profile name
	fileName := filepath.Base(filePath)
	baseName := strings.TrimSuffix(fileName, filepath.Ext(fileName))
	baseName = url.QueryEscape(baseName)
	baseName = strings.ReplaceAll(baseName, "_", "%20")

	// Open our jsonFile
	jsonFile, err := os.Open(filePath)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var server Server

	json.Unmarshal(byteValue, &server)

	// for debug
	fmt.Printf("read config form %s\n", os.Args[1])
	fmt.Println("--------------------------")
	fmt.Printf("Profile Name: %s\n", baseName)
	fmt.Printf("Server Method: %s\n", server.Method)
	fmt.Printf("Server Password: %s\n", server.Password)
	fmt.Printf("Server Server: %s\n", server.Server)
	fmt.Printf("Server Port: %d\n", server.Port)
	fmt.Println("")
	// generate QR code
	config := fmt.Sprintf("%s:%s@%s:%d", server.Method, server.Password, server.Server, server.Port)
	urlStr := fmt.Sprintf("ss://%s#%s", b64.StdEncoding.EncodeToString([]byte(config)), baseName)
	fmt.Println(urlStr)
	qrCode(urlStr)
	fmt.Println("")
}

func usage() {
	fmt.Printf("Usage:\n%s ss_config.json\n", os.Args[0])
}

func qrCode(urlStr string) {

	// QR code
	config := qrterminal.Config{
		Level:     qrterminal.L,
		Writer:    os.Stdout,
		BlackChar: qrterminal.BLACK,
		WhiteChar: qrterminal.WHITE,
		QuietZone: 1,
	}
	qrterminal.GenerateWithConfig(urlStr, config)
}
