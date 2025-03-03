package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"Sushi/backend/models"
)

type Basket []models.BasketItem

func SendBasket(basketJson string) {
	var basket Basket

	err := json.Unmarshal([]byte(basketJson), &basket)
	if err != nil {
		log.Println("Ошибка при парсинге входных данных корзины:", err)
		return
	}

	data, err := json.MarshalIndent(basket, "", "  ")
	if err != nil {
		log.Println("Ошибка при сериализации корзины:", err)
		return
	}

	err = ioutil.WriteFile("basket.json", data, 0644)
	if err != nil {
		log.Println("Ошибка при сохранении корзины:", err)
	}
}
