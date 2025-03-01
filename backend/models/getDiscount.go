package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Schedule struct {
    ID primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
    BusStop string `json:"busStop"`
    Direction string `json:"direction"`
    DepartureTime string `json:"departureTime"`
}
