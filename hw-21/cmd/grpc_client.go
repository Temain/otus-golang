package cmd

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	event "github.com/Temain/otus-golang/hw-21/internal/proto"
	"google.golang.org/grpc"

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
		end := make(chan interface{})
		// go writeRoutine(end, ctx, c)

		stream, err := conn.List(ctx, &event.ListRequest{})
		if err != nil {
			log.Fatalf("error on list events: %v", err)
		}

		for {
			event, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("error on receive list of events: %v", err)
			}
			log.Printf("received event: %v", event)
		}

		<-end
	},
}

//func writeRoutine(end chan interface{}, ctx context.Context, conn event.EventServiceClient) {
//	scanner := bufio.NewScanner(os.Stdin)
//OUTER:
//	for {
//		select {
//		case <-ctx.Done():
//			break OUTER
//		default:
//			if !scanner.Scan() {
//				break OUTER
//			}
//			str := scanner.Text()
//
//			if str == "end" {
//				break OUTER
//			}
//			log.Printf("To server %v\n", str)
//
//			msg, err := conn.SendMessage(context.Background(), &event.EventMessage{
//				Title: str,
//				// Description:
//				Date: ptypes.TimestampNow(),
//			})
//
//			if err != nil {
//				errMsg := status.Convert(err)
//				fmt.Printf("err %s %s", errMsg.Code(), errMsg.Message())
//			}
//
//			if msg != nil {
//				created, _ := ptypes.Timestamp(msg.Date)
//				created = created.Local()
//				fmt.Printf("%s created %s", msg.Title, created)
//			}
//
//		}
//	}
//
//	log.Printf("finished writeRoutine")
//	close(end)
//}
