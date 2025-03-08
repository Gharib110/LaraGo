package middleware

import (
	"myapp/data"

	"github.com/Gharib110/LaraGo"
)

type Middleware struct {
	App    *Lara.Lara
	Models data.Models
}
