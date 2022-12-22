package cogen

import (
	"io/ioutil"

	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

func CoverGen(opts ...Option) {
	cfg := NewOptions(opts...)

	dc := gg.NewContext(cfg.Width, cfg.Height)
	dc.SetRGB255(cfg.CoverRGB[0], cfg.CoverRGB[1], cfg.CoverRGB[2])
	dc.Clear()

	fontFace, err := getOpenTypeFontFace(cfg.FontFilePath, cfg.FontSize, cfg.DPI)
	if err != nil {
		panic(err)
	}

	dc.SetFontFace(*fontFace)
	dc.SetRGB255(cfg.FontRGB[0], cfg.FontRGB[1], cfg.FontRGB[2])
	dc.DrawStringWrapped(cfg.Title, cfg.X, cfg.Y, cfg.AX, cfg.AY, cfg.MaxWordWidth, cfg.LineSpacing, gg.AlignCenter)

	// Start compressing images.
	src, err := imaging.Open(cfg.ImagePath)
	if err != nil {
		panic(err)
	}

	src = imaging.Resize(src, cfg.ResizeWidth, cfg.ResizeHeight, imaging.Lanczos)
	dc.DrawImageAnchored(src, cfg.AnchoredX, cfg.AnchoredY, cfg.AnchoredAX, cfg.AnchoredAY)
	dc.SavePNG(cfg.SavingFileName + ".png")
}

func getOpenTypeFontFace(fontFilePath string, fontSize, dpi float64) (*font.Face, error) {
	fontData, err := ioutil.ReadFile(fontFilePath)
	if err != nil {
		return nil, err
	}
	otfFont, err := opentype.Parse(fontData)
	if err != nil {
		return nil, err
	}
	otfFace, err := opentype.NewFace(otfFont, &opentype.FaceOptions{
		Size: fontSize,
		DPI:  dpi,
	})
	if err != nil {
		return nil, err
	}
	return &otfFace, nil
}
