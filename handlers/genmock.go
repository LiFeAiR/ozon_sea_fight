package handlers

import "github.com/LiFeAiR/ozon_sea_fight/system"

//go:generate ../bin/mockery --name=(.+)Mock --case=underscore

type AppMock interface {
	system.App
}
