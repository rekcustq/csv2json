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

	"./data"
)

// type Data data.IPBlockList
type Data data.User

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
			switch f.Type().Name() {
				case "string": {
					f.SetString(d)
				}
				case "int": {
					if d == "" {
						d = "0"
					}	
					n, err := strconv.ParseInt(d, 10, 64)
					check(err)
					f.SetInt(n)
				}
				case "float": {
					if d == "" {
						d = "0"
					}
					n, err := strconv.ParseFloat(d, 64)
					check(err)
					f.SetFloat(n)
				}
				case "bool": {
					if d == "" {
						d = "false"
					}
					b, err := strconv.ParseBool(d)
					check(err)
					f.SetBool(b)
				}
				// ... them 1 dong thu
			}
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
	filename := "test.csv"
	t := time.Now().UnixNano()
	res := Csv2Json(filename)
	Save2File("test.json", res)
	t = time.Now().UnixNano() - t
	fmt.Println(t / 1000)
}