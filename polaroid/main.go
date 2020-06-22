package main

import (
	"fmt"
	"polaroid/admin"
	"polaroid/config"
	"polaroid/server"
)

func main() {
	conf := config.NewConf()
	conf["pk"] = config.Pk()
	srv := server.NewServer(conf)
	adm := admin.NewAdmin(srv.Data)
	srv.AddHH(adm)
	serve := srv.Servers(conf)
	err := serve.ListenAndServe()
	fmt.Println(err)
}
