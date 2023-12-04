package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

var dnsServers = [2]string{"178.22.122.100", "185.51.200.2"}

func main() {
	fmt.Println("start setting dns")
	var interfaceName string = "Wi-Fi 2"
	if isDnsSetted(interfaceName) {
		removeDnsCommand(interfaceName)
		fmt.Println("REMOVED")
	}else {
		for i := 0; i < 2; i++ {
			setDnsCommand(dnsServers[i], interfaceName, i+1)
		}
		fmt.Println("DONE")
	}
	fmt.Println("press enter to exit")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}


func isDnsSetted(interfaceName string)(bool){
	command := fmt.Sprintf("Get-DnsClientServerAddress -AddressFamily IPv4 -InterfaceAlias '%s'", interfaceName)
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


	strOut := string(out)
	regex, err := regexp.Compile("\\{(.*?)\\}")

	regexoutput := regex.FindAllString(strOut, 1)[0]
	regexoutput = regexoutput[1:len(regexoutput)-1]
	arrdns := strings.Split(regexoutput, ", ")
	if len(arrdns) != 2 {
		return false
	}
	for count, dns := range arrdns {
		if dns != dnsServers[count] {
			return false
		}
	}
	return true
}

func removeDnsCommand(interfaceName string){
	command := fmt.Sprintf("netsh interface ipv4 delete dns name=\"%s\" all", interfaceName)
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
