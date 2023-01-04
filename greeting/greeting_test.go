package greeting

import "testing"

func TestSayHello(t *testing.T) {
	greeting := SayHello()
	if greeting != "hello" {
		t.Fatal("Expected greeting to be hello")
	}
}

func TestSayGoodbye(t *testing.T) {
	greeting := SayGoodBye()
	if greeting != "goodbye" {
		t.Fatal("Expected greeting to be goodbye")
	}
}
