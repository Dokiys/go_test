package pkg

import (
	"testing"

	"github.com/gookit/color"
)

func TestColor(t *testing.T) {
	color.Redln("red")
	color.Yellowln("yellow")
	color.RGBFromString("170,187,204").Color().Println("rgb 170,187,204")
}
