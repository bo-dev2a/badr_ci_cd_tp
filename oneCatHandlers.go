package main

import "net/http"

func getCat(req *http.Request) (int, any) {
	catId := req.PathValue("catId")

	cat, ok := catsDatabase[catId]
	if !ok {
		Logger.Info("Cat not found: ", catId)
		return http.StatusNotFound, "Cat not found"
	}

	Logger.Info("Cat found: ", catId)
	return http.StatusOK, cat
}

func deleteCat(req *http.Request) (int, any) {
	catId := req.PathValue("catId")

	_, ok := catsDatabase[catId]
	if !ok {
		Logger.Info("Cat not found: ", catId)
		return http.StatusNotFound, "Cat not found"
	}

	delete(catsDatabase, catId)

	Logger.Info("Cat deleted: ", catId)
	return http.StatusNoContent, nil
}
