package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os/exec"
	"strconv"
)

var port int = 8080

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")

	// net name
	out := ""
	cmd := exec.Command("iwgetid", "-r")
	netName, err := cmd.Output()
	if err != nil {
		out += string(fmt.Sprintln(err))
	}
	out += string(netName)

	// IP
	out += fmt.Sprintln(GetOutboundIP())

	// sensors
	cmd = exec.Command("sensors")
	temperature, _ := cmd.Output()
	out += string(temperature)

	io.WriteString(w, out)
}

// debug
func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}

// Get preferred outbound ip of this machine
func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

func main() {
	log.Println("gmonitor starting up...")

	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)

	if err := http.ListenAndServe(":"+strconv.Itoa(port), nil); err != nil {
		log.Fatal("fail to start server:", err)
	}
}
