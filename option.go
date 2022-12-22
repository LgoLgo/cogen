package cogen

// Option is the only struct that can be used to set Options.
type Option struct {
	F func(o *Options)
}

type Options struct {
	// Title sets article title.
	//
	// Default: LgoLgo
	Title string

	// SavingFileName sets file name.
	//
	// Default: LgoLgo
	SavingFileName string

	// FontFilePath sets font file path.
	//
	// Default: font/syqh.ttf
	FontFilePath string

	// ImagePath sets image path
	//
	// Default: img/gopher.png
	ImagePath string

	// CoverRGB sets cover's RGB255.
	//
	// Default: []int{47, 54, 66}
	CoverRGB []int

	// FontRGB sets font's RGB255.
	//
	// Default: []int{238, 241, 247}
	FontRGB []int

	// FontSize sets font size.
	//
	//Default: 76
	FontSize float64

	// DPI sets font's DPI.
	//
	// Default: 72
	DPI float64

	// LineSpacing sets line spacing.
	//
	// Default: 1.1
	LineSpacing float64

	// Width sets cover width.
	//
	// Default: 1343
	Width int

	// ResizeWidth sets image width.
	//
	// Default: 1000
	ResizeWidth int

	// MaxWordWidth sets max word's width.
	//
	// Default: 660.0
	MaxWordWidth float64

	// Height sets height.
	//
	// Default: 734
	Height int

	// ResizeHeight sets image height.
	//
	// Default: 500
	ResizeHeight int

	// X sets font's X coordinate.
	//
	// Default: 800.0
	X float64

	// AX sets font's AX.
	//
	// Default: 0.5
	AX float64

	// AnchoredX sets image's X coordinate.
	//
	// Default: 200
	AnchoredX int

	// AnchoredAX set image's AX.
	//
	// Default: 0.5
	AnchoredAX float64

	// Y sets font's Y coordinate.
	//
	// Default: 300.0
	Y float64

	// AY sets font's AY.
	//
	// Default: 0.6
	AY float64

	// AnchoredY sets image's Y coordinate.
	//
	// Default: 220
	AnchoredY int

	// AnchoredAY set image's AY.
	//
	// Default: 0.5
	AnchoredAY float64
}

func (o *Options) Apply(opts []Option) {
	for _, op := range opts {
		op.F(o)
	}
}

// OptionsDefault is the default options.
var OptionsDefault = Options{
	Title:          "LgoLgo",
	SavingFileName: "LgoLgo",
	FontFilePath:   "font/syqh.ttf",
	ImagePath:      "img/gopher.png",
	CoverRGB:       []int{47, 54, 66},
	FontRGB:        []int{238, 241, 247},
	FontSize:       76,
	DPI:            72,
	LineSpacing:    1.1,
	Width:          1343,
	ResizeWidth:    1000,
	MaxWordWidth:   660.0,
	Height:         734,
	ResizeHeight:   500,
	X:              800.0,
	AX:             0.5,
	AnchoredX:      200,
	AnchoredAX:     0.5,
	Y:              300.0,
	AY:             0.6,
	AnchoredY:      220,
	AnchoredAY:     0.5,
}

func NewOptions(opts ...Option) *Options {
	options := &Options{
		Title:          OptionsDefault.Title,
		SavingFileName: OptionsDefault.SavingFileName,
		FontFilePath:   OptionsDefault.FontFilePath,
		ImagePath:      OptionsDefault.ImagePath,
		CoverRGB:       OptionsDefault.CoverRGB,
		FontRGB:        OptionsDefault.FontRGB,
		FontSize:       OptionsDefault.FontSize,
		DPI:            OptionsDefault.DPI,
		LineSpacing:    OptionsDefault.LineSpacing,
		Width:          OptionsDefault.Width,
		ResizeWidth:    OptionsDefault.ResizeWidth,
		MaxWordWidth:   OptionsDefault.MaxWordWidth,
		Height:         OptionsDefault.Height,
		ResizeHeight:   OptionsDefault.ResizeHeight,
		X:              OptionsDefault.X,
		AX:             OptionsDefault.AX,
		AnchoredX:      OptionsDefault.AnchoredX,
		AnchoredAX:     OptionsDefault.AnchoredAX,
		Y:              OptionsDefault.Y,
		AY:             OptionsDefault.AY,
		AnchoredY:      OptionsDefault.AnchoredY,
		AnchoredAY:     OptionsDefault.AnchoredAY,
	}
	options.Apply(opts)
	return options
}

