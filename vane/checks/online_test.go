package checks

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/bearded-web/vane/vane/site"
)

func TestOnline500x(t *testing.T) {
	tsOk := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}))
	defer tsOk.Close()

	tsError := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		code := http.StatusInternalServerError
		http.Error(w, http.StatusText(code), code)
	}))
	defer tsError.Close()

	s, err := site.NewSite(tsOk.URL)
	assert.NoError(t, err)
	assert.True(t, Online(s))

	s, err = site.NewSite(tsError.URL)
	assert.NoError(t, err)
	assert.False(t, Online(s))

	s, err = site.NewSite(fakeHTTPaddress)
	assert.NoError(t, err)
	assert.False(t, Online(s))
}
