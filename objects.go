package objects

import "fmt"

type Object struct {
	ID   string       `json:"id,omitmepty"`
	Type string       `json:"type"`
	Name string       `json:"name,omitmepty"`
	Prob float64      `json:"prob,omitmepty"`
	BBox [4]float64   `json:"bbox,omitmepty"`
	Area [][2]float64 `json:"area,omitmepty"`

	Data map[string]interface{} `json:"data,omitmepty"`
}

func (o Object) String() string {
	if o.Name != "" {
		return fmt.Sprintf("%s:%s", o.Type, o.Name)
	}

	return o.Type
}

func (o Object) IsValid() bool {
	return o.Type != ""
}
