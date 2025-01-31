package lara

import "net/http"

func (l *Lara) SessionLoad(next http.Handler) http.Handler {
	return l.Session.LoadAndSave(next)
}
