package daemon

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type MockMetrics struct {
	mock.Mock
}

func (m *MockMetrics) Read(rdr io.Reader) error {
	args := m.Called(rdr)
	return args.Error(0)
}

func (m *MockMetrics) Handler() http.Handler {
	args := m.Called()
	return args.Get(0).(http.Handler)
}

func TestUpdateHandlerUnsupportedMediaType(t *testing.T) {
	t.Parallel()

	m := &MockMetrics{}
	d := New(m)

	req, err := http.NewRequest("POST", "/metrics", bytes.NewReader([]byte("{}")))
	require.NoError(t, err)
	rr := httptest.NewRecorder()
	d.router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusUnsupportedMediaType, rr.Code)
	m.AssertExpectations(t)
}

func TestUpdateHandlerInternalServerError(t *testing.T) {
	t.Parallel()

	m := &MockMetrics{}
	m.On("Read", mock.Anything).Return(errors.New("mock error"))
	d := New(m)

	req, err := http.NewRequest("POST", "/metrics", bytes.NewReader([]byte("{}")))
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	http.Handler(d.router).ServeHTTP(rr, req)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	assert.Equal(t, "mock error", rr.Body.String())
	m.AssertExpectations(t)
}

func TestUpdateHandlerStatusAccepted(t *testing.T) {
	t.Parallel()

	m := &MockMetrics{}
	m.On("Read", mock.Anything).Return(nil)
	d := New(m)

	req, err := http.NewRequest("POST", "/metrics", bytes.NewReader([]byte("{}")))
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	http.Handler(d.router).ServeHTTP(rr, req)
	assert.Equal(t, http.StatusNoContent, rr.Code)
	m.AssertExpectations(t)
}
