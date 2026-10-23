// Command mark writes the go-symbolic mark, svg/color/go-symbolic.svg:
//
//	cd mark && go run . > ../svg/color/go-symbolic.svg
//
// go-symbolic is computer algebra, and the mark is its emblem: the integral
// sign, the operation a CAS is judged by. It is one stroke, a stem with a
// hook at each end.
//
// Nothing is placed by hand. The program lists the four points and two
// control points of the top hook, the stem runs from the top hook to the point
// that is its image under a half turn about the centre of the square, and the
// bottom hook is the top hook turned the same way and walked backwards. Change
// the numbers and the sign redraws itself, still symmetric.
//
// The rest is the family grammar every go-* mark shares: a 256-unit rounded
// square (rx 56) filled with a diagonal two-stop gradient, a white line glyph
// at stroke-width 13 with round caps and joins. logos/brandgen renders the
// format matrix from the file this writes; the kit workflow runs this again
// and fails if a byte of svg/color differs.
package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

const (
	light = "#D98CE0" // the gradient's light stop
	dark  = "#9810A4" // its dark stop

	stroke = 13.0 // the family's line width
	centre = 128.0
)

type point struct{ x, y float64 }

// turn is the half turn about the centre of the square.
func turn(p point) point { return point{2*centre - p.x, 2*centre - p.y} }

// The top hook: it starts at its free end a, leaves upward through c1 and c2,
// and comes down to the head of the stem at b.
var (
	a  = point{172, 90}
	c1 = point{172, 70}
	c2 = point{152, 62}
	b  = point{144, 82}
)

func pt(p point) string { return num(p.x) + " " + num(p.y) }

// num prints a coordinate with at most one decimal and no trailing zero.
func num(v float64) string {
	s := fmt.Sprintf("%.1f", math.Round(v*10)/10)
	s = strings.TrimSuffix(s, ".0")
	if s == "-0" {
		s = "0"
	}
	return s
}

func main() {
	d := fmt.Sprintf("M%sC%s %s %sL%sC%s %s %s",
		pt(a), pt(c1), pt(c2), pt(b), // the top hook
		pt(turn(b)),                           // the stem
		pt(turn(c2)), pt(turn(c1)), pt(turn(a))) // the bottom hook, walked backwards

	fmt.Fprintf(os.Stdout, `<svg viewBox="0 0 256 256" xmlns="http://www.w3.org/2000/svg" role="img" aria-label="go-symbolic">
  <defs><linearGradient id="sygrad" x1="0" y1="0" x2="1" y2="1"><stop offset="0" stop-color="%s"/><stop offset="1" stop-color="%s"/></linearGradient></defs>
  <rect width="256" height="256" rx="56" fill="url(#sygrad)"/>
  <g fill="none" stroke="#FFFFFF" stroke-width="%s" stroke-linecap="round" stroke-linejoin="round">
    <path d="%s"/>
  </g>
</svg>
`, light, dark, num(stroke), d)
}
