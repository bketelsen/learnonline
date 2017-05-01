package actions

import (
	"fmt"

	"github.com/bketelsen/learnonline/models"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/envy"
	"github.com/markbates/pop"
	"github.com/satori/go.uuid"
	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
)

func init() {
	stripe.Key = envy.Get("STRIPE_SECRET_KEY", "sk_test_CWb5jl75dB7SBNJJKC15ySmi")
}

func PurchasesCreate(c buffalo.Context) error {

	// Token is created using Stripe.js or Checkout!
	// Get the payment token submitted by the form:
	token := c.Request().FormValue("stripeToken")

	tx := c.Value("tx").(*pop.Connection)
	id, err := c.ParamInt("course_id")
	if err != nil {
		return err
	}
	course, err := models.GetCourse(id)
	if err != nil {
		return err
	}

	p := &models.Purchase{
		CourseID: course.ID,
		UserID:   c.Value("current_user_id").(uuid.UUID),
	}

	b, err := tx.Where("course_id = ? and user_id = ?", p.CourseID, p.UserID).Exists(p)
	if err != nil {

		c.Flash().Add("error", fmt.Sprintf("error: %s", err.Error()))
		fmt.Println(err)
		return err
	}
	if b {
		c.Flash().Add("info", fmt.Sprintf("You have already purchased %s.", course.Title))
		return c.Redirect(302, fmt.Sprintf("/classroom/%s", course.Slug))
	}

	// Charge the user's card:
	params := &stripe.ChargeParams{
		Amount:   uint64(course.Price),
		Currency: "usd",
		Desc:     fmt.Sprintf("GopherAcademy - %s", course.Title),
	}
	params.SetSource(token)

	_, err = charge.New(params)
	if err != nil {
		c.Flash().Add("danger", fmt.Sprintf("There was a problem creating the charge: %v", err))
		return c.Redirect(302, "/courses")
	}

	err = tx.Create(p)
	if err != nil {
		c.Flash().Add("danger", fmt.Sprintf("There was a problem creating the purchase: %v", err))
		return c.Redirect(302, "/courses")
	}
	c.Flash().Add("success", fmt.Sprintf("You have successfully purchased %s!", course.Title))
	return c.Redirect(302, fmt.Sprintf("/classroom/%s", course.Slug))
}

func setStripeKeys(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		c.Set("stripe_pk", envy.Get("STRIPE_PUBLIC_KEY", "pk_test_niJp8y50uDCfVLk6AywgAk5B"))
		return next(c)
	}
}
