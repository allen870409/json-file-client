package main

import (
	"flag"
	"fmt"
	"net/http"
	"io/ioutil"
	"bytes"
	"log"
	"os"
)

var host = "http://127.0.0.1:9000"
var inputPath, outputPath, verb, path string
var limit int


func main() {
	flag.Parse()
	verb = flag.Arg(0)

	switch verb{
	case "create":
		Create()
	case "get":
		Get()
	case "delete":
		Delete()
	case "update":
		Update()
	case "list":
		List()
	default:
		fmt.Fprintln(os.Stdout, "wrong verb!!!")
		return
	}
}


func readJson(filePath string) string {
	json, err := ioutil.ReadFile(filePath)
	if err != nil {
		return ""
	}
	return string(json)
}

func Create() {
	flagSet := flag.NewFlagSet(os.Args[0],flag.ExitOnError)
	flagSet.StringVar(&inputPath, "i", "", "入力先")
	flagSet.Parse(os.Args[2:])
	var jsonString string
	if inputPath == "" {
		path = flag.Arg(1)
		jsonString = flag.Arg(2)
	}else{
		path = flag.Arg(2)
		jsonString = readJson(inputPath)
	}

	body := bytes.NewBuffer([]byte(jsonString))

	client := &http.Client{}
	req, _ := http.NewRequest("PUT", host + path, body)
	res, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
		return
	}
	result, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("%s", result)
}

func Get() {
	flagSet := flag.NewFlagSet(os.Args[0],flag.ExitOnError)
	flagSet.StringVar(&outputPath, "o", "", "出力先")
	flagSet.Parse(os.Args[2:])

	if outputPath == "" {
		path = flag.Arg(1)
		resp, err := http.Get(host + path)
		if err != nil {
			logError(err)
			return
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		fmt.Printf("%s\n", body)
	}else{
		path = flag.Arg(2)
		resp, err := http.Get(host + path)
		if err != nil {
			logError(err)
			return
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		err = ioutil.WriteFile(outputPath, body, os.ModeAppend)
		if err != nil {
			logError(err)
			return
		}
		fmt.Fprintln(os.Stdout, "Finish writing:", outputPath)
		return
	}
}

func Delete() {
	path = flag.Arg(1)
	client := &http.Client{}
	req, _ := http.NewRequest("DELETE", host + path, nil)
	res,err := client.Do(req)

	if err != nil {
		log.Fatal(err)
		return
	}
	result, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("%s", result)
}

func Update() {
	flagSet := flag.NewFlagSet(os.Args[0],flag.ExitOnError)
	flagSet.StringVar(&inputPath, "i", "", "入力先")
	flagSet.Parse(os.Args[2:])
	var jsonString string
	if inputPath == "" {
		path = flag.Arg(1)
		jsonString = flag.Arg(2)
	}else{
		path = flag.Arg(2)
		jsonString = readJson(inputPath)
	}
	body := bytes.NewBuffer([]byte(jsonString))

	res,err := http.Post(host + path, "application/json;charset=utf-8", body)

	if err != nil {
		log.Fatal(err)
		return
	}
	result, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("%s", result)
}

func List() {
	flagSet := flag.NewFlagSet(os.Args[0],flag.ExitOnError)
	flagSet.IntVar(&limit, "limit", 0, "最大取得数")
	flagSet.Parse(os.Args[3:])
	path = flag.Arg(1)
	limitStr := ""
	if limit > 0{
		limitStr = "?limit=" + fmt.Sprintf("%d",  limit)
	}
	resp, err := http.Get(host + path + limitStr)
	if err != nil {
		logError(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("%s\n", body)
}

func logError(err error) {
	fmt.Fprintln(os.Stderr, err)
}


