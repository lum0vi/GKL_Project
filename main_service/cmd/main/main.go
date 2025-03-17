package main

import (
	"context"
	"fmt"
	"io"
	"main_sevice/internal/logger"
	"net/http"
	"sync"
	"text/template"

	"github.com/labstack/echo/v4"
	api "github.com/lum0vi/api_gkl"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var (
	data      chan int                           = make(chan int)
	requestt  chan Requests                      = make(chan Requests)
	send_data chan api.GetParserElementsResponse = make(chan api.GetParserElementsResponse)
	mx        sync.RWMutex
)

type Requests struct {
	SiteUrl   string
	Selection []string
}

// Определяем структуру для регистрации шаблонов
type TemplateRegistry struct {
	templates *template.Template
}

// Реализуем интерфейс echo.Renderer
func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func GetSelData(c echo.Context) error {
	fmt.Println("yes")
	var StringParamSelec []string = []string{"div.EventList_event__OjvqQ", "img.Image_image__vhbZk",
		".Tags_age__6uF5I", ".Show_price__YStM_ price", "h2", ".Show_show-time__iv3r5"}
	data <- 1
	requestt <- Requests{SiteUrl: "https://megakino42.ru/?facility=yubileynyy", Selection: StringParamSelec}
	fmt.Println("2")
	res := <-send_data
	fmt.Println("3")
	c.Render(http.StatusOK, "index.html", nil)
	return c.String(http.StatusOK, fmt.Sprint(res))
}

func Middleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println(c.Path())
		return nil
	}
}

func main() {
	ctx := context.Background()
	ctx, _ = logger.New(ctx)
	e := echo.New()
	e.GET("/get", GetSelData)
	grpcConn, err := grpc.NewClient("localhost:50051", grpc.WithInsecure())
	if err != nil {
		logger.GetLoggerFromCtx(ctx).Info(ctx, "failed to connect grpc service", zap.Error(err))
	}
	mainservice := api.NewParserClient(grpcConn)
	e.Renderer = &TemplateRegistry{
		templates: template.Must(template.ParseGlob("web/templates/*.html")),
	}
	defer grpcConn.Close()
	go func() {
		for {
			mod := <-data
			fmt.Println("4")
			if mod == 1 {
				req := api.GetParserElementsRequest{}
				req_p := <-requestt
				req.SiteUrl = req_p.SiteUrl
				req.Selection = req_p.Selection
				resp, err := mainservice.GetParserElements(ctx, &req)
				fmt.Println("5")
				if err != nil {
					logger.GetLoggerFromCtx(ctx).Info(ctx, "error with using service method", zap.Error(err))
				} else {
					send_data <- *resp
				}
			}
		}
	}()
	e.Logger.Fatal(e.Start(":1323"))
}
