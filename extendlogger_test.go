package extendlogger

import (
	"bytes"
	"testing"
)

func TestInfoln(t *testing.T) {
	var buf bytes.Buffer

	l := NewExtendLogger(&buf)
	l.Infoln("test msg")
	got := buf.String()

	want := "[*] INFO test msg\n"

	if got != want {
		t.Errorf("got '%v', want '%v'", got, want)
	}

}

func TestWarnln(t *testing.T) {
	var buf bytes.Buffer

	l := NewExtendLogger(&buf)
	l.Warnln("test msg")
	got := buf.String()

	want := "[*] WARNING test msg\n"

	if got != want {
		t.Errorf("got '%v', want '%v'", got, want)
	}
}

func TestErrorln(t *testing.T) {
	var buf bytes.Buffer

	l := NewExtendLogger(&buf)
	l.Errorln("test msg")
	got := buf.String()

	want := "[*] ERROR test msg\n"

	if got != want {
		t.Errorf("got '%v', want '%v'", got, want)
	}

}
