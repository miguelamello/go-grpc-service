package main

import (
	//"context"
	//"fmt"
	//"log"
	//"net"
	//"google.golang.org/grpc"
	ppb "go_grpc_service/go_grpc_service"
)

type server struct{}

func (s *server) sendResponse( req *ppb.Request, stream ppb.GoGrpcService_SendResponseServer) error {
	//log.Printf("Received: %v", req.GetMessage())
	//return nil
	return stream.Send(&ppb.Response{Message: "Hello " + req.GetMessage()})
}
