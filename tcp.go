//10.100.23.147:33546, IP:port

package main
import (. "fmt"
	"runtime"
	"time"
	."net"
)

func checkError(err error) {
	if err != nil {
		Println("Feil %v", err)
		return
	}
	
}

func read(listener *TCPConn) {
	buffer := make([]byte, 1024)
	//tcp_addr, err := ResolveTCPAddr("tcp", "10.100.23.147:33546")
	//listener, err := DialTCP("tcp", nil, tcp_addr) // connects to a server
	//checkError(err)
	Println("Hei!")
	for {
		Println("Hei2!")
		_,err := listener.Read(buffer)
		checkError(err)
		Printf(string(buffer))
	}	
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	tcp_addr, err := ResolveTCPAddr("tcp", "10.100.23.147:33546")
	checkError(err)
	conn, err := DialTCP("tcp", nil, tcp_addr)
	checkError(err)
	conn.Write([]byte("Connect to: 10.100.23.147:33546\x00"))
	
	go read(conn)
	
	for {

		time.Sleep(100*time.Millisecond)
		// WRITE
		//Println("Er du der server??")
		_, err := conn.Write([]byte("Jeg er her\n\x00")) // \x00
		checkError(err)		
	}
}