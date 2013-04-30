package convert

import "math"

func HSVtoRGB(h, s, v int) (int, int, int) { //Convert HSV to RGB
	hf := float64(h)
	sf := float64(s) / 100.0
	vf := float64(v) / 100.0
	cf := vf * sf
	m := vf - cf
	xf := cf - cf*math.Abs(math.Mod(hf/60.0, 2)-1)
	r, g, b := 0.0, 0.0, 0.0
	if hf < 60 { //Piecewise function for converting HSV to RGB (source: Wikipedia)
		r, g, b = cf, xf, 0.0
	} else if hf < 120 {
		r, g, b = xf, cf, 0.0
	} else if hf < 180 {
		r, g, b = 0.0, cf, xf
	} else if hf < 240 {
		r, g, b = 0.0, xf, cf
	} else if hf < 300 {
		r, g, b = xf, 0.0, cf
	} else {
		r, g, b = cf, 0.0, xf
	}
	return int((r + m) * 255), int((g + m) * 255), int((b + m) * 255)
}

func HSLtoRGB(h, s, l int) (int, int, int) { //Convert HSL to RGB
	hf := float64(h)
	sf := float64(s) / 100.0
	lf := float64(l) / 100.0
	cf := (1 - math.Abs(2.*float64(l)-1.)) * sf
	mf := lf - 0.5*cf
	xf := cf - cf*math.Abs(math.Mod(hf/60.0, 2)-1)
	r, g, b := 0.0, 0.0, 0.0
	if hf < 60 { //Piecewise function for converting HSV to RGB (source: Wikipedia)
		r, g, b = cf, xf, 0.0
	} else if hf < 120 {
		r, g, b = xf, cf, 0.0
	} else if hf < 180 {
		r, g, b = 0.0, cf, xf
	} else if hf < 240 {
		r, g, b = 0.0, xf, cf
	} else if hf < 300 {
		r, g, b = xf, 0.0, cf
	} else {
		r, g, b = cf, 0.0, xf
	}
	return int((r + m) * 255), int((g + m) * 255), int((b + m) * 255)
}
