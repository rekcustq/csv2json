package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

type Data struct {
	Time string `json:"time"`
	IP   string `json:"ip"`
	Port string `json:"port"`
	Date string `json:"date"`
	Name string `json:"name"`
}

func check(err error) {
	if err != nil {
		panic(err)
	}
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
			rawData.Time = tmp[0]
			rawData.IP = tmp[1]
			rawData.Port = tmp[2]
			rawData.Date = tmp[3]
			rawData.Name = tmp[4]
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
	res := Csv2Json(filename)
	Save2File("test2.json", res)
	// fmt.Println(res)
}
