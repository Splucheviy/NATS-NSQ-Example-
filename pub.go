package main

import (
    "log"
    "github.com/nats-io/nats.go"
)

func Publish() {
    // Подключаемся к серверу
    nc, err := nats.Connect(nats.DefaultURL) // nats://127.0.0.1:4222
    if err != nil {
        log.Fatal(err)
    }
    defer nc.Close()

    // Publish - отправляем сообщение в топик "updates"
    err = nc.Publish("updates", []byte("Привет! Это сообщение от сервиса А."))
    if err != nil {
        log.Fatal(err)
    }

    // Обеспечиваем flush, чтобы гарантировать доставку
    nc.Flush()
}