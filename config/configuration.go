package config

type Application struct {
	Port string
}

func Setup() *Application {
	return &Application{
		Port: "8080",
	}
}
