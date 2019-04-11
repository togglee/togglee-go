package helpers

import (
	"encoding/json"
	. "github.com/feaggle/feaggle/pkg/models"
	"github.com/patrickmn/go-cache"
	"io/ioutil"
	"net/http"
	"time"
)

// ToggleCache -- cache remote updater
type ToggleCacheImp struct {
	Project string
	URL     string
	cache   *cache.Cache
}

func refreshToggles(project string, url string) *Toggles{
	response, err := http.Get(url + "/" + project + "/toggles")
	toggles := Toggles{}
	if err == nil {
		data, _ := ioutil.ReadAll(response.Body)
		_ = json.Unmarshal(data, &toggles)
	}
	return &toggles
}

func NewToggleCacheImp(project string, url string, expiration time.Duration) ToggleCacheImp {
	tc := cache.New(expiration, 0)
	toggles := refreshToggles(project, url)
	tc.Set("toggles", toggles, cache.DefaultExpiration)
	tc.OnEvicted(func(k string, v interface{}) {
		toggles := refreshToggles(project, url)
		if toggles == nil {
			return
		}
		tc.Set("toggles", toggles, cache.DefaultExpiration)
	})
	return ToggleCacheImp{ project, url, tc }
}

func (h *ToggleCacheImp) Toggles() *Toggles {
	toggles, _ := h.cache.Get("toggles")
	return toggles.(*Toggles)
}



