package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// GetMenu возвращает меню
func GetMenu() string {
	fmt.Println("Функция GetMenu() вызвана!") 

	data, err := ioutil.ReadFile("menu.json") 
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return "[]"
	}

	var menu []map[string]interface{}
	err = json.Unmarshal(data, &menu)
	if err != nil {
		fmt.Println("Ошибка парсинга JSON:", err)
		return "[]"
	}

	result, _ := json.Marshal(menu)
	return string(result)
}
