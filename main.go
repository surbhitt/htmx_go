package main

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Templates struct {
    templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Templates {
    return &Templates{
        templates: template.Must(template.ParseGlob("views/*.html")),
    }
}

type Contact struct {
    Name string
    Email string
}

func newContact(name, email string) Contact {
    return Contact {
        Name: name,
        Email: email,
    }
}

type Contacts = []Contact

type Data struct {
    Contacts Contacts
}

func newData() Data {
    return Data{
        Contacts: []Contact{
            newContact("surbhit", "surbhit@gmail.com"),
            newContact("tushar", "tushar@gmail.com"),
        },
    }
}

func main(){
    e := echo.New()
    e.Use(middleware.Logger())
            
    data := newData()
    // for _, x := range data.contacts {
    //     fmt.Printf("name: %s, email: %s\n", x.name, x.email)
    // }
    e.Renderer = newTemplate()
    e.GET("/", func(c echo.Context) error {
        return c.Render(200, "index", data)
    })
    e.POST("/contacts", func(c echo.Context) error {
        data.Contacts = append(data.Contacts, newContact(c.FormValue("name"), c.FormValue("email")))
        return c.Render(200, "display", data)
    })
    e.Logger.Fatal(e.Start(":8080"))
}
