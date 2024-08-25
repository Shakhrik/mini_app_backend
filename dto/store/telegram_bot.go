package store

import "time"

type TelegramBotCreateParams struct {
	ID          string
	Name        string
	Token       string
	Description string
	ImageUrl    string
}

type TelegramBot struct {
	ID          string    `db:"id"`
	Name        string    `db:"name"`
	Token       string    `db:"token"`
	Description string    `db:"description"`
	ImageUrl    string    `db:"image_url"`
	CreateAt    time.Time `db:"created_at"`
}
