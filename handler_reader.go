package main

import (
	"net/http"	
)

func handlerReader(w http.ResponseWriter,r *http.Request){
	respondWithJSON(w,200,struct{}{})
}