package colorize

import (
	"bytes"
	"image/color"
	"testing"

	"golang.org/x/image/colornames"
)

func TestColorize(t *testing.T) {
	testCases := []struct {
		fgColor   color.Color
		bgColor   color.Color
		input     []byte
		wantBytes []byte
	}{
		{
			input:     []byte("no colors"),
			wantBytes: []byte("no colors"),
		},
		{
			fgColor:   colornames.Darkblue,
			input:     []byte("no colors"),
			wantBytes: []byte("\033[38;2;0;0;139mno colors\033[0m"),
		},
		{
			bgColor:   colornames.Darkblue,
			input:     []byte("no colors"),
			wantBytes: []byte("\033[48;2;0;0;139mno colors\033[0m"),
		},
		{
			fgColor:   colornames.Darkblue,
			bgColor:   colornames.Darkblue,
			input:     []byte("no colors"),
			wantBytes: []byte("\033[38;2;0;0;139m\033[48;2;0;0;139mno colors\033[0m"),
		},
	}

	for _, tc := range testCases {
		colorizer := New(tc.fgColor, tc.bgColor)
		gotBytes := colorizer.Bytes(tc.input)
		if diff := bytes.Compare(gotBytes, tc.wantBytes); diff != 0 {
			t.Fatalf("got: %s, want: %s", string(gotBytes), string(tc.wantBytes))
		}
	}
}

func TestColorizeString(t *testing.T) {
	testCases := []struct {
		fgColor    color.Color
		bgColor    color.Color
		input      string
		wantString string
	}{
		{
			input:      "no colors",
			wantString: "no colors",
		},
		{
			fgColor:    colornames.Darkblue,
			input:      "no colors",
			wantString: "\033[38;2;0;0;139mno colors\033[0m",
		},
		{
			bgColor:    colornames.Darkblue,
			input:      "no colors",
			wantString: "\033[48;2;0;0;139mno colors\033[0m",
		},
		{
			fgColor:    colornames.Darkblue,
			bgColor:    colornames.Darkblue,
			input:      "no colors",
			wantString: "\033[38;2;0;0;139m\033[48;2;0;0;139mno colors\033[0m",
		},
	}

	for _, tc := range testCases {
		colorizer := New(tc.fgColor, tc.bgColor)
		if gotString := colorizer.colorizeString(tc.input); gotString != tc.wantString {
			t.Fatalf("got: %s, want: %s", gotString, tc.wantString)
		}
	}
}
