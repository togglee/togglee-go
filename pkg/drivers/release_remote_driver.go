package drivers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Toggles struct {
	Releases []ReleaseToggle
}

type ReleaseToggle struct {
	Name string
	Active bool
}

// ReleaseDriver -- release toggles driver resolver
type ReleaseRemoteDriver struct {
	Project      string
	URL          string
	Toggles		 Toggles
}

// IsActive -- release toggles driver resolver function
func (h *ReleaseRemoteDriver) IsActive(name string) *bool {
	response, err := http.Get(h.URL + "/" + h.Project + "/toggles")
	if err != nil {
		for _, value := range h.Toggles.Releases {
			if value.Name == name {
				return &value.Active
			}
		}
		return nil
	}
	data, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(data, &h.Toggles)
	for _, value := range h.Toggles.Releases {
		if value.Name == name {
			return &value.Active
		}
	}
	return nil
}
