// pp is the unified tool for managing PushPanel.io services from the command line(CLI).

// Copyright 2019 PushPanel Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/user"
)

func callpush(command string, arguments string, ppkey []byte) {
	url := "https://dash.pushpanel.io/api/" + command + arguments
	rtype := "POST"
	rdata := []byte(``)
	if len(arguments) > 1 {
		rtype = "POST"
		rdata = []byte(arguments[1:])
	} else {
		rtype = "GET"
	}
	req, err := http.NewRequest(rtype, url, bytes.NewBuffer(rdata))
	req.Header.Add("Authorization", "Bearer "+string(ppkey))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		error(resp.StatusCode)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func error(resp int) {
	fmt.Printf("PushPanel does not recoginize this command, error in data or no authorization (%d). Try  %s help %s \n", resp, os.Args[0], os.Args[1])
	os.Exit(127)
}

func usage() {
	fmt.Printf("%s \n %s <command> [<args|subcommand>] \n", "Usage:", os.Args[0])
	flag.PrintDefaults()
}

func replaceAtIndex(input string, r byte, i int) string {
	return input[:i] + string(r) + url.QueryEscape(input[i+1:])
}

func main() {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	kpath := usr.HomeDir + "/.pushpanel"
	if _, err := os.Stat(kpath); os.IsNotExist(err) {
		fmt.Printf("Key file not found. Please put your token into %s file. \n", kpath)
		os.Exit(1)
	}
	keyfile, err := os.Open(kpath)
	if err != nil {
		log.Fatal(err)
	}
	defer keyfile.Close()
	ppkey, err := ioutil.ReadAll(keyfile)
	if err != nil {
		log.Fatal(err)
	}
	if len(os.Args) == 1 {
		usage()
		os.Exit(1)
	}

	cmdb := ""
	argb := ""

	for _, arg := range os.Args[1:] {
		if string(arg[0]) == "-" {
			argb = argb + replaceAtIndex(arg, '&', 0)
		} else {
			cmdb = cmdb + arg + "/"
		}
	}

	callpush(cmdb, argb, ppkey)

}
