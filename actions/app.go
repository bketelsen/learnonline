package actions

import (
	"fmt"
	"time"

	"github.com/bketelsen/learnonline/models"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/middleware"
	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/packr"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/pop"
)

// ENV is used to help switch settings based on where the
// application is being run. Default is "development".
var ENV = envy.Get("GO_ENV", "development")
var app *buffalo.App

// Sites is the list of sites this app can serve
var Sites map[string]*models.Site
var defaultSite = "training"

// App is where all routes and middleware for buffalo
// should be defined. This is the nerve center of your
// application.
func App() *buffalo.App {
	if app == nil {
		app = buffalo.Automatic(buffalo.Options{
			Env:         ENV,
			SessionName: "_teachweb_session",
		})

		app.Use(setSite)
		app.Use(middleware.PopTransaction(models.DB))
		app.Use(setCurrentUser)
		app.Use(trackLastURL)
		app.Use(setStripeKeys)

		app.Use(getCourses)

		app.Use(setTitle)
		app.Use(func(next buffalo.Handler) buffalo.Handler {
			return func(c buffalo.Context) error {
				c.Set("year", time.Now().Year())
				return next(c)
			}
		})

		app.GET("/", HomeHandler)
		app.GET("/training/online", CoursesIndex)
		app.GET("/training/live", WebinarsIndex)
		app.GET("/blog", BlogIndex)
		app.GET("/blog/{slug}", BlogShow)
		app.GET("/training/online/{slug}", CoursesShow)
		app.GET("/training/live/{slug}", WebinarsShow)
		app.POST("/training/online/purchase/{course_id}", authorize(PurchasesCreate))
		app.GET("/classroom/{slug}", authorize(ClassroomShow))
		app.GET("/classroom/{slug}/{module}", authorize(ClassroomModuleShow))
		app.ServeFiles("/assets", packr.NewBox("../public/assets"))

		auth := app.Group("/auth")
		auth.Middleware.Replace(trackLastURL, func(next buffalo.Handler) buffalo.Handler {
			return func(c buffalo.Context) error {
				return next(c)
			}
		})
		auth.GET("/{provider}", buffalo.WrapHandlerFunc(gothic.BeginAuthHandler))
		auth.GET("/{provider}/callback", AuthCallback)
		app.GET("/logout", AuthLogout)
	}

	return app
}

func setSite(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		req := c.Request()
		host := req.Host
		c.LogField("Host", host)
		set := false
		for h, s := range Sites {
			fmt.Printf("comparing %s to %s", host, h)
			if s.Baseurl == host {
				fmt.Println("!! Setting HOST", h, s)
				c.Set("site", s)
				set = true

				r.HTMLLayout = s.Basedir + "/" + "application.html"
			}
		}
		if !set {
			c.Set("site", defaultSite)
		}
		return next(c)
	}
}

func trackLastURL(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		req := c.Request()
		if req.Method == "GET" {
			c.Session().Set("last_url", req.URL.Path)
			err := c.Session().Save()
			if err != nil {
				return err
			}
		}
		return next(c)
	}
}

func setTitle(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		if c.Value("title") == nil {
			c.Set("title", "Classroom")
		}

		return next(c)
	}
}
func getCourses(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		tx := c.Value("tx").(*pop.Connection)
		courses := models.GetCourses()
		if c.Value("current_user") != nil {
			cu := c.Value("current_user").(*models.User)
			courses.MarkPurchases(tx, cu)
		}
		c.Set("courses", courses)

		return next(c)
	}
}
