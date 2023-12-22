package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/gorilla/csrf"
	"github.com/joho/godotenv"
	"github.com/psanti93/galleryValleyv1/controllers"
	"github.com/psanti93/galleryValleyv1/migrations"
	"github.com/psanti93/galleryValleyv1/models"
	"github.com/psanti93/galleryValleyv1/templates"
	"github.com/psanti93/galleryValleyv1/views"
)

type config struct {
	PSQL models.PostgresConfig
	SMTP models.SMTConfig
	CSRF struct {
		Key    string
		Secure bool
	}
	Server struct {
		Address string // "localhost:3000"
	}
}

func loadEnvConfig() (config, error) {
	var cfg config
	// load the env file with the smtp
	err := godotenv.Load()
	if err != nil {
		return cfg, err
	}
	// TODO: psql set up read from environment varibles
	cfg.PSQL = models.DefaultPostgresConfig()

	//TODO: smtp
	cfg.SMTP.Host = os.Getenv("SMTP_HOST")
	portStr := os.Getenv("SMTP_PORT")
	cfg.SMTP.Port, err = strconv.Atoi(portStr)
	if err != nil {
		return cfg, err
	}
	cfg.SMTP.Username = os.Getenv("SMTP_USERNAME")
	cfg.SMTP.Password = os.Getenv("SMTP_PASSWORD")

	// TODO: CSRF read from env variables
	cfg.CSRF.Key = "gFvi45R4fy5xNBlnEeZtQbfAVCYEIAUX"
	cfg.CSRF.Secure = false

	// TODO address read from an environment variable
	cfg.Server.Address = ":3000"
	return cfg, nil
}

func main() {

	cfg, err := loadEnvConfig()
	if err != nil {
		panic(err)
	}

	//Set Up database:

	db, err := models.Open(cfg.PSQL)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	// running the migrations
	err = models.MigrateFS(db, migrations.EmbedMigratonFS, ".")
	if err != nil {
		panic(err)
	}

	// Set Up services

	userService := &models.UserService{
		DB: db}

	sessionService := &models.SessionService{
		DB: db,
	}
	passwordResetService := &models.PasswordResetService{
		DB: db,
	}

	emailService := models.NewEmailService(cfg.SMTP)

	//Set up middleware
	umw := controllers.UserMiddleware{
		SessionService: sessionService,
	}

	csrfMw := csrf.Protect([]byte(cfg.CSRF.Key), csrf.Secure(cfg.CSRF.Secure)) // csrf.Secure() by default it's true, it requires an https secure connection, false for now cause local we don't have https connection. set to true in prod

	// set up controller:
	usersC := controllers.Users{
		UserService:          userService,
		SessionService:       sessionService,
		PasswordResetService: passwordResetService,
		EmailService:         emailService,
	}

	//set up router and routes

	r := chi.NewRouter()
	// setting up the middleware in the router
	r.Use(csrfMw)
	r.Use(umw.SetUser)

	// parsing the template prior to executing it
	tpl := views.Must(views.ParseFS(templates.FS, "home.gohtml", "tailwind.gohtml"))
	r.Get("/", controllers.StaticHandler(tpl))
	//Contact
	tpl = views.Must(views.ParseFS(templates.FS, "contact.gohtml", "tailwind.gohtml"))
	r.Get("/contact", controllers.StaticHandler(tpl))
	// FAQ
	tpl = views.Must(views.ParseFS(templates.FS, "faq.gohtml", "tailwind.gohtml"))
	r.Get("/faq", controllers.FAQ(tpl))

	// Sign Up Page Routes
	usersC.Templates.SignUp = views.Must(views.ParseFS(templates.FS, "signup.gohtml", "tailwind.gohtml"))
	r.Get("/signup", usersC.SignUp)
	r.Post("/users", usersC.CreateUser)

	// Sign In Page Routes
	usersC.Templates.SignIn = views.Must(views.ParseFS(templates.FS, "signin.gohtml", "tailwind.gohtml"))
	r.Get("/signin", usersC.SignIn)
	r.Post("/signin", usersC.ProcessSignIn)

	// Reset Password Route
	usersC.Templates.ForgotPasswordTemplate = views.Must(views.ParseFS(templates.FS, "forgot-pw.gohtml", "tailwind.gohtml"))
	r.Get("/forgot-pw", usersC.ForgotPassword)
	r.Post("/forgot-pw", usersC.ProcessForgotPassword)

	// Checking our email
	usersC.Templates.CheckYoureEmail = views.Must(views.ParseFS(templates.FS, "check-your-email.gohtml", "tailwind.gohtml"))

	// getting the cookie of the user
	// r.Get("/users/me", usersC.CurrentUser)

	//anything that has the prefix "/users/me"
	r.Route("/users/me", func(r chi.Router) {
		r.Use(umw.RequireUser)
		r.Get("/", usersC.CurrentUser)
		r.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "Hello!")
		})

	})

	//Add sign outrout
	r.Post("/signout", usersC.ProcessSignOut)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	fmt.Printf("Starting server on port %s ....\n ", cfg.Server.Address)

	err = http.ListenAndServe(cfg.Server.Address, r)

	if err != nil {
		panic(err)
	}

}
