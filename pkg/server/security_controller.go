package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"golang.org/x/crypto/bcrypt"
)

func firewall(c *fiber.Ctx) error {
	sess, err := store.Get(c)
	if err != nil {
		log.Error("unable to open session store ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	logged_in := sess.Get("logged_in")

	if logged_in == true {
		uid := sess.Get("userid").(int64)
		user, err := q.GetUserById(ctx, uid)
		if err == nil {
			c.Locals("User", user)
			return c.Next()
		}
	}

	sess.Set("logged_in", false)

	return c.RedirectToRoute("login", fiber.Map{})
}

func GetLogin(c *fiber.Ctx) error {
	log.Info("login route")
	return c.Render("login/login", fiber.Map{}, "")
}

func PostLogin(c *fiber.Ctx) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	user, err := q.GetUserByEmail(ctx, username)

	if err == nil {
		if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) == nil {
			sess, err := store.Get(c)
			if err != nil {
				log.Error("Error opening store", err)
				return c.SendStatus(fiber.StatusInternalServerError)
			}

			sess.Set("logged_in", true)
			sess.Set("userid", user.ID)
			sess.Save()
			if c.FormValue("redirectUrl") != "" {
				return c.Redirect(c.FormValue("redirectUrl"))
			} else {
				return c.Redirect("/")
			}
		}

		pwd, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
		log.Warn("Login attempt for user", username, "with invalid password", password, "expected:", string(pwd))
	}
	log.Warn("Login attempt for invalid user", username, "password", password)

	return c.Render("login/login", fiber.Map{
		"error":       "Login failure",
		"redirectUrl": c.FormValue("redirectUrl"),
	})
}

func GetLogout(c *fiber.Ctx) error {
	sess, err := store.Get(c)
	if err != nil {
		log.Error("unable to open session store ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	sess.Set("logged_in", false)
	return c.Redirect("login")
}
