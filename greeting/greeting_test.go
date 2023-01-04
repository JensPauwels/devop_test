package greeting

import "testing"

func TestSayHello(t *testing.T) {
	greeting := SayHello()
	// this should fail
	if greeting == "hello" {
		t.Fatal("Expected greeting to be hello")
	}
}
