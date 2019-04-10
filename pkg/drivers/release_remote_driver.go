package drivers

import (
	"encoding/json"
	. "github.com/feaggle/feaggle/pkg/models"
	"io/ioutil"
	"net/http"
)

// ReleaseDriver -- release toggles driver resolver
type ReleaseRemoteDriver struct {
	Project string
	URL     string
	cache   Toggles
}

func (h *ReleaseRemoteDriver) getToggleValue(name string) *bool {
	for _, value := range h.cache.Releases {
		if value.Name == name {
			return &value.Active
		}
	}
	return nil

}

// IsActive -- release toggles driver resolver function
func (h *ReleaseRemoteDriver) IsActive(name string) *bool {
	response, err := http.Get(h.URL + "/" + h.Project + "/toggles")
	if err != nil {
		return h.getToggleValue(name)
	}
	data, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(data, &h.cache)
	return h.getToggleValue(name)
}
