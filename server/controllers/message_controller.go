package controllers

import (
	"context"
	"net/http"
	"server/configs"
	"server/models"
	"server/responses"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var messageCollections *mongo.Collection = configs.GetCollection(configs.DB, "MessageManagement")
var messageValidate = validator.New()

func SendMessage(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var message models.Message
	defer cancel()

	if err := c.BodyParser(&message); err != nil {
		return c.JSON(responses.MessageResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	if validationErr := messageValidate.Struct(&message); validationErr != nil {
		return c.JSON(responses.MessageResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	newMessage := models.Message{
		Id:             primitive.NewObjectID(),
		SenderID:       message.SenderID,
		ReceiverID:     message.ReceiverID,
		MessageContent: message.MessageContent,
		Time:           message.Time,
	}

	result, err := messageCollections.InsertOne(ctx, newMessage)

	if err != nil {
		return c.JSON(responses.MessageResponse{Status: http.StatusInternalServerError, Message: "Failed to send new message", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.JSON(responses.MessageResponse{Status: http.StatusOK, Message: "Success!", Data: &fiber.Map{"data": result}})
}
