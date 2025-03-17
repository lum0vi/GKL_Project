package service

import (
	"context"
	"fmt"
	"parser_service/internal/connection"
	"sync"

	api "github.com/lum0vi/api_gkl"
)

type Service struct {
	api.UnimplementedParserServer
	mx *sync.Mutex
}

func (s *Service) GetParserElements(ctx context.Context, req *api.GetParserElementsRequest) (*api.GetParserElementsResponse, error) {
	var mas []*api.SelectionElement
	var len int
	fmt.Println("yes")
	mas, len = connection.Connection(ctx, req.SiteUrl, req.Selection)
	resp := new(api.GetParserElementsResponse)
	resp.Lenght = int32(len)
	resp.Selectionelement = mas
	return resp, nil
}

func New() *Service {
	return &Service{mx: &sync.Mutex{}}
}
