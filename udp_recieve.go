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

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // limits num of threads to num of cores
	buffer := make([]byte, 1024) // make an array with size 1024*bytes
	udp_addr, err := ResolveUDPAddr("udp", ":30000") // store adress of port 30000 into udpAddr, udpAddr hvor vi skal lese fra
	checkError(err)
	conn, err := ListenUDP("udp", udp_addr) // open connection for recieving udp messages
	// conn is a socket (connection), port
	checkError(err)
	
	for {
		time.Sleep(1000*time.Millisecond)
		n,err := conn.Read(buffer) // read data from conn into variable buffer
		checkError(err)
		Printf("Rcv %d bytes: %s\n",n, buffer)
	}
}