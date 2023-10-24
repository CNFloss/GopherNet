package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/CNFloss/GopherNet/api/data"
)

type Users struct {
	logger *log.Logger
	dataSource data.DataSource
}

func NewUsers(ds data.DataSource, l *log.Logger) *Users {
	return &Users{dataSource: ds, logger:l}
}

func (u *Users) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	u.logger.Println("Hello Users")

	if r.Method == http.MethodPost {
		d, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "Oops", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(rw, "%s\n", d)
	}
}