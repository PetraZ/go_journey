package main

import (
	context "context"
	fmt "fmt"
	"net"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type addServer struct{}

func main() {
	// Layer 3 tcp listener
	listener, err := net.Listen("tcp", "localhost:7011")
	if err != nil {
		fmt.Println(err)
	}
	// you can add options if you want...
	srv := grpc.NewServer()
	RegisterAddServiceServer(srv, &addServer{})
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}

	//exmaple curl : grpcurl -d '{"a":1,"b":2}' -v -plaintext localhost:7011 main.AddService/Add
}

func (s *addServer) Add(ctx context.Context, r *Request) (*Response, error) {
	return &Response{Result: r.A + r.B}, nil
}
