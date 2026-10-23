# go-symbolic — brand assets

Official logos for the **go-symbolic** organization.

<img src="avatar/go-symbolic.png" width="160" alt="go-symbolic">

The mark is the integral sign, the operation a computer algebra system is judged by: one stroke, a stem with a hook at each end. It is written by `mark/main.go`, which computes the layout from a
few numbers; nothing is placed by hand. Its gradient (`#D98CE0` to `#9810A4`)
is a violet (hue 295) that sets go-symbolic apart from the indigo of go-numerics (267), the magenta of go-bignum and go-dataframe (321), the rose of go-dsp (335), the emerald of go-ndarray and go-fft, the blue of go-plot, the coral of go-sparse and the orange of go-approx; its dark end reads
on white at 7.13:1 and its light end on a dark page at 8.07:1.

## Formats

- **SVG** — `svg/color`, `svg/white` (white, transparent), `svg/black` (black, transparent)
- **PNG** 16→1024 px, 3 variants — `png/<variant>/<size>/`
- **JPG** (color, white background) — `jpg/<size>/`
- **ICO** (Windows) — `ico/`
- **ICNS** (macOS) — `icns/`
- **GitHub avatar** 512 px — `avatar/`
- **Social preview** 1280×640 (repo banner) — `social/`

`check/kit.py` counts the whole matrix on every pull request, and the `mark`
job runs `mark/main.go` again and fails if `svg/color/go-symbolic.svg` differs.

---
*Generated from `svg/color/go-symbolic.svg` by the fleet's Go brand pipeline (brandgen) — rounded square badge, line glyph, color / white / black variants.*
