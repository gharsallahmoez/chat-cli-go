package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/gharsallahmoez/chat/src/config"
	pb "github.com/gharsallahmoez/chat/src/pb"
	"github.com/gharsallahmoez/chat/src/utils"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"strings"

	"google.golang.org/grpc"
)

// clientHandler
type clientHandler struct {
	stream     pb.Services_ChatServiceClient
	clientName string
	config *config.Config
	logger *logrus.Logger
}

func main() {
	logger := utils.GetLogger()
	conf, err := config.MakeConfig()
	if err != nil {
		logger.Fatalf( "failed to create the configuration %v",err)
	}

	// connect to the server
	conn, err := grpc.Dial(conf.Server.Host+":"+conf.Server.Port, grpc.WithInsecure())
	if err != nil {
		logger.Fatalf( "failed to connect to the server %v",err)
	}

	// defer close connection
	defer conn.Close()

	// create grpc client
	client := pb.NewServicesClient(conn)

	// create stream with background context
	stream, err := client.ChatService(context.Background())
	if err != nil {
		logger.Fatalf("Failed to call ChatService, %v",err)
	}

	// implement communication with gRPC server
	ch := clientHandler{stream: stream,config: conf,logger: logger}
	ch.clientConfig()
	go ch.sendMessage()
	go ch.receiveMessage()

	// blocker
	bl := make(chan bool)
	<-bl
}

func (ch *clientHandler) clientConfig() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter your username : ")
	for{
		name, err := reader.ReadString('\n')
		if err != nil {
			ch.logger.Fatalf(" Failed to read from console %v", err)
		}
		name = strings.Trim(name, "\r\n")
		if !utils.IsValidUsername(name){
			fmt.Printf("please choose a username between 1 and 20 character \n")
			continue
		}
			ch.clientName = strings.Trim(name, "\r\n")
			err = ch.sendHelloMessage()
			if err != nil {
				ch.logger.Fatalf(" Failed to send hello message %v", err)
			}
			break

		}
}

//send message
func (ch *clientHandler) sendHelloMessage() error{

	clientMessageBox := &pb.FromClient{
		Name: ch.clientName,
		Body: "added to the chat",
	}
	err := ch.stream.Send(clientMessageBox)
	return err
}
//send message
func (ch *clientHandler) sendMessage() {
	// create a loop
	for {
		reader := bufio.NewReader(os.Stdin)
		clientMessage, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf(" Failed to read from console :: %v", err)
		}
		clientMessage = strings.Trim(clientMessage, "\r\n")

		clientMessageBox := &pb.FromClient{
			Name: ch.clientName,
			Body: clientMessage,
		}

		err = ch.stream.Send(clientMessageBox)

		if err != nil {
			log.Printf("Error while sending message to server :: %v", err)
		}
	}
}

//receive message
func (ch *clientHandler) receiveMessage() {
	// create a loop
	for {
		mssg, err := ch.stream.Recv()
		if err != nil {
			log.Printf("Error in receiving message from server :: %v", err)
		}
		// print message to console
		fmt.Printf("%s : %s \n",mssg.Name,mssg.Body)

	}
}
