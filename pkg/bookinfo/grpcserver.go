package bookinfo

import (
	"awesomeProject/internal/app/store"
	"awesomeProject/pkg/api"
	"context"
	"github.com/sirupsen/logrus"
)

type GRPCServer struct {
	config *Config
	logger *logrus.Logger
	store  *store.Store
}

type Config struct { // Config GRPCServer'а
	BindAddr string        `toml:"bind_addr"`
	LogLevel string        `toml:"log_level"`
	Store    *store.Config // Config хранилища
}

func NewConfig() *Config { // Конструктор нового Config'а
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
		Store:    store.NewConfig(),
	}
}

func (s *GRPCServer) ConfigureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil
}

func (s *GRPCServer) ConfigureStore() error {
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}
	s.store = st
}
func (s *GRPCServer) GetBooks(ctx context.Context, auth *api.Author) (*api.Books, error) {
	// здесь мы должны взять значение поля Name переданного методу объекта типа Author и по этому значению искать в базе книги
	return &api.Books{Titles: auth.Name}, nil
}

func (s *GRPCServer) GetAuthors(ctx context.Context, book *api.Book) (*api.Authors, error) {
	return &api.Authors{Names: "Some Names"}, nil
}
