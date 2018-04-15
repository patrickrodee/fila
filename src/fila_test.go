package fila

import (
  "testing"
)

func TestConnect(t *testing.T) {
  conn := Connect("https://foo/bar")
  if !conn.Open {
    t.Errorf("Connection should be open")
  }
}

func TestDisconnect(t *testing.T) {
  conn := Connect("https://foo/bar")
  conn.Disconnect()
  if conn.Open {
    t.Errorf("Connection should not be open")
  }
}