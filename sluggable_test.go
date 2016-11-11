package sluggable

import "testing"

func TestSluggable(t *testing.T) {
	str := "hello-world"

	slug := Slugify("HELLO WORLD")

	if slug != str {
		t.Errorf("Expected %v got %v", str, slug)
	}
}
