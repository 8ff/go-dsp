/*
 * Copyright (c) 2012 Matt Jibson <matt.jibson@gmail.com>
 *
 * Permission to use, copy, modify, and distribute this software for any
 * purpose with or without fee is hereby granted, provided that the above
 * copyright notice and this permission notice appear in all copies.
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

package spectral

import (
	"testing"

	"github.com/NectGmbH/go-dsp/dsputils"
)

type pwelchTest struct {
	fs          float64
	opts        *PwelchOptions
	x, p, freqs []float64
}

var pwelchTests = []pwelchTest{
	{ // zero-length input
		0,
		&PwelchOptions{},
		[]float64{},
		[]float64{},
		[]float64{},
	},
	{
		2,
		&PwelchOptions{},
		[]float64{
			0,
			1,
			2,
			3,
			4,
			5,
			6,
			7,
			8,
			9,
			10,
			11,
			12,
			13,
			14,
			15,
			16,
			17,
			18,
			19,
			20,
			21,
			22,
			23,
			24,
			25,
			26,
			27,
			28,
			29,
			30,
			31,
			32,
			33,
			34,
			35,
			36,
			37,
			38,
			39,
			40,
			41,
			42,
			43,
			44,
			45,
			46,
			47,
			48,
			49,
			50,
			51,
			52,
			53,
			54,
			55,
			56,
			57,
			58,
			59,
			60,
			61,
			62,
			63,
			64,
			65,
			66,
			67,
			68,
			69,
			70,
			71,
			72,
			73,
			74,
			75,
			76,
			77,
			78,
			79,
			80,
			81,
			82,
			83,
			84,
			85,
			86,
			87,
			88,
			89,
			90,
			91,
			92,
			93,
			94,
			95,
			96,
			97,
			98,
			99,
		},
		[]float64{
			3.66817103e+04,
			6.16097526e+04,
			3.70964854e+04,
			1.76858083e+04,
			8.82747121e+03,
			5.58636625e+03,
			3.86686565e+03,
			2.79695091e+03,
			2.14687978e+03,
			1.68918004e+03,
			1.36571705e+03,
			1.13024093e+03,
			9.48033939e+02,
			8.08850444e+02,
			6.97809757e+02,
			6.08092372e+02,
			5.35404251e+02,
			4.74620274e+02,
			4.24037212e+02,
			3.81226909e+02,
			3.44548926e+02,
			3.13192558e+02,
			2.85886182e+02,
			2.62122493e+02,
			2.41303266e+02,
			2.22870690e+02,
			2.06594463e+02,
			1.92060902e+02,
			1.79062190e+02,
			1.67411631e+02,
			1.56878696e+02,
			1.47375046e+02,
			1.38742768e+02,
			1.30879468e+02,
			1.23716560e+02,
			1.17146757e+02,
			1.11127186e+02,
			1.05591138e+02,
			1.00482309e+02,
			9.57717459e+01,
			9.14056404e+01,
			8.73592894e+01,
			8.36025117e+01,
			8.01022290e+01,
			7.68443525e+01,
			7.38005900e+01,
			7.09550385e+01,
			6.82933042e+01,
			6.57951735e+01,
			6.34526724e+01,
			6.12504908e+01,
			5.91777124e+01,
			5.72271084e+01,
			5.53860529e+01,
			5.36493451e+01,
			5.20085636e+01,
			5.04559628e+01,
			4.89876620e+01,
			4.75956980e+01,
			4.62762918e+01,
			4.50247688e+01,
			4.38355570e+01,
			4.27063668e+01,
			4.16321728e+01,
			4.06100428e+01,
			3.96373615e+01,
			3.87101146e+01,
			3.78267782e+01,
			3.69842029e+01,
			3.61800421e+01,
			3.54128094e+01,
			3.46796320e+01,
			3.39793658e+01,
			3.33100629e+01,
			3.26698301e+01,
			3.20577904e+01,
			3.14719152e+01,
			3.09112634e+01,
			3.03746526e+01,
			2.98605643e+01,
			2.93684407e+01,
			2.88968774e+01,
			2.84450603e+01,
			2.80122875e+01,
			2.75973586e+01,
			2.71998759e+01,
			2.68188936e+01,
			2.64536948e+01,
			2.61038720e+01,
			2.57684964e+01,
			2.54472465e+01,
			2.51395088e+01,
			2.48446551e+01,
			2.45624511e+01,
			2.42921985e+01,
			2.40336109e+01,
			2.37863119e+01,
			2.35497603e+01,
			2.33238184e+01,
			2.31079809e+01,
			2.29019795e+01,
			2.27056035e+01,
			2.25183990e+01,
			2.23402769e+01,
			2.21708920e+01,
			2.20099898e+01,
			2.18574728e+01,
			2.17129732e+01,
			2.15764231e+01,
			2.14476081e+01,
			2.13262901e+01,
			2.12124459e+01,
			2.11057929e+01,
			2.10062684e+01,
			2.09137648e+01,
			2.08280657e+01,
			2.07491945e+01,
			2.06769518e+01,
			2.06112729e+01,
			2.05521368e+01,
			2.04993557e+01,
			2.04529802e+01,
			2.04128917e+01,
			2.03790224e+01,
			2.03514209e+01,
			2.03299362e+01,
			2.03146325e+01,
			2.03054705e+01,
			1.01511907e+01,
		},
		[]float64{
			0,
			0.0078125,
			0.015625,
			0.0234375,
			0.03125,
			0.0390625,
			0.046875,
			0.0546875,
			0.0625,
			0.0703125,
			0.078125,
			0.0859375,
			0.09375,
			0.1015625,
			0.109375,
			0.1171875,
			0.125,
			0.1328125,
			0.140625,
			0.1484375,
			0.15625,
			0.1640625,
			0.171875,
			0.1796875,
			0.1875,
			0.1953125,
			0.203125,
			0.2109375,
			0.21875,
			0.2265625,
			0.234375,
			0.2421875,
			0.25,
			0.2578125,
			0.265625,
			0.2734375,
			0.28125,
			0.2890625,
			0.296875,
			0.3046875,
			0.3125,
			0.3203125,
			0.328125,
			0.3359375,
			0.34375,
			0.3515625,
			0.359375,
			0.3671875,
			0.375,
			0.3828125,
			0.390625,
			0.3984375,
			0.40625,
			0.4140625,
			0.421875,
			0.4296875,
			0.4375,
			0.4453125,
			0.453125,
			0.4609375,
			0.46875,
			0.4765625,
			0.484375,
			0.4921875,
			0.5,
			0.5078125,
			0.515625,
			0.5234375,
			0.53125,
			0.5390625,
			0.546875,
			0.5546875,
			0.5625,
			0.5703125,
			0.578125,
			0.5859375,
			0.59375,
			0.6015625,
			0.609375,
			0.6171875,
			0.625,
			0.6328125,
			0.640625,
			0.6484375,
			0.65625,
			0.6640625,
			0.671875,
			0.6796875,
			0.6875,
			0.6953125,
			0.703125,
			0.7109375,
			0.71875,
			0.7265625,
			0.734375,
			0.7421875,
			0.75,
			0.7578125,
			0.765625,
			0.7734375,
			0.78125,
			0.7890625,
			0.796875,
			0.8046875,
			0.8125,
			0.8203125,
			0.828125,
			0.8359375,
			0.84375,
			0.8515625,
			0.859375,
			0.8671875,
			0.875,
			0.8828125,
			0.890625,
			0.8984375,
			0.90625,
			0.9140625,
			0.921875,
			0.9296875,
			0.9375,
			0.9453125,
			0.953125,
			0.9609375,
			0.96875,
			0.9765625,
			0.984375,
			0.9921875,
			1,
		},
	},
}

func TestPwelch(t *testing.T) {
	for _, v := range pwelchTests {
		p, freqs := Pwelch(v.x, v.fs, v.opts)
		if !dsputils.PrettyClose(p, v.p) {
			t.Error("Pwelch Pxx error\n   input:", v.x, "\n  output:", p, "\nexpected:", v.p)
		}

		if !dsputils.PrettyClose(freqs, v.freqs) {
			t.Error(
				"Pwelch freqs error\n   input:",
				v.x,
				"\n  output:",
				freqs,
				"\nexpected:",
				v.freqs,
			)
		}
	}
}
