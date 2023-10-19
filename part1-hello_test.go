package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris","")
		want := "Hello Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'world'", func(t *testing.T) {
		got := Hello("","")
		want := "Hello world"
    assertCorrectMessage(t, got, want)
	})
  t.Run("can people hablo espanol?", func(t *testing.T){
    got := Hello("Ezikel", "Spanish")
    want := "Hola Ezikel"
    assertCorrectMessage(t, got, want)
  })
  t.Run("tu parles fracais?", func (t *testing.T){
    got := Hello("bonaparte", "French")
    want := "Bonjour bonaparte"
    assertCorrectMessage(t, got, want)
  })

}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
