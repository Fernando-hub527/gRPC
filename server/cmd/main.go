package main

import (
	"fmt"
	"net"

	"github.com/Fernando-hub527/gRPC/internal/pb"
	"github.com/Fernando-hub527/gRPC/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

/*
	Para realizar os testes, você pode usar o evans, o evans é um cliente que te permite acessar o servidor grpc
	e simular as chamadas às funções, basta instalar seguindo as orientações oficiais e rodar o comando:
	evans -r repl
*/

func main() {
	serviceEncomenda := service.NewEncomendaService()

	grpcServer := grpc.NewServer() // CRiação do servidor grpc

	pb.RegisterEncomendaServiceServer(grpcServer, serviceEncomenda) // Aqui você precisa vincular o servidor grpc e o serviço criado que você implementou
	reflection.Register(grpcServer)                                 // Não entendi para que serve

	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(listen); err != nil {
		panic(err)
	}

	fmt.Println("Setvidor iniciado")
}
