package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TransactionStatus string

const (
	StatusPending    TransactionStatus = "peding"
	StatusProcessing TransactionStatus = "processing"
	StatusApproved   TransactionStatus = "approved"
	StatusRejected   TransactionStatus = "rejected"
	StatusRefunded   TransactionStatus = "refunded"
)

type Transaction struct {
	ID             primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	IdempotencyKey string             `json:"idempotency_key" bson: "idempotency_key"`
	Amount         int64              `json:"amount" bson: "amount"`
	Status         TransactionStatus  `json: "status" bson:"status"`
	CustomerEmail  string             `json: "customer_email" bson "customer-_email"` //bson:"customer_email" → MongoDB armazena com mesmo nome
	Description    string             `json: "description" bson: "description"`
	WebhookURL     string             `json: "webhook_url,omitempty" bson: "webhook_url.omitempty"`     //omitempty = se vazio, não aparece no JSON nem no MongoDB.
	CardLastFour   string             `json:"card_last_four,omitempty" bson:"card_last_four,omitempty"` //omitempty = se vazio, não aparece no JSON nem no MongoDB.
	CreatedAt      time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt      time.Time          `json:"updated_at" bson:"updated_at"`
}
