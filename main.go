package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	fmt.Println("start setting dns")

	dnsServers := [2]string{"178.22.122.100", "185.51.200.2"}
	var interfaceName string = "Wi-Fi"
	if len(os.Args) > 1 {
		interfaceName = os.Args[1]
	}
	_ = interfaceName

	for i := 0; i < 2; i++ {
		setDnsCommand(dnsServers[i], interfaceName, i+1)
	}
	
	fmt.Println("DONE")
	fmt.Println("press enter to exit")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func setDnsCommand(dns string, interfaceName string,index int) {
	command := fmt.Sprintf("netsh interface ipv4 add dnsserver \"%s\" address=%s index=%d", interfaceName, dns, index)
	var commandManager *exec.Cmd = exec.Command("powershell", command)

	stdin, err := commandManager.StdinPipe()
	defer stdin.Close()
	if err != nil {
		log.Fatal(err)
	}

	out, err := commandManager.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	_ = out
	fmt.Println(dns + " DONE")
}
