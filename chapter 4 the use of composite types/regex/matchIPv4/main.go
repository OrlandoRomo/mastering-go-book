// to run we use the following command:
// go run main.go /tmp/ipv4.log | sort -rn | uniq -c | sort -rn
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
	"regexp"
)

var LOG_FILE = "/tmp/ipv4.log"

func init() {
	f, err := os.OpenFile(LOG_FILE, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()
	ipLog := log.New(f, "log: ", log.LstdFlags)
	ipLog.SetFlags(log.LstdFlags | log.Lshortfile)
	ipv4 := []string{
		"76.140.138.233",
		"99.47.183.5",
		"123.6.93.171",
		"236.240.41.23",
		"108.219.108.48",
		"62.198.148.92",
		"183.123.84.114",
		"130.130.111.246",
		"142.187.56.22",
		"225.143.16.169",
		"231.0.160.212",
		"237.248.182.11",
		"220.148.195.118",
		"223.53.217.208",
		"79.56.91.1",
		"104.35.91.197",
		"255.25.52.238",
		"224.101.164.196",
		"75.1.171.93",
		"36.173.126.154",
		"170.176.144.253",
		"27.157.14.204",
		"69.116.172.71",
		"203.130.28.37",
		"116.124.117.167",
		"201.245.107.93",
		"2.234.215.168",
		"69.197.170.163",
		"148.191.26.201",
		"38.40.216.7",
		"17.67.107.93",
		"108.133.52.31",
		"28.78.51.27",
		"216.138.155.86",
		"5.252.175.86",
		"9.116.95.190",
		"106.149.56.65",
		"122.222.70.201",
		"114.95.227.70",
		"212.39.108.104",
		"0.256.1.256",
	}

	for _, ip := range ipv4 {
		ipLog.Printf("ip registered: %s\n", ip)
	}
}

func main() {
	arguments := os.Args
	if len(arguments) < 1 {
		fmt.Printf("usage: %s logFile\n", filepath.Base(arguments[0]))
		os.Exit(1)
	}

	for _, file := range arguments[1:] {
		f, err := os.Open(file)
		if err != nil {
			fmt.Printf("error opening %s file %v", f.Name(), err)
			os.Exit(1)
		}
		r := bufio.NewReader(f)
		for {
			line, err := r.ReadString('\n')
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Printf("error reading file %s", err)
				break
			}
			ip := findIP(line)
			trial := net.ParseIP(ip)
			if trial.To4() == nil {
				continue
			}
			fmt.Println("Valid IP: ", ip)
		}
	}
}

func findIP(input string) string {
	partIP := "(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])"
	grammar := partIP + "\\." + partIP + "\\." + partIP + "\\." + partIP
	matchIP := regexp.MustCompile(grammar)
	return matchIP.FindString(input)
}
