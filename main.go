package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type Data struct {
	Firstseen	string	`json:"firstseen"`
	DstIP		string	`json:"dstip"`
	DstPort		int		`json:"dstport"`
	LastOnline	string	`json:"lastonline"`
	Malware		string	`json:"malware"`
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func parse(tmp []string) Data {
	var res Data
	for i, d := range tmp {
		f := reflect.ValueOf(&res).Elem().Field(i)
		if f.CanSet() {
			if f.Type().Name() == "string" {
				f.SetString(d)
			}
			if f.Type().Name() == "int" {
				n, err := strconv.ParseInt(d, 10, 64)
				check(err)
				f.SetInt(n)
			}
			if f.Type().Name() == "float" {
				n, err := strconv.ParseFloat(d, 64)
				check(err)
				f.SetFloat(n)
			}
			if f.Type().Name() == "bool" {
				b, err := strconv.ParseBool(d)
				check(err)
				f.SetBool(b)
			}
			// ... them 1 dong thu
		}
	}

	return res
}

func Csv2Json(filename string) []Data {
	f, err := os.OpenFile(filename, os.O_RDONLY|os.O_CREATE, 0644)
	check(err)
	defer f.Close()

	r := bufio.NewReader(f)
	var data []Data
	for line, _, err := r.ReadLine(); err != io.EOF; line, _, err = r.ReadLine() {
		if string(line[0]) == "#" {
			continue
		}
		var rawData Data
		tmp := strings.Split(string(line), ",")
		// fmt.Println(len(tmp), tmp)
		if len(tmp) < 1 {
			panic("empty line")
		} else {
			rawData = parse(tmp)
		}
		data = append(data, rawData)
	}

	return data
}

func Save2File(filename string, rawData []Data) {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0644)
	check(err)
	err = os.Truncate(filename, 0)
	check(err)
	defer f.Close()

	js, err := json.MarshalIndent(rawData, "", "  ")
	// js, err := json.Marshal(rawData)
	check(err)
	fmt.Fprintln(f, string(js))
}

func main() {
	filename := "test2.csv"
	t := time.Now().UnixNano()
	res := Csv2Json(filename)
	Save2File("test2.json", res)
	t = time.Now().UnixNano() - t
	fmt.Println(t / 1000)
}
