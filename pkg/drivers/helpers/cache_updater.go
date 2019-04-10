package helpers

import (
	"encoding/json"
	. "github.com/feaggle/feaggle/pkg/models"
	"io/ioutil"
	"net/http"
)

// ToggleCache -- cache remote updater
type ToggleCacheImp struct {
	Project string
	URL     string
	cache   Toggles
}

// Validate -- validates the cache if expired
func (h *ToggleCacheImp) Validate() {
	response, err := http.Get(h.URL + "/" + h.Project + "/toggles")
	if err != nil {
		return
	}
	data, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(data, &h.cache)
}

func (h *ToggleCacheImp) Cache() *Toggles {
	return &h.cache
}



