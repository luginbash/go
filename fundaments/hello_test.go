package main

import "testing"

func TestHello(t *testing.T) {
  got := Hello("Name")
  want := "Hello Name"

  if got != want {
    t.Errorf("got %q want %q", got, want)
  }
}
