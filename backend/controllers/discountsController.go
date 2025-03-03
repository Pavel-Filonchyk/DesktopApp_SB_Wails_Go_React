package controllers

import (
	"context"
	"encoding/json"

	"Sushi/backend/db"
	"Sushi/backend/models"

	"go.mongodb.org/mongo-driver/bson"
)

// Функция для получения всех даных напрямую без http.ResponseWriter
func GetDiscountsDirect() string {
	collection := db.GetCollection("schedules")

	// Запрос к базе
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return `[]` 
	}
	defer cursor.Close(context.TODO())

	// Необходимо распарсить данные
	var schedules []models.Schedule
	if err := cursor.All(context.TODO(), &schedules); err != nil {
		return `[]` 
	}

	data, _ := json.Marshal(schedules)
	return string(data)
}
