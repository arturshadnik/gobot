// db access layer
package db

import (
	"context"
	"log"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/arturshadnik/gobot/backend/internal/models"
)

var Client *firestore.Client

func InitFirestore() {
	c := context.Background()
	var err error
	log.Print("Starting firestore")
	Client, err = firestore.NewClient(c, "research-query")
	if err != nil {
		log.Fatalf("Failed to initialize Firestore client: %v", err)
	}
}

func CloseFirestore() {
	if Client != nil {
		Client.Close()
	}
}

func StoreMessage(message, level, role string) error {
	c := context.Background()
	docRef, _, err := Client.Collection("messages").Add(c, map[string]interface{}{
		"body":      message,
		"role":      role,
		"timestamp": time.Now(),
	})
	if err != nil {
		log.Fatalf("Write to firestore failed! %v", err)
	}

	messageId := docRef.ID
	log.Print(messageId)
	return nil
}

func LoadConversation(id, level string) (models.Conversation, error) {
	return models.Conversation{}, nil
}
