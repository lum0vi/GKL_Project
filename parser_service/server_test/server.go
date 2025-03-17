package main

// import (
// 	"context"
// 	"fmt"
// 	"parser_service/internal/logger"

// 	api "github.com/lum0vi/api_gkl"
// 	"google.golang.org/grpc"
// )

// func main() {
// 	ctx := context.Background()
// 	ctx, _ = logger.New(ctx)
// 	grpcConn, _ := grpc.Dial("localhost:50051", grpc.WithInsecure())
// 	defer grpcConn.Close()
// 	mainservice := api.NewParserClient(grpcConn)
// 	res, _ := mainservice.GetParserElements(ctx, &api.GetParserElementsRequest{SiteUrl: "https://megakino42.ru/?facility=yubileynyy",
// 		Selection: "div.EventList_event__OjvqQ"})

// 	for _, val := range res.Selectionelement {
// 		fmt.Println(val.Title, val.AgeLimit, val.HrefImage, val.HrefSeans, val.Price, val.TimeSeans)
// 	}
// }
