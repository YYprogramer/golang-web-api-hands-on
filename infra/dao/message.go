package dao

import (
	"math/rand"
	"time"

	"golang-web-api-hands-on/repository"
)

type messageRepository struct{}

func NewMessageRepository() repository.MessageRepository {
	return &messageRepository{}
}

func (m *messageRepository) GetMessage() string {
	messages := []string{
		"Change before you have to.",
		"There is always light behind the clouds.",
		"If you can dream it, you can do it.",
		"Love the life you live. Live the life you love.",
		"変わる前に変える。",
		"雲の後ろにはいつも光がある。",
		"夢見ることができれば、それを実現できる。",
		"生きる人生を愛し、愛する人生を生きる。",
	}

	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(messages))

	return messages[randomIndex]
}
