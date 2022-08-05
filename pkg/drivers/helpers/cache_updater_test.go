package helpers_test

import (
	"fmt"
	. "github.com/togglee/togglee-go/pkg/drivers/helpers"
	. "github.com/togglee/togglee-go/pkg/models"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type fakeHandler struct {
	expectedResult bool
	failNextCalls bool
	callCount int
}

func (h *fakeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.callCount+=1
	if h.failNextCalls && h.callCount > 1 {
		panic("DOES NOT EXIST")
		return
	}
	result := fmt.Sprintf(`{
							"releases": [
								{
									"name": "TOGGLE_NAME",
									"active": %v
								}
							]
						}`, h.expectedResult)
	if r.URL.Path == "/myProject/toggles" {
		_, _ = fmt.Fprint(w, result)
		return
	} else if  r.URL.Path == "/wrongBody/toggles" {
		_, _ = fmt.Fprint(w, "pepe")
		return
	} else if  r.URL.Path == "/panic/toggles" {
		panic("DOES NOT EXIST")
		return
	}
	panic("WHY ARE YOU HERE")
}

func Test_Release_Remote_Driver(t *testing.T) {
	t.Run("takes true value from remote", func(t *testing.T) {
		handler := &fakeHandler{
			expectedResult: true,
			failNextCalls: false,
			callCount: 0,
		}
		server := httptest.NewServer(handler)
		defer server.Close()
		driver := NewToggleCacheImp("myProject", server.URL, 0)
		assert.Equal(t, driver.Toggles(), &Toggles{
			Releases:
				[]ReleaseToggle{
					{Name: "TOGGLE_NAME", Active: true},
				},
		})
	})

	t.Run("takes false value from remote", func(t *testing.T) {
		handler := &fakeHandler{
			expectedResult: false,
			failNextCalls: false,
			callCount: 0,
		}
		server := httptest.NewServer(handler)
		defer server.Close()
		driver := NewToggleCacheImp("myProject", server.URL , 0)
		assert.Equal(t, driver.Toggles(), &Toggles{
			Releases:
			[]ReleaseToggle{
				{Name: "TOGGLE_NAME", Active: false},
			},
		})
	})

	t.Run("if no remote return nil", func(t *testing.T) {
		handler := &fakeHandler{
			expectedResult: false,
			failNextCalls: false,
			callCount: 0,
		}
		server := httptest.NewServer(handler)
		defer server.Close()
		driver := NewToggleCacheImp("panic", server.URL, 0)
		assert.Equal(t, driver.Toggles(), &Toggles{ Releases: nil })
	})

	t.Run("if wrong body return nil", func(t *testing.T) {
		handler := &fakeHandler{
			expectedResult: false,
			failNextCalls: false,
			callCount: 0,
		}
		server := httptest.NewServer(handler)
		defer server.Close()
		driver := NewToggleCacheImp("wrongBody", server.URL, 0)
		assert.Equal(t, driver.Toggles(), &Toggles{ Releases: nil })
	})

	t.Run("if service is down return last know value", func(t *testing.T) {
		handler := &fakeHandler{
			expectedResult: true,
			failNextCalls: true,
			callCount: 0,
		}
		server := httptest.NewServer(handler)
		defer server.Close()
		driver := NewToggleCacheImp("myProject", server.URL, 0)
		assert.Equal(t, driver.Toggles(), &Toggles{
			Releases:
			[]ReleaseToggle{
				{Name: "TOGGLE_NAME", Active: true},
			},
		})
		assert.Equal(t, driver.Toggles(), &Toggles{
			Releases:
			[]ReleaseToggle{
				{Name: "TOGGLE_NAME", Active: true},
			},
		})
	})
}
