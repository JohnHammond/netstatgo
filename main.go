package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"
	"syscall"

	"github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/process"
)

func protocolToString(connType uint32, ip string) string {
	isIPv6 := strings.Contains(ip, ":")
	switch connType {
	case syscall.SOCK_STREAM: // TCP
		if isIPv6 {
			return "TCP6"
		}
		return "TCP"
	case syscall.SOCK_DGRAM: // UDP
		if isIPv6 {
			return "UDP6"
		}
		return "UDP"
	default:
		return "Unknown"
	}
}

func main() {
	// Define command-line flags
	portFlag := flag.Int("port", 0, "filter by port number")
	processFlag := flag.String("process", "", "filter by process name")
	flag.Parse()

	// Get network connections
	conns, err := net.Connections("all")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%-7s %-25s %-25s %-10s %-10s\n", "Proto", "Local Address", "Foreign Address", "State", "PID/Name")

	for _, conn := range conns {
		// Filter out connections with no associated process
		if conn.Pid == 0 {
			continue
		}

		// Check if connection matches port filter
		if *portFlag > 0 && int(conn.Laddr.Port) != *portFlag {
			continue
		}

		// Get process name
		proc, err := process.NewProcess(conn.Pid)
		if err != nil {
			log.Println(err)
			continue
		}
		procName, err := proc.Name()
		if err != nil {
			log.Println(err)
			continue
		}

		// Check if connection matches process filter
		if *processFlag != "" && !strings.Contains(strings.ToLower(procName), strings.ToLower(*processFlag)) {
			continue
		}

		// Convert the protocol to a string
		protocol := protocolToString(conn.Type, conn.Laddr.IP)

		// Print connection details
		localAddr := conn.Laddr.IP + ":" + strconv.Itoa(int(conn.Laddr.Port))
		remoteAddr := conn.Raddr.IP + ":" + strconv.Itoa(int(conn.Raddr.Port))
		fmt.Printf("%-7s %-25s %-25s %-10s %-10d/%s\n", protocol, localAddr, remoteAddr, conn.Status, conn.Pid, procName)
	}
}
