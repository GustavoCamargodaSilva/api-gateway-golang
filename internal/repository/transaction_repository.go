package repository

import (
	"context"

	"github.com/GustavoCamargodaSilva/payment-gateway/internal/domain"
)

// -> Vai definir os contratos para persistencia
type TransactionRepository interface {

	// -> Criar uma nova transaçao no banco
	Create(ctx context.Context, transaction *domain.Transaction) error

	// -> Buscar a transaçao com base no ID
	FindByID(ctx context.Context, id string) (*domain.Transaction, error)

	// -> Buscar uma transaçao pela chave de idempotencia
	FindByIdempotencyKey(ctx context.Context, key string) (*domain.Transaction, error)

	// -> Apenas atualiza o status da transaçao
	UpdateStatus(ctx context.Context, id string, status domain.TransactionStatus) error

	// ListAll retorna todas as transações (com paginação futura)
	List(ctx context.Context, limit, offset int) ([]*domain.Transaction, error)
}
