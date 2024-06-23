package main

import (
	"echo-templ/view"
	"echo-templ/view/component"
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

var palette = component.Base16{
    Scheme: "Campfire",
    Author: "Me",
    Base00: "#14171F",
    Base01: "#323848",
    Base02: "#3F475A",
    Base03: "#6D7A88",
    Base04: "#97A4AF",
    Base05: "#EFC164",
    Base06: "#2A2F3C",
    Base07: "#DDD7CA",
    Base08: "#A885C1",
    Base09: "#F35955",
    Base0A: "#F3835D",
    Base0B: "#468966",
    Base0C: "#3A8098",
    Base0D: "#70ADC2",
    Base0E: "#67CC8E",
    Base0F: "#DDD7CA",
}

func main() {
    e := echo.New()

    e.GET("/", HomeHandler)

    e.Logger.Fatal(e.Start(":3000"))
}

func Render(ctx echo.Context, statusCode int, t templ.Component) error {
    buf := templ.GetBuffer()
    defer templ.ReleaseBuffer(buf)

    if err := t.Render(ctx.Request().Context(), buf); err != nil {
        return err
    }

    return ctx.HTML(statusCode, buf.String())
}

func HomeHandler(c echo.Context) error {
    return Render(c, http.StatusOK, view.Index(palette))
}
