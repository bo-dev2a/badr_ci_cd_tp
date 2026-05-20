package apitests

import (
	"fmt"
	"net/http"
	"testing"
)

var initCatId string

func init() {
	// Remise à zéro de la base avant les tests
	ids := []string{}
	call("GET", "/cats", nil, nil, &ids)

	for _, id := range ids {
		code := 0
		call("DELETE", "/cats/"+id, nil, &code, nil)
		fmt.Println("DELETE /cats ->", code)
	}

	// Création d'un chat de référence
	call("POST", "/cats", &CatModel{Name: "Toto"}, nil, &initCatId)
}

// Simple — lister les chats, vérifier le code 200 et la présence du chat créé
func TestGetCats(t *testing.T) {
	code := 0
	result := []string{}
	err := call("GET", "/cats", nil, &code, &result)
	if err != nil {
		t.Error("Request error", err)
	}

	fmt.Println("GET /cats ->", code, result)

	if code != http.StatusOK {
		t.Errorf("Expected 200, got %d", code)
	}

	if len(result) != 1 {
		t.Errorf("Expected 1 item, got %d", len(result))
		return
	}

	if result[0] != initCatId {
		t.Errorf("Expected ID %s, got %s", initCatId, result[0])
	}
}

// Avancé — créer un chat et vérifier qu'il est bien récupérable
func TestCreateAndGetCat(t *testing.T) {
	// Création
	newId := ""
	code := 0
	err := call("POST", "/cats", &CatModel{Name: "Felix", Color: "Black"}, &code, &newId)
	if err != nil {
		t.Error("Request error", err)
	}

	if code != http.StatusCreated {
		t.Errorf("Expected 201, got %d", code)
	}

	// Récupération
	cat := CatModel{}
	err = call("GET", "/cats/"+newId, nil, &code, &cat)
	if err != nil {
		t.Error("Request error", err)
	}

	if code != http.StatusOK {
		t.Errorf("Expected 200, got %d", code)
	}

	if cat.Name != "Felix" {
		t.Errorf("Expected name Felix, got %s", cat.Name)
	}
}

// Complexe — supprimer un chat et vérifier qu'il n'est plus accessible
func TestDeleteCat(t *testing.T) {
	// Création d'un chat à supprimer
	newId := ""
	call("POST", "/cats", &CatModel{Name: "Temp"}, nil, &newId)

	// Suppression
	code := 0
	err := call("DELETE", "/cats/"+newId, nil, &code, nil)
	if err != nil {
		t.Error("Request error", err)
	}

	if code != http.StatusNoContent {
		t.Errorf("Expected 204, got %d", code)
	}

	// Vérification — le chat ne doit plus exister
	err = call("GET", "/cats/"+newId, nil, &code, nil)
	if err != nil {
		t.Error("Request error", err)
	}

	if code != http.StatusNotFound {
		t.Errorf("Expected 404, got %d", code)
	}
}
