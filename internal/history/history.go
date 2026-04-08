package history

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sort"
	"time"
)

type Result struct {
	WPM      float64   `json:"wpm"`
	Accuracy float64   `json:"accuracy"`
	Duration float64   `json:"duration_secs"`
	Mode     string    `json:"mode"`
	Date     time.Time `json:"date"`
}

type History struct {
	Results []Result `json:"results"`
}

func path() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".termtype.json")
}

func Load() History {
	data, err := os.ReadFile(path())
	if err != nil {
		return History{}
	}
	var h History
	json.Unmarshal(data, &h)
	return h
}

func (h *History) Save() error {
	data, err := json.MarshalIndent(h, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path(), data, 0644)
}

func (h *History) Add(r Result) {
	h.Results = append(h.Results, r)
	h.Save()
}

func (h *History) BestWPM() float64 {
	best := 0.0
	for _, r := range h.Results {
		if r.WPM > best {
			best = r.WPM
		}
	}
	return best
}

func (h *History) Last(n int) []Result {
	if len(h.Results) <= n {
		return h.Results
	}
	return h.Results[len(h.Results)-n:]
}

func (h *History) BestWPMByMode(mode string) float64 {
	best := 0.0
	for _, r := range h.Results {
		if r.Mode == mode && r.WPM > best {
			best = r.WPM
		}
	}
	return best
}

func (h *History) RecentByMode(mode string, n int) []Result {
	var filtered []Result
	for _, r := range h.Results {
		if r.Mode == mode {
			filtered = append(filtered, r)
		}
	}
	sort.Slice(filtered, func(i, j int) bool {
		return filtered[i].Date.Before(filtered[j].Date)
	})
	if len(filtered) <= n {
		return filtered
	}
	return filtered[len(filtered)-n:]
}
