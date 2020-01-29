// UDP server : 10.100.23.147
// Our pc : 10.100.23.230


package main
import (. "fmt" 
	"runtime"
	"time"
	."net"
)

func checkError(err error) {
	if err != nil {
		Println("Noe gikk galt %v", err)
		return
	}
	
}

func read(udp_addr *UDPAddr) {
	buffer := make([]byte, 1024)
	conn, err := ListenUDP("udp", udp_addr)
	checkError(err)
	for {
		buffer = nil
		n,err := conn.Read(buffer)
		checkError(err)
		Printf("Rcv %d bytes: %s\n",n, buffer)
	}	
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	udpAddr, err := ResolveUDPAddr("udp", "10.100.23.147:20015") // IP of what we are sending to
	//udpAddr2, err := ResolveUDPAddr("udp", "10.100.23.147:20016")
	checkError(err)
	conn, err := DialUDP("udp", nil, udpAddr) // open connection for sending udp messages
	checkError(err)
	
	//go read(udpAddr) // start thread (go routine) read
	
	for {
		time.Sleep(1000*time.Millisecond)
		
		// WRITE
		Println("sender meldinger")
		_, err := conn.Write([]byte("Eg er her\n")) // \x00
		checkError(err)
				
		// READ
		//buffer := make([]byte, 1024)
		//n,err := conn.Read(buffer)
		//checkError(err)
		//Printf("Rcv %d bytes: %s\n",n, buffer)
	}
}