package domain

type TransactionStatus string

const (
	StatusPending    TransactionStatus = "peding"
	StatusProcessing TransactionStatus = "processing"
	StatusApproved   TransactionStatus = "approved"
	StatusRejected   TransactionStatus = "rejected"
	StatusRefunded   TransactionStatus = "refunded"
)
