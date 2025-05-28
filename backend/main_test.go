package main

import (
	"encoding/json"
	"math/big"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestFibonacci(t *testing.T) {
	// テーブル駆動テスト
	tests := []struct {
		name     string
		input    int
		expected *big.Int
		wantErr  error
	}{
		{
			name:     "n = 0",
			input:    0,
			expected: big.NewInt(0),
			wantErr:  nil,
		},
		{
			name:     "n = 1",
			input:    1,
			expected: big.NewInt(1),
			wantErr:  nil,
		},
		{
			name:     "n = 2",
			input:    2,
			expected: big.NewInt(1),
			wantErr:  nil,
		},
		{
			name:     "n = 10",
			input:    10,
			expected: big.NewInt(55),
			wantErr:  nil,
		},
		{
			name:     "n = -1",
			input:    -1,
			expected: nil,
			wantErr:  ErrInvalidInput,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fibonacci(tt.input)
			if tt.wantErr != nil {
				assert.ErrorIs(t, err, tt.wantErr)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestHandleFibonacci(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/fib", handleFibonacci)

	tests := []struct {
		name         string
		query        string
		expectedCode int
		expectedBody map[string]interface{}
	}{
		{
			name:         "正常系: n = 5",
			query:        "n=5",
			expectedCode: http.StatusOK,
			expectedBody: map[string]interface{}{
				"result": big.NewInt(5),
			},
		},
		{
			name:         "エラー: 無効な入力",
			query:        "n=invalid",
			expectedCode: http.StatusBadRequest,
			expectedBody: map[string]interface{}{
				"error": ErrBadRequest.Error(),
			},
		},
		{
			name:         "エラー: 負の数",
			query:        "n=-1",
			expectedCode: http.StatusBadRequest,
			expectedBody: map[string]interface{}{
				"error": ErrInvalidInput.Error(),
			},
		},
		{
			name:         "エラー: パラメータなし",
			query:        "",
			expectedCode: http.StatusBadRequest,
			expectedBody: map[string]interface{}{
				"error": ErrBadRequest.Error(),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/fib?"+tt.query, nil)
			router.ServeHTTP(recorder, req)

			assert.Equal(t, tt.expectedCode, recorder.Code)

			var response map[string]interface{}
			err := json.Unmarshal(recorder.Body.Bytes(), &response)
			assert.NoError(t, err)

			if tt.expectedCode == http.StatusOK {
				result := response["result"].(float64)
				expected := tt.expectedBody["result"].(*big.Int).Int64()
				assert.Equal(t, float64(expected), result)
			} else {
				assert.Equal(t, tt.expectedBody["error"], response["error"])
			}
		})
	}
}
