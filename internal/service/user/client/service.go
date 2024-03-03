package client

import (
	"github.com/arslanovdi/omp-bot/internal/model/user"
)

type DummyClientService struct {
	clients []user.Client
}

func NewClientService() *DummyClientService {
	return &DummyClientService{
		clients: []user.Client{ // Заглушка 10 клиентов
			{"one"},
			{"two"},
			{"three"},
			{"four"},
			{"five"},
			{"six"},
			{"seven"},
			{"eight"},
			{"nine"},
			{"ten"},
		},
	}
}
