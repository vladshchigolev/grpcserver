package store

type Config struct {
	DatabaseURL string `toml:"database_url"` // Строка подключения к базе
}

func NewConfig() *Config { // Функция-конструктор объекта, или helper-функция
	return &Config{}
}
