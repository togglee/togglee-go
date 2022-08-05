package helpers

import (
	"encoding/json"
	. "github.com/togglee/togglee-go/pkg/models"
	"github.com/patrickmn/go-cache"
	"io/ioutil"
	"net/http"
	"time"
)

// ToggleCache -- cache remote updater
type ToggleCacheImp struct {
	URL     string
	cache   *cache.Cache
}

func refreshToggles(url string) *Toggles{
	response, err := http.Get(url)
	toggles := Toggles{}
	if err == nil {
		data, _ := ioutil.ReadAll(response.Body)
		_ = json.Unmarshal(data, &toggles)
	}
	return &toggles
}

func NewToggleCacheImp(url string, expiration time.Duration) ToggleCacheImp {
	tc := cache.New(expiration, 0)
	toggles := refreshToggles(url)
	tc.Set("toggles", toggles, cache.DefaultExpiration)
	tc.OnEvicted(func(k string, v interface{}) {
		toggles := refreshToggles(url)
		if toggles == nil {
			tcl, _ := tc.Get("toggles")
			toggles = tcl.(*Toggles)
		}
		tc.Set("toggles", toggles, cache.DefaultExpiration)
	})
	return ToggleCacheImp{ url, tc }
}

func (h *ToggleCacheImp) Toggles() *Toggles {
	toggles, _ := h.cache.Get("toggles")
	return toggles.(*Toggles)
}



