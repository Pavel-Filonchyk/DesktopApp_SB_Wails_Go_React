package main

import (
	"context"
	"fmt"
	"encoding/json"
	"io/ioutil"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup сохраняет контекст при запуске приложения
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) GetMenu() string {
	fmt.Println("Функция GetMenu() вызвана!") 

	data, err := ioutil.ReadFile("menu.json") 
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return "[]"
	}

	// Определяю структуру для всего меню (не изменяем оригинальную структуру!)
	var menu []map[string]interface{}

	err = json.Unmarshal(data, &menu)
	if err != nil {
		fmt.Println("Ошибка парсинга JSON:", err)
		return "[]"
	}

	//fmt.Println("Отправляем в React:", menu)

	result, _ := json.Marshal(menu)
	return string(result)
}
