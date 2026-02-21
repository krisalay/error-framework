package core

import "testing"

func TestBuilder_Build(t *testing.T) {

	err := New().
		WithMessage("msg").
		WithCode("CODE").
		WithStatus(400).
		WithSensitive(false).
		Build()

	if err.Message != "msg" {
		t.Fatal("Message incorrect")
	}

	if err.Code != "CODE" {
		t.Fatal("Code incorrect")
	}

	if err.Status != 400 {
		t.Fatal("Status incorrect")
	}
}
