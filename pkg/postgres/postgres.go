package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Config struct { // Конфигурация для подключения к базе данных PostgreSQL
	User     string `envconfig:"POSTGRES_USER"     required:"true"`
	Password string `envconfig:"POSTGRES_PASSWORD" required:"true"`
	Port     string `envconfig:"POSTGRES_PORT"     required:"true"`
	Host     string `envconfig:"POSTGRES_HOST"     required:"true"`
	DBName   string `envconfig:"POSTGRES_DB"       required:"true"`
}

type Pool struct { // Структура для хранения пула соединений с базой данных PostgreSQL, которая будет использоваться для выполнения запросов к базе данных
	*pgxpool.Pool
}

func New(ctx context.Context, c Config) (*Pool, error) { // Конструктор для создания нового пула соединений с базой данных
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", // Формируем строку подключения (DSN)
		c.User, c.Password, c.Host, c.Port, c.DBName)

	cfg, err := pgxpool.ParseConfig(dsn) // Парсим строку подключения в конфигурацию для pgxpool
	if err != nil {
		return nil, fmt.Errorf("pgxpool.ParseConfig: %w", err)
	}

	pool, err := pgxpool.NewWithConfig(ctx, cfg) // Создаем новый пул соединений с базой данных, используя конфигурацию
	if err != nil {
		return nil, fmt.Errorf("pgxpool.NewWithConfig: %w", err)
	}

	return &Pool{pool}, nil
}

func (p *Pool) Close() { // Метод для закрытия пула соединений с базой данных
	p.Pool.Close()
}
