package admin

import (
	"net/http"
	"polaroid/server"
	"polaroid/tool"
	"polaroid/types"
)

type Admin struct {
	Data *types.Data
}

func NewAdmin(s *types.Data) (a *Admin) {
	a = new(Admin)
	a.Data = s
	return
}

func Admins(e *types.Data) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		head := tool.NewHeader(r, w, "gopiko-admin", e)
		head.AddCss("oui")
		head.Jointure("layout.html", "admin.html")
	})
}

func (a *Admin) WWW(s *server.Server) {
	s.NewR("/admin", "admin", []string{"GET"}, Admins(s.Data), 1)
}
