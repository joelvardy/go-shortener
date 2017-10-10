package main

import (
	"net/http"
	"github.com/bahlo/goat"
	"github.com/gorilla/handlers"
	"github.com/siddontang/go-mysql/client"
	"github.com/mattheath/base62"
	"fmt"
	"net"
	"os"
)

type App struct {
	Database *client.Conn
}

type Response struct {
	Error bool   `json:"error"`
	Code  string `json:"code,omitempty"`
}

func (app *App) createHandler(w http.ResponseWriter, r *http.Request, p goat.Params) {
	r.ParseForm()
	result, err := app.Database.Execute(fmt.Sprintf("insert into links (url) values ('%s')", r.Form.Get("url")))
	if err != nil {
		goat.WriteJSON(w, Response{
			Error: true,
		})
	} else {
		code := base62.EncodeInt64(int64(result.InsertId))
		goat.WriteJSON(w, Response{
			Error: false,
			Code:  code,
		})
	}
}

func (app *App) redirectHandler(w http.ResponseWriter, r *http.Request, p goat.Params) {
	id := base62.DecodeToBigInt(p["code"])
	result, _ := app.Database.Execute(fmt.Sprintf("select url from links where id = %d", id))
	url, _ := result.GetStringByName(0, "url")
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	app.Database.Execute(fmt.Sprintf("insert into link_clicks (link_id, ip_address) values (%d, '%s')", id, ip))
	http.Redirect(w, r, url, 301)
}

func main() {

	mysql, err := client.Connect(os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"))
	if err != nil {
		panic("Unable to connect to database")
	}

	app := App{
		Database: mysql,
	}

	server := goat.New()
	server.Post("/create", "create", app.createHandler)
	server.Get("/:code", "redirect", app.redirectHandler)
	server.Use(handlers.CORS())
	server.Run(":8080")

}
