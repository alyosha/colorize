package colorize

import (
	"bytes"
	"fmt"
	"image/color"
	"strings"
)

type Colorizer struct {
	foregroundColor, backgroundColor color.Color
}

func New(foregroundColor, backgroundColor color.Color) *Colorizer {
	return &Colorizer{
		foregroundColor: foregroundColor,
		backgroundColor: backgroundColor,
	}
}

func (c *Colorizer) UpdateForeground(foregroundColor color.Color) {
	c.foregroundColor = foregroundColor
}

func (c *Colorizer) UpdateBackground(backgroundColor color.Color) {
	c.backgroundColor = backgroundColor
}

func (c *Colorizer) Println(input interface{}) {
	fmt.Println(c.colorizeString(fmt.Sprintf("%s", input)))
}

func (c *Colorizer) Print(input interface{}) {
	input, trimmed := trimTrailingNewLine(input.(string))
	if trimmed {
		defer fmt.Print("\n")
	}
	fmt.Print(c.colorizeString(fmt.Sprintf("%s", input)))
}

func (c *Colorizer) Printf(format string, args ...interface{}) {
	format, trimmed := trimTrailingNewLine(format)
	if trimmed {
		defer fmt.Print("\n")
	}
	fmt.Print(c.colorizeString(fmt.Sprintf(format, args...)))
}

func (c *Colorizer) Sprintf(format string, args ...interface{}) string {
	return c.colorizeString(fmt.Sprintf(format, args...))
}

func (c *Colorizer) Bytes(input []byte) []byte {
	return c.colorize(input)
}

func (c *Colorizer) colorize(input []byte) []byte {
	var (
		buf            bytes.Buffer
		colorSpecified bool
	)

	if c.foregroundColor != nil {
		r, g, b := rgb(c.foregroundColor)
		buf.WriteString(fmt.Sprintf("\x1b[38;2;%d;%d;%dm", r, g, b))
		colorSpecified = true
	}

	if c.backgroundColor != nil {
		r, g, b := rgb(c.backgroundColor)
		buf.WriteString(fmt.Sprintf("\x1b[48;2;%d;%d;%dm", r, g, b))
		colorSpecified = true
	}

	buf.Write(input)

	if colorSpecified {
		buf.WriteString("\x1b[0m")
	}

	return buf.Bytes()
}

func (c *Colorizer) colorizeString(input string) string {
	return string(c.colorize([]byte(input)))
}

func rgb(c color.Color) (uint8, uint8, uint8) {
	rgba := color.RGBAModel.Convert(c).(color.RGBA)
	return rgba.R, rgba.G, rgba.B
}

func trimTrailingNewLine(input string) (string, bool) {
	var trimmed bool
	if strings.HasSuffix(input, "\n") {
		input = strings.TrimSuffix(input, "\n")
		trimmed = true
	}

	return input, trimmed
}
