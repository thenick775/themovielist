package analytics

import "image/color"

type maskConf struct {
	File  string     `json:"file"`
	Color color.RGBA `json:"color"`
}

// Default Word Cloud Configuration
type confImg struct {
	FontMaxSize     int          `json:"font_max_size"`
	FontMinSize     int          `json:"font_min_size"`
	RandomPlacement bool         `json:"random_placement"`
	FontFile        string       `json:"font_file"`
	Colors          []color.RGBA `json:"colors"`
	BackgroundColor color.RGBA   `yaml:"background_color"`
	Width           int          `json:"width"`
	Height          int          `json:"height"`
	Mask            maskConf     `json:"mask"`
}

type listsSummary struct {
	Name                  string
	TotalContentCount     int
	ContentCountPerRating map[int]int
}
