package icon

import (
	"testing"
)

func TestResource(t *testing.T) {
	res := Resource()

	if res == nil {
		t.Fatal("Resource() returned nil")
	}

	if res.Name() != "logo.svg" {
		t.Errorf("Expected resource name 'logo.svg', got '%s'", res.Name())
	}

	if len(res.Content()) == 0 {
		t.Error("Resource content is empty")
	}
}

func TestLogoDataEmbedded(t *testing.T) {
	if len(logoData) == 0 {
		t.Error("logoData is empty - SVG file not embedded")
	}

	// Check it starts with SVG content
	content := string(logoData)
	if len(content) < 5 || content[:5] != "<svg " && content[:5] != "<?xml" {
		t.Error("logoData does not appear to be valid SVG content")
	}
}
