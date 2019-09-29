package main

import (
	"context"
	"fmt"
	accountpb "github.com/OdaDaisuke/gae_sand/pb/account"
	contentpb "github.com/OdaDaisuke/gae_sand/pb/content"
	"github.com/OdaDaisuke/gae_sand/services/main/config"
	gateway "github.com/OdaDaisuke/gae_sand/services/main/gateways"
	handler "github.com/OdaDaisuke/gae_sand/services/main/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {
	configs := config.NewAppConfig()

	grpcFac, err := gateway.NewGrpcGatewayFactory(configs.GrpcHosts)
	if err != nil {
		log.Printf("gRPC initialize error")
		panic(err)
	}

	accountClient := accountpb.NewAccountServiceClient(grpcFac.AccountConn)
	contentClient := contentpb.NewContentServiceClient(grpcFac.ContentConn)
	ctx := context.TODO()

	accountRoute := handler.NewAccountRoute(accountClient, ctx)
	contentRoute := handler.NewContentRoute(contentClient, ctx)

	r := mux.NewRouter()
	r.Handle("/account/signin", accountRoute.SigninHandler)
	r.Handle("/content/{id}", contentRoute.GetContentHandler)
	r.Handle("/contents/{after_id}/{before_id}", contentRoute.GetContentsHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	log.Printf("Listening on port %s", port)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
