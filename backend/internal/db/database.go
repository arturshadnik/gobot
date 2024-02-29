// db access layer
package db

import (
	"context"
	"log"
	"reflect"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/arturshadnik/gobot/backend/internal/models"
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
	docRef, _, err := Client.Collection("messages").Add(c, map[string]interface{}{
		"content":   message,
		"role":      role,
		"timestamp": time.Now(),
	})
	if err != nil {
		log.Printf("Write to firestore failed! %v", err)
		return err
	}

	convoId := level + userId
	convoRef := getConvoRef(convoId)

	messageId := docRef.ID
	_, err = convoRef.Get(c)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			messages := []string{messageId}
			_, err = Client.Collection("conversations").Doc(convoId).Create(c, map[string]interface{}{
				"messages": messages,
			})
			if err != nil {
				return err
			}
		} else {
			log.Printf("Something went wrong: %v", err)
			return err
		}

	} else {
		_, err = convoRef.Update(c, []firestore.Update{{Path: "messages", Value: firestore.ArrayUnion(messageId)}})
		if err != nil {
			return err
		}
	}
	return nil
}

func LoadConversation(userId, level string) (models.Conversation, error) {
	c := context.Background()
	convoId := level + userId

	convoRef := getConvoRef(convoId)
	convoSnapshot, err := convoRef.Get(c)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			log.Printf("Conversation not found: %v", err)
		} else {
			log.Printf("Something went wrong: %v", err)
		}
		return models.Conversation{}, err
	}
	var conversation models.Conversation
	convoSnapshot.DataTo(&conversation)

	return conversation, nil
}

func GetMessages(msgIds []string) ([]map[string]any, error) {
	var messages []map[string]any
	c := context.Background()

	for _, msgId := range msgIds {
		msg, err := loadOneMessage(msgId, c)
		if err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}
	return messages, nil
}

func ClearConversation(convoId string) error {
	c := context.Background()
	convoRef := getConvoRef(convoId)

	convoSnapshot, err := convoRef.Get(c)
	if err != nil {
		log.Printf("Failed to load convo: %v", err)
		return err
	}

	msgIdsInterface, err := convoSnapshot.DataAt("messages")
	log.Print(reflect.TypeOf(msgIdsInterface))
	if err != nil {
		log.Printf("Failed to load convo: %v", err)
		return err
	}
	msgIds, ok := msgIdsInterface.([]interface{})
	if !ok {
		log.Printf("interface not a slice of strings: %v", err)
		return err
	}

	for _, msgId := range msgIds {
		err = deleteOneMessage(msgId.(string), c)
		if err != nil {
			log.Printf("Failed to delete %v: %v", msgId, err)
			return err
		}
	}

	_, err = convoRef.Delete(c)
	if err != nil {
		return err
	} else {
		return nil
	}
}

// helper
func getConvoRef(convoId string) *firestore.DocumentRef {
	conversations := Client.Collection("conversations")
	convoRef := conversations.Doc(convoId)
	return convoRef
}

func loadOneMessage(msgId string, c context.Context) (map[string]any, error) {
	msgRef := Client.Collection("messages").Doc(msgId)
	msgSnapshot, err := msgRef.Get(c)
	if err != nil {
		return nil, err
	}
	var msgMap models.ConvoMessage

	msgSnapshot.DataTo(&msgMap)
	returnMsg := map[string]any{
		"role":      msgMap.Role,
		"content":   msgMap.Content,
		"timestamp": msgMap.Timestamp,
	}

	return returnMsg, nil
}

func deleteOneMessage(msgId string, c context.Context) error {
	msgRef := Client.Collection("messages").Doc(msgId)
	_, err := msgRef.Delete(c)
	if err != nil {
		return err
	} else {
		return nil
	}
}
