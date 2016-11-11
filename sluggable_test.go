package sluggable

import "testing"

func TestSluggable(t *testing.T) {
	str := "hello-world"

	slug := Slugify("HELLO WORLD")

	if slug != str {
		t.Errorf("Expected %v got %v", str, slug)
	}

	slug2 := Slugify("This is a test ---")
	str2 := "this-is-a-test"

	if slug2 != str2 {
		t.Errorf("Expected %v got %v", str2, slug2)
	}

	slug3 := Slugify("This -- is a ## test ---")
	str3 := "this-is-a-test"

	if slug3 != str3 {
		t.Errorf("Expected %v got %v", str3, slug3)
	}
}
