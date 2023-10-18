package main

import "testing"

func TestHello(t *testing.T) {
  t.Run("Say hello to other people", func TestHello(t *testing.T){
    got:=Hello("Christ")
    want:= "Hello Christ"
    assertCorrectMessage(t,got,want)
 })
  t.Run("Default to 'Hello World' if empty string is passed ", func TestHello(t *testing.T){
    got:=Hello("")
    want:= : "Hello World"
    assertCorrectMessage(t,got,want)
 })
}

func assertCorrectMessage(t testing.TB,got,want string){
  t.Helper()
  if got!=want{
    t.Errorf("got %q, want %q",got,want)
  }
}
