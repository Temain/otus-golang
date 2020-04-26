package api

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	event "github.com/Temain/otus-golang/hw-26/internal/proto"
	"google.golang.org/grpc"

	"github.com/golang/protobuf/ptypes"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var GrpcClientCmd = &cobra.Command{
	Use:   "grpc_client",
	Short: "run grpc client",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("running gRPC client...")

		ctx, _ := context.WithTimeout(context.Background(), 5*time.Minute)
		cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("could not connect: %v", err)
		}
		defer cc.Close()

		conn := event.NewEventServiceClient(cc)
		termChan := make(chan os.Signal)
		go selectRoutine(ctx, conn, termChan)
		signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)
		<-termChan
	},
}

func selectRoutine(ctx context.Context, conn event.EventServiceClient, termChan chan os.Signal) {
OUTER:
	for {
		method := selectMethod()
		switch method {
		case "List":
			listEvents(ctx, conn)
			break
		case "Search":
			searchEvent(ctx, conn)
			break
		case "Add":
			addEvent(ctx, conn)
			break
		case "Update":
			updateEvent(ctx, conn)
			break
		case "Delete":
			deleteEvent(ctx, conn)
			break
		case "Exit":
			close(termChan)
			break OUTER
		}
	}
}

func selectMethod() string {
	prompt := promptui.Select{
		Label: "Select method",
		Items: []string{"List", "Search", "Add", "Update", "Delete", "Exit"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("prompt failed %v\n", err)
		return ""
	}

	fmt.Printf("you choose %q\n", result)

	return result
}

func listEvents(ctx context.Context, conn event.EventServiceClient) {
	stream, err := conn.List(ctx, &event.ListRequest{})
	if err != nil {
		log.Fatalf("error on list events: %v", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error on receive list of events: %v", err)
		}
		log.Printf("received event: %v\n", msg)
	}
}

func searchEvent(ctx context.Context, conn event.EventServiceClient) {
	sample := time.Date(2020, 04, 22, 10, 00, 00, 00, time.UTC)
	created, err := ptypes.TimestampProto(sample)
	if err != nil {
		log.Fatalf("wrong event date: %v", err)
	}
	response, err := conn.Search(ctx, &event.SearchRequest{Date: created})
	if err != nil {
		log.Fatalf("error on search event: %v", err)
	}

	msg := response.Event
	if msg == nil {
		log.Println("event not found")
	}

	log.Printf("found event: %v\n", msg)
}

func addEvent(ctx context.Context, conn event.EventServiceClient) {
	created, err := ptypes.TimestampProto(time.Now())
	if err != nil {
		log.Fatalf("wrong event date: %v", err)
	}

	message := &event.EventMessage{
		Title:       "Sample event title",
		Description: "Sample event description",
		Created:     created,
	}
	request := &event.AddRequest{Event: message}
	response, err := conn.Add(ctx, request)
	if err != nil {
		log.Fatalf("error on add event: %v", err)
	}

	success := response.Success
	if !success {
		log.Fatalf("new event not added")
	}

	log.Printf("added new event: %v\n", message)
}

func updateEvent(ctx context.Context, conn event.EventServiceClient) {
	created, err := ptypes.TimestampProto(time.Now())
	if err != nil {
		log.Fatalf("wrong event date: %v", err)
	}

	message := &event.EventMessage{
		Id:          2,
		Title:       "Evening tea (updated)",
		Description: "Not bad (updated)",
		Created:     created,
	}
	request := &event.UpdateRequest{Event: message}
	response, err := conn.Update(ctx, request)
	if err != nil {
		log.Fatalf("error on update event: %v", err)
	}

	success := response.Success
	if !success {
		log.Fatalf("event not updated")
	}

	log.Printf("event updated: %v\n", message)
}

func deleteEvent(ctx context.Context, conn event.EventServiceClient) {
	request := &event.DeleteRequest{Id: 1}
	response, err := conn.Delete(ctx, request)
	if err != nil {
		log.Fatalf("error on delete event: %v", err)
	}

	success := response.Success
	if !success {
		log.Printf("event not deleted")
	}

	log.Println("event(id=1) deleted")
}
