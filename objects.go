package objects

type Object struct {
	ID   string       `json:"id"`
	Type string       `json:"type"`
	Name string       `json:"name"`
	Prob float64      `json:"prob"`
	BBox [4]float64   `json:"bbox"`
	Area [][2]float64 `json:"area"`
}
