package repository

import (
	"context"
	"errors"
	"time"

	"github.com/GustavoCamargodaSilva/payment-gateway/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoTransactionRepository struct {
	collection *mongo.Collection
}

func NewMongoTransactionRepository(db *mongo.Database) TransactionRepository {
	collection := db.Collection("transactions")

	createIndexes(collection)

	return &mongoTransactionRepository{
		collection: collection,
	}
}

func (r *mongoTransactionRepository) Create(ctx context.Context, transaction *domain.Transaction) error {
	if transaction == nil {
		return errors.New("transaction is nil")
	}

	_, err := r.collection.InsertOne(ctx, transaction)
	if err != nil {
		return err
	}

	return nil
}

func createIndexes(collection *mongo.Collection) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	idempotencyIndex := mongo.IndexModel{
		Keys:    bson.D{{Key: "idempotency_key", Value: 1}},
		Options: options.Index().SetUnique(true),
	}

	statusIndex := mongo.IndexModel{
		Keys: bson.D{{Key: "status", Value: 1}},
	}

	emailIndex := mongo.IndexModel{
		Keys: bson.D{{Key: "customer_email", Value: 1}},
	}

	_, err := collection.Indexes().CreateMany(ctx, []mongo.IndexModel{
		idempotencyIndex,
		statusIndex,
		emailIndex,
	})

	if err != nil {
		// Log do erro (println é simples, em produção usaríamos um logger)
		println("Warning: failed to create indexes:", err.Error())
	}
}

func (r *mongoTransactionRepository) FindByID(ctx context.Context, id string) (*domain.Transaction, error) {

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, ErrorNotFound
	}

	var transaction domain.Transaction

	err = r.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&transaction)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, ErrorNotFound
		}
		return nil, err
	}
	return &transaction, nil
}

func (r *mongoTransactionRepository) FindByIdempotencyKey(ctx context.Context, key string) (*domain.Transaction, error) {
	var transaction domain.Transaction

	err := r.collection.FindOne(ctx, bson.M{"idempotency_key": key}).Decode(&transaction)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, ErrorNotFound
		}
		return nil, err
	}

	return &transaction, nil
}

func (r *mongoTransactionRepository) UpdateStatus(ctx context.Context, id string, status domain.TransactionStatus) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ErrorNotFound
	}

	update := bson.M{
		"$set": bson.M{
			"status":     status,
			"updated_at": time.Now(),
		},
	}

	result, err := r.collection.UpdateOne(ctx, bson.M{"_id": objectID}, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return ErrorNotFound
	}

	return nil // Sucesso!
}

func (r *mongoTransactionRepository) List(ctx context.Context, limit, offset int) ([]*domain.Transaction, error) {
	findOptions := options.Find().
		SetLimit(int64(limit)).
		SetSkip(int64(offset)).
		SetSort(bson.D{{Key: "created_at", Value: -1}})

	cursor, err := r.collection.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	var transactions []*domain.Transaction

	if err := cursor.All(ctx, &transactions); err != nil {
		return nil, err
	}

	return transactions, nil
}
