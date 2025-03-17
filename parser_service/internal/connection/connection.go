package connection

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"parser_service/internal/findselection"

	api "github.com/lum0vi/api_gkl"
)

type Connect interface {
	Connection() (resp *http.Response, err error)
}

type Site struct {
	Url      string
	Bytebody []byte
}

func (s *Site) Connection() (resp *http.Response) {
	body, err := http.Get(s.Url)
	if err != nil {
		l := log.Default()
		l.SetOutput(os.Stdout)
		l.Printf("Error connection: %s", err.Error())
	}
	return body
}

func (s *Site) GetBody() {
	body := s.Connection()
	bytebody, err := io.ReadAll(body.Body)
	if err != nil {
		l := log.Default()
		l.SetOutput(os.Stdout)
		l.Printf("Error connection: %s", err.Error())
	}
	s.Bytebody = bytebody
}

func Connection(ctx context.Context, url string, selection []string) ([]*api.SelectionElement, int) {
	site := Site{Url: url} // https://megakino42.ru/?facility=yubileynyy
	body := site.Connection()
	mas, len := findselection.Findselection(body.Body, selection) // "div.EventList_event__OjvqQ"
	return mas, len
	// jsongdata, err := json.Marshal(mas)
	// if err != nil {
	// 	l := log.Default()
	// 	l.SetOutput(os.Stdout)
	// 	l.Printf("Erro coding to json: %s", err.Error())
	// }
	// fmt.Println(string(jsongdata))
}
