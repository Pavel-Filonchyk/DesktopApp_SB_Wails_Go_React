package main

import (
	"context"
	"Sushi/backend/controllers"
	"Sushi/backend/db"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// сохраняет контекст при запуске приложения
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) GetAllMenu() string {
	return controllers.GetMenu()
}

func (a *App) GetAllDiscounts() string {
	db.ConnectDB()
	return controllers.GetDiscountsDirect()
}

func (a *App) GetAllBasket() string {
	return controllers.GetBasket()
}

func (a *App) SendAllBasket(basketJson string) {
	controllers.SendBasket(basketJson)
}