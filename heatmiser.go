package heatmiser

import "net/http"
import "io/ioutil"
import "regexp"
import "strconv"

type HeatMiser struct {
  baseURL string
  body []byte
}

func New(baseURL string) *HeatMiser {
  h := HeatMiser{baseURL: baseURL}
  h.Perform()
  return &h
}

func (h *HeatMiser) Perform() {
  resp, _ := http.Get(h.baseURL + "/right.htm")
  defer resp.Body.Close()

  h.body, _ = ioutil.ReadAll(resp.Body)
}

func (h *HeatMiser) Current() float64 {
  current := h.dataExtract(`<b>Actual.*font size='5'>(\d{1,2}\.\d)`)
  val, _ := strconv.ParseFloat(current, 64)
  return val
}

func (h *HeatMiser) Target() float64 {
  target := h.dataExtract(`<b>Set.*font size='4'>(\d{1,2})`)
  val, _ := strconv.ParseFloat(target, 64)
  return val
}

func (h *HeatMiser) Enabled() bool {
  onOff := h.dataExtract(`<b>Heat Status.*font size='4'>(O[NF]{1,2})`)
  return onOff == `ON`
}

func (h *HeatMiser) dataExtract(regex string) string {
  r := regexp.MustCompile(regex)
  result := r.FindSubmatch(h.body)
  return string(result[1])
}
