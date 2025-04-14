package integration

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"avito_intern/internal/handler"
	"avito_intern/internal/repository"
	"avito_intern/internal/service"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func setupTestServer(t *testing.T) (*sqlx.DB, http.Handler) {
	db, err := sqlx.Connect("postgres", "postgres://postgres:03795@localhost:5432/avito_intern?sslmode=disable")
	if err != nil {
		t.Fatalf("failed to connect to test db: %v", err)
	}
	repoComposite := repository.NewRepository(db)
	svc := service.NewService(repoComposite)
	h := handler.NewHandler(svc)
	router := h.InitRoutes()
	return db, router
}

func getDummyToken(t *testing.T, router http.Handler, role string) string {
	body, _ := json.Marshal(map[string]string{"role": role})
	req := httptest.NewRequest(http.MethodPost, "/dummyLogin", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	var resp map[string]string
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	if err != nil || resp["token"] == "" {
		t.Fatalf("failed to get dummy token for role %s", role)
	}
	return resp["token"]
}

func TestFullFlow(t *testing.T) {
	db, router := setupTestServer(t)
	defer db.Close()

	modToken := getDummyToken(t, router, "moderator")
	empToken := getDummyToken(t, router, "employee")

	pvzBody := map[string]string{"city": "Москва"}
	bodyBytes, _ := json.Marshal(pvzBody)
	req := httptest.NewRequest(http.MethodPost, "/pvz/", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+modToken)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusCreated, rr.Code)

	var pvzResp map[string]string
	err := json.Unmarshal(rr.Body.Bytes(), &pvzResp)
	assert.NoError(t, err)
	pvzID := pvzResp["id"]
	assert.NotEmpty(t, pvzID)

	recBody := map[string]string{"pvzId": pvzID}
	bodyBytes, _ = json.Marshal(recBody)
	req = httptest.NewRequest(http.MethodPost, "/receptions", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+empToken)
	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusCreated, rr.Code)

	for i := 0; i < 50; i++ {
		prodBody := map[string]interface{}{"pvzId": pvzID, "type": "обувь"}
		bodyBytes, _ = json.Marshal(prodBody)
		req = httptest.NewRequest(http.MethodPost, "/products", bytes.NewReader(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+empToken)
		rr = httptest.NewRecorder()
		router.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusCreated, rr.Code, fmt.Sprintf("product addition %d", i+1))
	}

	closeURL := fmt.Sprintf("/pvz/%s/close_last_reception", pvzID)
	req = httptest.NewRequest(http.MethodPost, closeURL, nil)
	req.Header.Set("Authorization", "Bearer "+empToken)
	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)

	var closedRec map[string]interface{}
	err = json.Unmarshal(rr.Body.Bytes(), &closedRec)
	assert.NoError(t, err)
	assert.Equal(t, "close", closedRec["status"])
}
