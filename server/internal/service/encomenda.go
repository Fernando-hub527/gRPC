package service

import (
	"context"
	"fmt"

	"github.com/Fernando-hub527/gRPC/internal/pb"
)

/*
Aqui você começa a implementar suas funcionalidades de acordo com o que foi definido no
arquivo .proto.

UnimplementedEncomendaServiceServer serve
*/
type EncomendaService struct {
	pb.UnimplementedEncomendaServiceServer
}

/*
Aqui estamos disponibilizando uma função para criar a estrutura EncomendaService, se você é novo em go e já trabalhou
com linguagens orientadas a objeto, essa função é algo parecido com um constructor
*/
func NewEncomendaService() *EncomendaService {
	return &EncomendaService{}
}

/*
Aqui você implementa o serviço seguindo a interface que está definida no arquivo *_grpc gerado a partir do
arquivo .proto
*/
func (encomendaS *EncomendaService) SendEncomenda(ctx context.Context, encomenda *pb.Encomenda) (*pb.EncomendaRastreio, error) {
	envio := &pb.EncomendaRastreio{Encomenda: encomenda, CodRastreio: 1}
	fmt.Println("Encomenda enviada com sucesso: \n", envio)
	return envio, nil
}
