package main

import (
	"log"
	"github.com/nats-io/nats.go"
)

func Subscriber() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Подписка с привязкой к очереди "workers"
	_, err = nc.QueueSubscribe("updates", "workers", func(m *nats.Msg) {
		log.Printf("Рабочий получил: %s", string(m.Data))
	})
	if err != nil {
		log.Fatal(err)
	}

	// Ожидаем сообщений, не даём завершиться приложению
	select {}
}