package keyValueStore_test

import (
	"github.com/devashishRaj/goTools/keyValueStore"
	"os"
	"testing"
)

func TestGet_ReturnsNotOKIfKeyDoesNotExist(t *testing.T) {
	t.Parallel()
	s, err := keyValueStore.OpenStore("dummy path")
	if err != nil {
		t.Fatal(err)
	}
	_, ok := s.Get("Key")
	if ok {
		t.Fatal("Unexpected OK")
	}
}

func TestGet_ReturnsValueAndOKIfKeyDoesExist(t *testing.T) {
	t.Parallel()
	s, err := keyValueStore.OpenStore("dummy path")
	if err != nil {
		t.Fatal(err)
	}
	s.Set("key", "value")
	v, ok := s.Get("key")
	if !ok {
		t.Fatal("not ok")
	}
	if v != "value" {
		t.Errorf("want 'value', got %q", v)

	}
}

func TestSetUpdatesExistingKeyToNewValue(t *testing.T) {
	t.Parallel()
	s, err := keyValueStore.OpenStore("dummy path")
	if err != nil {
		t.Fatal(err)
	}
	s.Set("Key", "Value")
	s.Set("Key", "Updated")
	v, ok := s.Get("Key")
	if !ok {
		t.Fatalf("Not OK")
	}
	if v != "Updated" {
		t.Fatalf("Wanted 'Update', got %q", v)
	}
}

func TestSave_SavesDataPersistently(t *testing.T) {
	t.Parallel()
	path := t.TempDir() + "/keyValueStoretest.store"
	// We open a store 's' at path and save data there
	s, err := keyValueStore.OpenStore(path)
	if err != nil {
		t.Fatal(err)
	}
	s.Set("A", "1")
	s.Set("B", "2")
	s.Set("C", "3")
	err = s.Save()
	if err != nil {
		t.Fatal(err)
	}
	// open a store 's2' from the same path to check if we can load saved data
	s2, err := keyValueStore.OpenStore(path)
	if err != nil {
		t.Fatal(err)
	}
	if v, _ := s2.Get("A"); v != "1" {
		t.Fatalf("want A=1, got A=%s", v)
	}
	if v, _ := s2.Get("B"); v != "2" {
		t.Fatalf("want B=2, got B=%s", v)
	}
	if v, _ := s2.Get("C"); v != "3" {
		t.Fatalf("want C=3, got C=%s", v)
	}
}

func TestOpenStore_ErrorsWhenPathUnreadable(t *testing.T) {
	t.Parallel()
	path := t.TempDir() + "/unreadable.store"
	if _, err := os.Create(path); err != nil {
		t.Fatal(err)
	}
	if err := os.Chmod(path, 0o000); err != nil {
		t.Fatal(err)
	}
	_, err := keyValueStore.OpenStore(path)
	if err == nil {
		t.Fatal("no error")
	}
}

func TestOpenStore_ReturnsErrorOnInvalidData(t *testing.T) {
	t.Parallel()
	_, err := keyValueStore.OpenStore("testdata/invalid.store")
	if err == nil {
		t.Fatal("no error")
	}
}

func TestSaveErrorsWhenPathUnwritable(t *testing.T) {
	t.Parallel()
	s, err := keyValueStore.OpenStore("bogus/unwritable.store")
	if err != nil {
		t.Fatal(err)
	}
	err = s.Save()
	if err == nil {
		t.Fatal("no error")
	}
}
