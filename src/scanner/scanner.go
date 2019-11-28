package scanner

import (
	"fmt"
	"net"
	"time"
	"strconv"
	"strings"
)

type Scanner struct {
	CIDR string
	Ports string
	Timeout time.Duration
	Protocol string
}

func New(Protocol string, CIDR string, Ports string, Timeout time.Duration) *Scanner {
	return &Scanner{
		CIDR: CIDR,
		Ports: Ports,
		Timeout: Timeout,
		Protocol: Protocol,
	}
}

func (s *Scanner) Scan() {
	workerPool := make(chan int, 100)

	ports := parsePorts(s.Ports)

	var openPorts []int
	results := make(chan int)

	for i := 0; i < cap(workerPool); i++ {
		go worker(workerPool, s.CIDR, s.Protocol, s.Timeout, results)
	}

	go func() {
		for port := range ports {
			workerPool <- port
		}
	}()


	for i := range ports {
		i++
		port := <-results

		if port != 0 {
			openPorts = append(openPorts, port)
		}
	}

	close(workerPool)
	close(results)

	fmt.Println(openPorts)
}

func parsePorts(ports string) []int {
	var parsedPorts []int

	for _, port := range strings.Split(ports, ",") {
		port := strings.Trim(port, "")

		portsRange := strings.Split(port, "-")

		switch len(portsRange) {
		case 1:
			intPort, err := strconv.Atoi(port)

			if err == nil {
				parsedPorts = append(parsedPorts, intPort)
			}
		case 2:
			firstPort, err1 := strconv.Atoi(portsRange[0])
			lastPort, err2 := strconv.Atoi(portsRange[1])

			if err1 == nil && err2 == nil {
				for i := firstPort; i <= lastPort; i++ {
					parsedPorts = append(parsedPorts, i)
				}
			}
		default:
			panic("Cannot parse ports")
		}
	}

	return parsedPorts
}

/**

 */
func worker(workerPool chan int, target string, protocol string, timeout time.Duration, results chan int) {
	for p := range workerPool {
		address := fmt.Sprintf("%s:%d", target, p)

		conn, err := net.DialTimeout(protocol, address, timeout)

		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}