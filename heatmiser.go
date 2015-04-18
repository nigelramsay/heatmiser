package heatmiser

import "net/http"
import "io/ioutil"
import "regexp"
import "strconv"

// HeatMiser holds the state of the HeatMiser wifi thermostat
type HeatMiser struct {
	baseURL string
	body    []byte
}

// New creates a new HeatMiser struct
func New(baseURL string) *HeatMiser {
	h := HeatMiser{baseURL: baseURL}
	h.perform()
	return &h
}

// perform downloads the current state from the HeatMiser thermostat
func (h *HeatMiser) perform() {
	resp, _ := http.Get(h.baseURL + "/right.htm")
	defer resp.Body.Close()

	h.body, _ = ioutil.ReadAll(resp.Body)
}

// Current returns the current reported temperature
func (h *HeatMiser) Current() float64 {
	current := h.dataExtract(`<b>Actual.*font size='5'>(\d{1,2}\.\d)`)
	val, _ := strconv.ParseFloat(current, 64)
	return val
}

// Target returns the target desired temperature
func (h *HeatMiser) Target() float64 {
	target := h.dataExtract(`<b>Set.*font size='4'>(\d{1,2})`)
	val, _ := strconv.ParseFloat(target, 64)
	return val
}

// Enabled returns true if the HeatMiser has enabled the burner
func (h *HeatMiser) Enabled() bool {
	onOff := h.dataExtract(`<b>Heat Status.*font size='4'>(O[NF]{1,2})`)
	return onOff == `ON`
}

func (h *HeatMiser) dataExtract(regex string) string {
	r := regexp.MustCompile(regex)
	result := r.FindSubmatch(h.body)
	return string(result[1])
}
