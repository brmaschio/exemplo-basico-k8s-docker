package main

import (
	"net/http"
    "net/http/httptest"
	"testing"
)

func TestHttp(t *testing.T){

    req, err := http.NewRequest("GET", "/", nil)
    if err != nil {
        t.Fatal(err)
	}
	
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(index)
    handler.ServeHTTP(rr, req)

	// TESTE DE STATUS
    if rr.Code != http.StatusOK {
		t.Errorf("Status Retornado Foi %v , Mas Era Esperado %v", rr.Code, http.StatusOK)
    }

    // VERIFICA BODY
    expected := `<b>Code.education Rocks!!!</b>`
    if rr.Body.String() != expected {
        t.Errorf("Retorno Do Body Inesteparo %v , Estava Aguardando %v", rr.Body.String(), expected)
	}
  
}
