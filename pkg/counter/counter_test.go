package counter_test

import (
	"encoding/json"
	"testing"

	"github.com/malisetti/wcjson/pkg/counter"
)

func mustJSON(t *testing.T, s string) any {
	t.Helper()
	var v any
	if err := json.Unmarshal([]byte(s), &v); err != nil {
		t.Fatalf("json.Unmarshal: %v", err)
	}
	return v
}

func TestKeyCount(t *testing.T) {
	tests := []struct {
		name string
		json string
		want int
	}{
		{"null", `null`, 0},
		{"scalar", `42`, 0},
		{"empty object", `{}`, 0},
		{"flat object", `{"a":1,"b":2}`, 2},
		{"nested object", `{"a":{"b":1,"c":2},"d":3}`, 4},
		{"array of scalars", `[1,2,3]`, 0},
		{"array of objects", `[{"x":1},{"y":2,"z":3}]`, 3},
		{"mixed", `{"items":[{"k":1},{"k":2}],"n":0}`, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := mustJSON(t, tt.json)
			if got := counter.KeyCount(v); got != tt.want {
				t.Fatalf("KeyCount() = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestLeafCount(t *testing.T) {
	tests := []struct {
		name string
		json string
		want int
	}{
		{"null", `null`, 1},
		{"scalar", `"x"`, 1},
		{"empty object", `{}`, 0},
		{"empty array", `[]`, 0},
		{"flat object", `{"a":1,"b":2}`, 2},
		{"nested object", `{"a":{"b":1,"c":2},"d":3}`, 3},
		{"array of scalars", `[1,2,3]`, 3},
		{"array of objects", `[{"x":1},{"y":2,"z":3}]`, 3},
		{"mixed", `{"items":[{"k":1},{"k":2}],"n":0}`, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := mustJSON(t, tt.json)
			if got := counter.LeafCount(v); got != tt.want {
				t.Fatalf("LeafCount() = %d, want %d", got, tt.want)
			}
		})
	}
}
