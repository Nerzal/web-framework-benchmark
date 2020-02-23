package router

import (
	"fmt"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
	"github.com/savsgio/atreugo/v10"
	"github.com/valyala/fasthttp"
)

func init() {
	r := echo.New()
	r.HideBanner = true

	r.GET("/hello", func(c echo.Context) error {
		return c.JSON(200, "hello")
	})

	r.GET("/hello/:id/world", func(c echo.Context) error {
		return c.JSON(200, "hello "+c.Param("id")+" world")
	})

	r.GET("/static", func(c echo.Context) error {
		return c.File("README.md")
	})

	go r.Start(":1338")

	config := &atreugo.Config{
		Addr: "0.0.0.0:1339",
	}

	a := atreugo.New(config)
	a.GET("/hello", func(ctx *atreugo.RequestCtx) error {
		return ctx.TextResponse("hello", 200)
	})

	a.GET("/static", func(ctx *atreugo.RequestCtx) error {
		return ctx.Response.SendFile("README.md")
	})

	a.GET("/hello/:id/world", func(ctx *atreugo.RequestCtx) error {
		return ctx.TextResponse(fmt.Sprintf("hello %s world", ctx.UserValue("id")), 200)
	})

	go a.ListenAndServe()

	gin.SetMode(gin.ReleaseMode)
	g := gin.New()
	g.GET("/hello", func(c *gin.Context) {
		c.String(200, "hello")
	})

	g.GET("/static", func(c *gin.Context) {
		c.File("README.md")
	})

	g.GET("/hello/:id/world", func(c *gin.Context) {
		c.String(200, "hello "+c.Param("id")+" world")
	})

	go g.Run("0.0.0.0:1337")
}

func Benchmark_Atreugo_Hello(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		status, _, err := fasthttp.Get(nil, "http://localhost:1339/hello")
		if err != nil {
			b.Error("failed to call hello: ", err)
		}

		if status != 200 {
			b.Error("received wrong statuscode")
		}
	}
}

func Benchmark_Echo_Hello(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		status, _, err := fasthttp.Get(nil, "http://localhost:1338/hello")
		if err != nil {
			b.Error("failed to call hello: ", err)
		}

		if status != 200 {
			b.Error("received wrong statuscode")
		}
	}
}

func Benchmark_Gin_Hello(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		status, _, err := fasthttp.Get(nil, "http://localhost:1337/hello")
		if err != nil {
			b.Error("failed to call hello: ", err)
		}

		if status != 200 {
			b.Error("received wrong statuscode")
		}
	}
}

func Benchmark_Atreugo_Hello_PathParam(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		status, _, err := fasthttp.Get(nil, "http://localhost:1339/hello/42/world")
		if err != nil {
			b.Error("failed to call hello: ", err)
		}

		if status != 200 {
			b.Error("received wrong statuscode")
		}
	}
}

func Benchmark_Echo_Hello_PathParam(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		status, _, err := fasthttp.Get(nil, "http://localhost:1338/hello/42/world")
		if err != nil {
			b.Error("failed to call hello: ", err)
		}

		if status != 200 {
			b.Error("received wrong statuscode")
		}
	}
}

func Benchmark_Gin_Hello_PathParam(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		status, _, err := fasthttp.Get(nil, "http://localhost:1337/hello/42/world")
		if err != nil {
			b.Error("failed to call hello: ", err)
		}

		if status != 200 {
			b.Error("received wrong statuscode")
		}
	}
}

func Benchmark_Atreugo_Static(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		status, _, err := fasthttp.Get(nil, "http://localhost:1339/static")
		if err != nil {
			b.Error("failed to call hello: ", err)
		}

		if status != 200 {
			b.Error("received wrong statuscode")
		}
	}
}

func Benchmark_Echo_Static(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		status, _, err := fasthttp.Get(nil, "http://localhost:1338/static")
		if err != nil {
			b.Error("failed to call hello: ", err)
		}

		if status != 200 {
			b.Error("received wrong statuscode")
		}
	}
}

func Benchmark_Gin_Static(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		status, _, err := fasthttp.Get(nil, "http://localhost:1337/static")
		if err != nil {
			b.Error("failed to call hello: ", err)
		}

		if status != 200 {
			b.Error("received wrong statuscode")
		}
	}
}
