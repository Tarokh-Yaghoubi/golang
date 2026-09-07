
package main

import "fmt"
import "time"


type Server struct {
	users map[string]string

}

func NewServer() *Server {
	return &Server {
		users: make(map[string]string)	
	}
}

func (s *Server) addUser(user string) {
	s.users[user] = user
}


func main() {
	userchannel := make(chan string) 	// this is like a pipe

	bufferedChannel := make(chan string, 2)	// this is like a pipe (this is a buffered channel and it will block when it is full)

	go func() {
		time.Sleep(2 * time.Second)
		userchannel <- "Tarokh"	//  it will block here on the SPAWNED goroutine
	}()

	// this will wait for the channel to get its data
	// because that unbuffered channel is blocking
	user := <- userchannel
	fmt.Println(user) 

	bufferedChannel <- "Michael Gray"	// this will not block because this channel is not full
	bufferedChannel <- "John Smith"
	

	firstData := <- bufferedChannel
	bufferedChannel <- "FOO"
	SecondData := <- bufferedChannel

	fmt.Println("data -> ", firstData)
	fmt.Println("second data -> ", SecondData)

}

// the syntax below means that this channel is only meant for "sending to a channel"
// and not READING FROM A CHANNEL
// mychannelName chan <- string
func sendMsg(msgChannel chan<- string) {
	msgChannel <- "This is the msg\n"

	// msg :=  <-  myChannel 	// this is problematic

}


// the syntax below means that this function is only meant for "reading from a channel", 
// so you cannot write to the channel which is passed as an argument to the func
// mychannelName <- chan string
func readMsg(msgChannel <- chan string) {
	msg := <- msgChannel
	fmt.Println("msg is => ", msgh)

	// msgChannel <- "This is a msg to be sent to the channel\n"	// this is problematic
}

