package dashboards

// TileBounds the position and size of a tile
type TileBounds struct {
	Top    int32 `json:"top"`    // the vertical distance from the top left corner of the dashboard to the top left corner of the tile, in pixels
	Left   int32 `json:"left"`   // the horizontal distance from the top left corner of the dashboard to the top left corner of the tile, in pixels
	Width  int32 `json:"width"`  // the width of the tile, in pixels
	Height int32 `json:"height"` // the height of the tile, in pixels
}

// NewTileBounds creates a new instance of TileBounds and initializes the fields
func NewTileBounds(left, top, width, height int32) *TileBounds {
	return &TileBounds{Top: top, Left: left, Width: width, Height: height}
}
