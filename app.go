package main

import (
	"context"
	"Sushi/backend/controllers"

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

// Вызов маршрутов
func (a *App) GetAllMenu() string {
	return controllers.GetMenu()
}
func (a *App) GetAllDiscounts() string {
	//return controllers.GetDiscounts()
	return "err"
}
func (a *App) SendBasket() string {
	return controllers.GetMenu()
}