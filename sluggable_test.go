package sluggable

import "testing"

func TestSluggable(t *testing.T) {
	str := "hello-world"

	slug := Slugify("Hello World")

	if slug != str {
		t.Errorf("Expected %v got %v", str, slug)
	}
}
