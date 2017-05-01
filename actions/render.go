package actions

import (
	"fmt"
	"math"
	"path/filepath"
	"time"

	"github.com/bketelsen/learnonline/models"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/packr"
	"github.com/leekchan/accounting"
)

var r *render.Engine
var money = accounting.Accounting{Symbol: "$", Precision: 2}

func init() {
	r = render.New(render.Options{
		// HTML layout to be used for all HTML requests:
		HTMLLayout: "application.html",

		// Box containing all of the templates:
		TemplatesBox: packr.NewBox("../templates"),

		// Add template helpers here:
		Helpers: render.Helpers{
			"inverted":  invertedHelper,
			"currency":  currencyHelper,
			"timestamp": timestamp,
		},
	})
}
func getSite(c buffalo.Context) *models.Site {

	fmt.Println("pulling site", c.Value("site"))
	site := c.Value("site").(*models.Site)
	fmt.Println(site)
	c.Set("templatebase", site.Basedir)
	return site
}

func sitePath(c buffalo.Context, path string) string {
	site := getSite(c)
	return (filepath.Join(site.Basedir, path))
}

func currencyHelper(price int) string {
	return money.FormatMoney(price / 100)
}
func invertedHelper(index int) string {
	if index%2 == 0 {
		return "timeline-inverted"
	}
	return ""
}

func timestamp(timestamp int64) string {
	timestamp = int64(math.Abs(float64(timestamp) / 1000.0))
	return time.Unix(timestamp, 0).Format("January 2, 2006")
}
