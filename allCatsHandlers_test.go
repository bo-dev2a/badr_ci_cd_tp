package main

import (
	"testing"
)

// Cas simple — map avec un seul élément

func TestListMapKeys_Single(t *testing.T) {
	db := map[string]Cat{
		"id1": {Name: "Toto"},
	}
	keys := listMapKeys(db)
	if len(keys) != 1 {
		t.Errorf("Expected 1 key, got %d", len(keys))
	}
}

// Cas avancé — vérification de la clé retournée

func TestListMapKeys_CorrectKey(t *testing.T) {
	db := map[string]Cat{
		"abc123": {Name: "Felix"},
	}
	keys := listMapKeys(db)
	if keys[0] != "abc123" {
		t.Errorf("Expected key 'abc123', got '%s'", keys[0])
	}
}

// Cas limite — map vide

func TestListMapKeys_Empty(t *testing.T) {
	db := map[string]Cat{}
	keys := listMapKeys(db)
	if len(keys) != 0 {
		t.Errorf("Expected 0 keys, got %d", len(keys))
	}
}

