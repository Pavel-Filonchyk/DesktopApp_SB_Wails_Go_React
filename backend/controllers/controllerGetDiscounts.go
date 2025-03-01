package controllers

import (
    "context"
    "encoding/json"
    "net/http"

    "go.mongodb.org/mongo-driver/bson"
    "Sushi/backend/db"
    "Sushi/backend/models"
)

func GetDiscounts(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    collection := db.GetCollection("schedules")
    
    cursor, err := collection.Find(context.TODO(), bson.M{})
    if err != nil {
        http.Error(w, "Failed to fetch schedules", http.StatusInternalServerError)
        return
    }
    defer cursor.Close(context.TODO())

    var schedules []models.Schedule
    if err := cursor.All(context.TODO(), &schedules); err != nil {
        http.Error(w, "Error decoding schedules", http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(schedules)
}
