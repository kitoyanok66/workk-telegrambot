package handlers

import (
	tb "github.com/mymmrac/telego"
)

type StartHandler struct {
	Bot     *tb.Bot
	Backend *service.AuthService
}
