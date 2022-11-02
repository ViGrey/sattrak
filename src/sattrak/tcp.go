package main

import (
	"net"
	"sync"
	//"time"
)

var (
	connections          = map[net.Conn]*sync.RWMutex{}
	connectionsTmp       = map[net.Conn]*sync.RWMutex{}
	connectionsCopyMutex = sync.RWMutex{}
)

func startServer() {
	l, _ := net.Listen("tcp", "127.0.0.1:25544")
	defer l.Close()
	for {
		conn, _ := l.Accept()
		if conn != nil {
			connections[conn] = &sync.RWMutex{}
			connectionsCopyMutex.Lock()
			connectionsTmp = map[net.Conn]*sync.RWMutex{}
			for conn := range connections {
				connectionsTmp[conn] = connections[conn]
			}
			writeToConn(conn, satInputs)
			writeToConn(conn, homeSunMoonInputs)
			connectionsCopyMutex.Unlock()
		}
	}
}

func writeToConn(conn net.Conn, content []byte) {
	if connectionsTmp[conn] != nil {
		if len(content) > 0 {
			connectionsTmp[conn].Lock()
			_, err := conn.Write(content)
			connectionsTmp[conn].Unlock()
			if err != nil {
				delete(connections, conn)
			}
		}
	} else {
		delete(connections, conn)
	}
}

func writeToAllConns(content []byte) {
	connectionsCopyMutex.Lock()
	connectionsTmp = map[net.Conn]*sync.RWMutex{}
	for conn := range connections {
		connectionsTmp[conn] = connections[conn]
	}
	var wg sync.WaitGroup
	for conn := range connectionsTmp {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			writeToConn(conn, content)
		}(&wg)
	}
	wg.Wait()
	connectionsCopyMutex.Unlock()
}
