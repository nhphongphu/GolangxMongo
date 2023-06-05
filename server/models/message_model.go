package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Message struct {
	Id             primitive.ObjectID `json:"id,omitempty"`
	SenderID       string             `json:"sender_id,omitempty" validate:"required"`
	ReceiverID     string             `json:"receiver_id,omitempty" validate:"required"`
	MessageContent string             `json:"message_content,omitempty" validate:"required"`
	Time           string             `json:"time,omitempty" validate:"required"`
}
