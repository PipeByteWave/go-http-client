package httptestwrap

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func NewRequestTestGo(t *testing.T, r *gin.Engine, method, url string, bod any, response any) (*httptest.ResponseRecorder, error) {
	assert.NotNil(t, response, "response no puede ser nil")
	assert.IsType(t, response, response, "response debe ser un puntero a una estructura o slice")
	var jsonData []byte
	var err error
	if bod != nil {
		jsonData, err = json.Marshal(bod)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, url ,bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	if len(jsonData) > 0 {
		req.Header.Set("Content-Type", "application/json")
	}

	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)

	err = json.Unmarshal(rr.Body.Bytes(), response)
	assert.NoError(t, err, "Error al leer el cuerpo de la respuesta")
	return rr, nil
}