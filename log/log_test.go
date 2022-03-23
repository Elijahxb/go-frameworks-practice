package log

import (
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"go.uber.org/zap"
)

func init() {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()
	sugar.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", "https://www.baidu.com",
		"attempt", 3,
		"backoff", time.Second,
	)
}

type Book struct {
	Title string `json:"title"`
}

type BookController struct {
	/* dependencies */
}

// GET: http://localhost:8080/api/v2/books
func (c *BookController) Get() []Book {
	//sugar.Info("mvc收到get请求")
	return []Book{
		{"Mastering Concurrency in Go"},
		{"Go Design Patterns"},
		{"Black Hat Go"},
	}
}

// POST: http://localhost:8080/api/v2/books
func (c *BookController) Post(b Book) int {
	//sugar.Info("mvc收到post请求")
	println("Received Book: " + b.Title)

	return iris.StatusCreated
}

// GET: http://localhost:8080/api/v1/books
func list(ctx iris.Context) {
	//sugar.Info("收到get请求")
	books := []Book{
		{"Mastering concurrency in Python"},
		{"Python Design Patterns"},
		{"Black Hat Python"},
	}
	ctx.JSON(books)
}

// POST: http://localhost:8080/api/v1/books
func create(ctx iris.Context) {
	//sugar.Info("收到post请求")
	var b Book
	err := ctx.ReadJSON(&b)
	// TIP: use ctx.ReadBody(&b) to bind
	// any type of incoming data instead.
	if err != nil {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
			Title("Book creation failure").DetailErr(err))
		// TIP: use ctx.StopWithError(code, err) when only
		// plain text responses are expected on errors.
		return
	}

	println("Received Book: " + b.Title)

	ctx.StatusCode(iris.StatusCreated)
}

func log_test() {
	app := iris.New()
	booksAPIV1 := app.Party("/api/v1/books")
	{
		booksAPIV1.Use(iris.Compression)
		booksAPIV1.Get("/", list)
		booksAPIV1.Post("/", create)
	}
	booksAPIV2 := app.Party("/api/v2/books")
	{
		booksAPIV2.Use(iris.Compression)
	}
	var m = mvc.New(booksAPIV2)
	m.Handle(new(BookController))
	app.Listen(":8080")
	//sugar.Info("app结束")
}
