// db access layer
package db

import (
	"context"
	"log"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/arturshadnik/gobot/backend/internal/models"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func StoreMessage(message, level, role, userId string) error {
	c := context.Background()

	convoId := level + userId
	convoRef := Client.Collection("conversations").Doc(convoId)

	_, err := convoRef.Get(c)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			_, err = Client.Collection("conversations").Doc(convoId).Create(c, map[string]interface{}{
				"timestamp": time.Now(),
			})
			if err != nil {
				return err
			}
		} else {
			log.Printf("Something went wrong: %v", err)
			return err
		}
	}
	_, err = convoRef.Collection("messages").NewDoc().Create(c, map[string]interface{}{
		"role": role, "content": message, "timestamp": time.Now(),
	})
	if err != nil {
		return err
	}

	return nil
}

func GetMessages(convoId string) ([]models.Message, error) {
	c := context.Background()

	convoRef := Client.Collection("conversations").Doc(convoId)

	_, err := convoRef.Get(c)

	if err != nil {
		return []models.Message{}, err
	}

	var messages []models.Message

	messagesIter := convoRef.Collection("messages").Documents(c)

	for {
		doc, err := messagesIter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return []models.Message{}, err
		}
		var newMessage models.Message
		doc.DataTo(&newMessage)
		messages = append(messages, newMessage)
	}
	return messages, nil
}

func ClearConversation(convoId string) error {
	c := context.Background()

	err := ClearMessages(c, convoId)
	if err != nil {
		return err
	}

	convoRef := Client.Collection("conversations").Doc(convoId)
	_, err = convoRef.Delete(c)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func ClearMessages(c context.Context, convoId string) error {
	miniBatch := 32

	messages := Client.Collection("conversations").Doc(convoId).Collection("messages")

	bulkwriter := Client.BulkWriter(c)

	for {
		iter := messages.Limit(miniBatch).Documents(c)
		numDel := 0

		for {
			doc, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				return err
			}

			bulkwriter.Delete(doc.Ref)
			numDel++
		}

		if numDel == 0 {
			bulkwriter.End()
			break
		}

		bulkwriter.Flush()
	}
	return nil
}
