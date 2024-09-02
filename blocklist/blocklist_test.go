package blocklist

import "testing"

func Test_BlockList(t *testing.T) {
	t.Run("make blocklist", func(t *testing.T) {
		b := New()
		if b == nil {
			t.Error("Blocklist is nil")
		}
	})
	t.Run("add item", func(t *testing.T) {
		b := New()
		err := b.BlockItems.add("http://example.com")
		if err != nil {
			t.Error("Error adding item")
		}
	})
	t.Run("add duplicate item", func(t *testing.T) {
		b := New()
		err := b.BlockItems.add("http://example.com")
		if err != nil {
			t.Error("Error adding item")
		}
		err = b.BlockItems.add("http://example.com")
		if err == nil {
			t.Error("Error adding duplicate item")
		}
	})
	t.Run("check item", func(t *testing.T) {
		b := New()
		err := b.BlockItems.add("http://example.com")
		if err != nil {
			t.Error("Error adding item")
		}
		err = b.BlockItems.Check("http://example.com")
		if err != nil {
			t.Error("Error checking item")
		}
	})
	t.Run("check missing item", func(t *testing.T) {
		b := New()
		err := b.BlockItems.Check("http://example.com")
		if err == nil {
			t.Error("Error checking missing item")
		}
	})
	t.Run("batch add items", func(t *testing.T) {
		b := New()
		err := b.BlockItems.Batch(BatchBlock{urls: []string{"http://example.com", "http://example2.com"}})
		if err != nil {
			t.Error("Error adding items")
		}
	})

	t.Run("remove item", func(t *testing.T) {
		b := New()
		err := b.BlockItems.add("http://example.com")
		if err != nil {
			t.Error("Error adding item")
		}
		err = b.BlockItems.remove("http://example.com")
		if err != nil {
			t.Error("Error removing item")
		}
	})
}
