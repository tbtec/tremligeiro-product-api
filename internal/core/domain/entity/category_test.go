package entity

import (
	"testing"
)

func TestCategoryFields(t *testing.T) {
	c := Category{
		ID:   1,
		Name: "Electronics",
	}

	if c.ID != 1 {
		t.Errorf("expected ID to be 1, got %d", c.ID)
	}
	if c.Name != "Electronics" {
		t.Errorf("expected Name to be 'Electronics', got '%s'", c.Name)
	}
}
