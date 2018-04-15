package fila

import (
  "time"
)

type connection struct {
  Created time.Time
  Url string
  Open bool
}

func (c *connection) Connect(url string) {
  c.Url = url
  c.Open = true
  c.Created = time.Now()
}

func (c *connection) Disconnect() {
  c.Open = false
}

func Connect(url string) *connection {
  conn := connection{}
  conn.Connect(url)
  return &conn
}
