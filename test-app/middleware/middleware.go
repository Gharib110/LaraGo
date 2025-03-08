package middleware

import (
	"test-app/data"

	"github.com/Gharib110/LaraGo"
)

type Middleware struct {
	App    *lara.Lara
	Models data.Models
}
