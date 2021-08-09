package grpc

import (
	pb "github.com/gharsallahmoez/chat/src/pb"
	"sync"
	"time"
)

type messageUnit struct {
	ClientName        string
	MessageBody       string
}

type messageHandle struct {
	MQue []messageUnit
	mu   sync.Mutex
}

var messageHandleObject = messageHandle{}

//define ChatService
func (runner *SvcRunner) ChatService(csi pb.Services_ChatServiceServer) error {
	errch := make(chan error)
	// receive messages - init a go routine
	go runner.receiveFromStream(csi, errch)
	// send messages - init a go routine
	go runner.sendToStream(csi, errch)
	return <-errch
}

//receive messages
func (runner *SvcRunner) receiveFromStream(csi_ pb.Services_ChatServiceServer, errch_ chan error) {
	//implement a loop
	for {
		mssg, err := csi_.Recv()
		if err != nil {
			runner.Logger.Printf("Error in receiving message from client :: %v", err)
			errch_ <- err
		} else {
			messageHandleObject.mu.Lock()
			messageHandleObject.MQue = append(messageHandleObject.MQue, messageUnit{
				ClientName:        mssg.Name,
				MessageBody:       mssg.Body,
			})
			if _, ok := runner.Users[mssg.Name]; !ok {
				runner.Users[mssg.Name] = struct{}{}
			}

			runner.Logger.Printf("%v", messageHandleObject.MQue[len(messageHandleObject.MQue)-1])
			messageHandleObject.mu.Unlock()
		}
	}
}

//send message
func (runner *SvcRunner) sendToStream(csi_ pb.Services_ChatServiceServer,  errch_ chan error) {
	for {
		// loop through messages in MQue
		for {
			time.Sleep(500 * time.Millisecond)
			messageHandleObject.mu.Lock()
			if len(messageHandleObject.MQue) == 0 {
				messageHandleObject.mu.Unlock()
				break
			}
			senderName4Client := messageHandleObject.MQue[0].ClientName
			message4Client := messageHandleObject.MQue[0].MessageBody
			messageHandleObject.mu.Unlock()
				for user := range(runner.Users) {
					// send message to designated client (do not send to the same client)
					if user != senderName4Client {
						err := csi_.Send(&pb.FromServer{Name: senderName4Client, Body: message4Client})
						if err != nil {
							errch_ <- err
						}
						messageHandleObject.mu.Lock()
						if len(messageHandleObject.MQue) > 1 {
							messageHandleObject.MQue = messageHandleObject.MQue[1:] // delete the message at index 0 after sending to receiver
						} else {
							messageHandleObject.MQue = []messageUnit{}
						}
						messageHandleObject.mu.Unlock()
					}
				}

		}
		time.Sleep(100 * time.Millisecond)
	}
}
