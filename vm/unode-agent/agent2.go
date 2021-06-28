package main

import (
	"fmt"
	"net"
	"os"
    "net/http"
    "encoding/json"
    "io/ioutil"
)


func getHostIpaddr() {
	conn, error := net.Dial("udp", "8.8.8.8:80")
	if error != nil {
		fmt.Println(error)
	}
	defer conn.Close()
	ipAddress := conn.LocalAddr().(*net.UDPAddr)
	fmt.Println(ipAddress)
}


type IP struct {
    Query string
}

func getHostIpaddrPublic() string {

// func getip2() string {
    req, err := http.Get("http://ip-api.com/json/")
    // req, err := http.Get("https://api64.ipify.org?format=json")
    if err != nil {
        return err.Error()
    }
    defer req.Body.Close()

    body, err := ioutil.ReadAll(req.Body)
    if err != nil {
        return err.Error()
    }

    var ip IP
    json.Unmarshal(body, &ip)

	fmt.Println("ip public:", ip )
    return ip.Query
}



func getHostName() {
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	fmt.Println("hostname:", name)
	// fqdn "github.com/Showmax/go-fqdn"
	// fmt.Println(fqdn.Get())
}






func main() {
	getHostIpaddr()
	getHostName()
    getHostIpaddrPublic()
}

