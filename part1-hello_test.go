package main

import "testing"

func TestHello(t *testing.T){
  t.Run("Say Hello to other people", func(t *testing.T){ got:= Hello("Jezus")
    want:= "Hello Jezus"

    if got!=want{
      t.Errorf("got %q, want %q",got,want)
    }
  })
  t.Run("Default back to Hello world in case of empty string", func(t *testing.T){
    got:= Hello("")
    want:= "Hello world"

    if got!=want{
      t.Errorf("got %q, want %q",got,want)
    }
  })
}
