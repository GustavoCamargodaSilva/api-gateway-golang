package domain

import (
	"errors"
)

var (
	// ErrInvalidAmount indica que o valor da transação é inválido
	ErrorInvalidAmount = errors.New("amount must be greater than zero")

	// ErrInvalidStatus indica que o status da transação é inválido para a operação
	ErrorInvalidStatus = errors.New("invalid transaction status for this operation")

	// ErrAlreadyRefunded indica que a transação já foi extornada
	ErrAlreadyRefunded = errors.New("transaction already refunded")

	// ErrCannotRefund indica que a transação não pode ser extornada
	ErrCannotRefund = errors.New("only approved transactions can be refunded")
)
