package internal

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"bytes"
)

func TestAPI(t *testing.T) {
	router := createRouter()
	server := httptest.NewServer(router)
	defer server.Close()

	resetearBD()

	matricula := "ABC123"
	url := server.URL + "/camion/" + matricula
	resp, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		t.Errorf("Esperado status 400, pero obtuve %d", resp.StatusCode)
	}

	matricula = "1234ABC"
	url = server.URL + "/camion/" + matricula
	resp, err = http.Get(url)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Esperado status 200, pero obtuve %d", resp.StatusCode)
	}

	var camion CamionMatricula
	err = json.NewDecoder(resp.Body).Decode(&camion)
	if err != nil {
		t.Fatal(err)
	}

	if camion.Matricula != matricula {
		t.Errorf("Esperado matricula %s, pero obtuve %s", matricula, camion.Matricula)
	}

	ID := "AAAAAA"
	url = server.URL + "/suministro/" + ID
	resp, err = http.Get(url)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		t.Errorf("Esperado status 400, pero obtuve %d", resp.StatusCode)
	}

	ID = "1"
	url = server.URL + "/suministro/" + ID
	resp, err = http.Get(url)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Esperado status 200, pero obtuve %d", resp.StatusCode)
	}

	var suministro SuministroID
	err = json.NewDecoder(resp.Body).Decode(&suministro)
	if err != nil {
		t.Fatal(err)
	}

	if suministro.ID != ID {
		t.Errorf("Esperado ID %s, pero obtuve %s", ID, suministro.ID)
	}

	ID = "AAAAAA"
	url = server.URL + "/suministro/" + ID + "/asignacion"
	resp, err = http.Get(url)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		t.Errorf("Esperado status 400, pero obtuve %d", resp.StatusCode)
	}

	ID = "3"
	url = server.URL + "/suministro/" + ID + "/asignacion"
	resp, err = http.Get(url)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Esperado status 200, pero obtuve %d", resp.StatusCode)
	}

	var asignacion Asignacion
	err = json.NewDecoder(resp.Body).Decode(&asignacion)
	if err != nil {
		t.Fatal(err)
	}

	if asignacion.IDSuministro != ID {
		t.Errorf("Esperado ID %s, pero obtuve %s", ID, asignacion.IDSuministro)
	}

	matricula = "AAAAAA"
	url = server.URL + "/camion/" + matricula
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		t.Fatal(err)
	}

	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		t.Errorf("Esperado status 400, pero obtuve %d", resp.StatusCode)
	}

	matricula = "1234ABE"
	url = server.URL + "/camion/" + matricula
	req, err = http.NewRequest("DELETE", url, nil)
	if err != nil {
		t.Fatal(err)
	}

	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Esperado status 200, pero obtuve %d", resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(&camion)
	if err != nil {
		t.Fatal(err)
	}

	if camion.Matricula != matricula {
		t.Errorf("Esperado matricula %s, pero obtuve %s", matricula, camion.Matricula)
	}

	ID = "AAAAAA"
	url = server.URL + "/suministro/" + ID
	req, err = http.NewRequest("DELETE", url, nil)
	if err != nil {
		t.Fatal(err)
	}
	
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		t.Errorf("Esperado status 400, pero obtuve %d", resp.StatusCode)
	}

	ID = "2"
	url = server.URL + "/suministro/" + ID
	req, err = http.NewRequest("DELETE", url, nil)
	if err != nil {
		t.Fatal(err)
	}

	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Esperado status 200, pero obtuve %d", resp.StatusCode)
	}


	ID = "AAAAAA"
	url = server.URL + "/suministro/" + ID + "/asignacion"
	req, err = http.NewRequest("POST", url, nil)
	if err != nil {
		t.Fatal(err)
	}

	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		t.Errorf("Esperado status 400, pero obtuve %d", resp.StatusCode)
	}

	ID = "2"
	url = server.URL + "/suministro/" + ID + "/asignacion"

	asignacionData := []string{"1234ABD"}

	jsonData, err := json.Marshal(asignacionData)
	if err != nil {
		t.Fatal(err)
	}

	req, err = http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Esperado status 200, pero obtuve %d", resp.StatusCode)
	}

	camionData := Camion{
		tipo:         	TipoSuministro(NORMAL),
		volumen_cm3:    1000,
		mma:            100,
	}
	matricula = "AAAAAA"
	url = server.URL + "/camion/" + matricula
	
	jsonData, err = json.Marshal(camionData)
	if err != nil {
		t.Fatal(err)
	}

	req, err = http.NewRequest("PUT", url, bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		t.Errorf("Esperado status 400, pero obtuve %d", resp.StatusCode)
	}

	matricula = "1234ABD"
	url = server.URL + "/camion/" + matricula

	jsonData, err = json.Marshal(camionData)
	if err != nil {
		t.Fatal(err)
	}

	req, err = http.NewRequest("PUT", url, bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Esperado status 200, pero obtuve %d", resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(&camion)
	if err != nil {
		t.Fatal(err)
	}

	if camion.Matricula != matricula {
		t.Errorf("Esperado matricula %s, pero obtuve %s", matricula, camion.Matricula)
	}

	ID = "AAAAAA"
	SuministroData := Suministro{
		direccion:      "Calle Falsa 123",
		tiempo:         10,
		volumen_cm3:    100,
		peso_kg:        1000,
		tipo:           TipoSuministro(NORMAL),
	}		
	url = server.URL + "/suministro/" + ID

	jsonData, err = json.Marshal(SuministroData)
	if err != nil {
		t.Fatal(err)
	}

	req, err = http.NewRequest("PUT", url, bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		t.Errorf("Esperado status 400, pero obtuve %d", resp.StatusCode)
	}

}