// WithTitle sets title.
func WithTitle(title string) Option {
	return Option{
		F: func(o *Options) {
			o.Title = title
		},
	}
}

// WithSavingFileName sets title.
func WithSavingFileName(savingFileName string) Option {
	return Option{
		F: func(o *Options) {
			o.SavingFileName = savingFileName
		},
	}
}

// WithFontFilePath sets font file path.
func WithFontFilePath(path string) Option {
	return Option{
		F: func(o *Options) {
			o.FontFilePath = path
		},
	}
}

// WithImagePath sets image path.
func WithImagePath(path string) Option {
	return Option{
		F: func(o *Options) {
			o.ImagePath = path
		},
	}
}

// WithCoverRGB sets cover's RGB255.
func WithCoverRGB(rgb []int) Option {
	return Option{
		F: func(o *Options) {
			o.CoverRGB = rgb
		},
	}
}

// WithFontRGB sets font's RGB255.
func WithFontRGB(rgb []int) Option {
	return Option{
		F: func(o *Options) {
			o.FontRGB = rgb
		},
	}
}

// WithFontSize sets font's size.
func WithFontSize(size float64) Option {
	return Option{
		F: func(o *Options) {
			o.FontSize = size
		},
	}
}

// WithDPI sets font's DPI.
func WithDPI(dpi float64) Option {
	return Option{
		F: func(o *Options) {
			o.DPI = dpi
		},
	}
}

// WithLineSpacing sets line spacing.
func WithLineSpacing(lineSpacing float64) Option {
	return Option{
		F: func(o *Options) {
			o.LineSpacing = lineSpacing
		},
	}
}

// WithWidth sets cover's width.
func WithWidth(width int) Option {
	return Option{
		F: func(o *Options) {
			o.Width = width
		},
	}
}

// WithResizeWidth sets image's width.
func WithResizeWidth(width int) Option {
	return Option{
		F: func(o *Options) {
			o.ResizeWidth = width
		},
	}
}

// WithMaxWordWidth sets max word width.
func WithMaxWordWidth(width float64) Option {
	return Option{
		F: func(o *Options) {
			o.MaxWordWidth = width
		},
	}
}

// WithHeight sets cover's height.
func WithHeight(height int) Option {
	return Option{
		F: func(o *Options) {
			o.Height = height
		},
	}
}

// WithResizeHeight sets image's height.
func WithResizeHeight(height int) Option {
	return Option{
		F: func(o *Options) {
			o.ResizeHeight = height
		},
	}
}

// WithX sets font's X coordinate.
func WithX(x float64) Option {
	return Option{
		F: func(o *Options) {
			o.X = x
		},
	}
}

// WithAX sets font's AX.
func WithAX(ax float64) Option {
	return Option{
		F: func(o *Options) {
			o.AX = ax
		},
	}
}

// WithAnchoredX sets image's X coordinate.
func WithAnchoredX(x int) Option {
	return Option{
		F: func(o *Options) {
			o.AnchoredX = x
		},
	}
}

// WithAnchoredAX sets image's AX.
func WithAnchoredAX(ax float64) Option {
	return Option{
		F: func(o *Options) {
			o.AnchoredAX = ax
		},
	}
}

// WithY sets font's Y coordinate.
func WithY(y float64) Option {
	return Option{
		F: func(o *Options) {
			o.Y = y
		},
	}
}

// WithAY sets font's AY.
func WithAY(ay float64) Option {
	return Option{
		F: func(o *Options) {
			o.AY = ay
		},
	}
}

// WithAnchoredY sets image's Y coordinate.
func WithAnchoredY(y int) Option {
	return Option{
		F: func(o *Options) {
			o.AnchoredY = y
		},
	}
}

// WithAnchoredAY sets image's AY.
func WithAnchoredAY(ay float64) Option {
	return Option{
		F: func(o *Options) {
			o.AnchoredAY = ay
		},
	}
}
