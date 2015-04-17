# HeatMiser

Golang library for HeatMiser Wifi Thermostat

## Usage

```
import "github.com/nigelramsay/heatmiser"

h := HeatMiser{baseURL: "http://192.168.1.7"}
h.Perform()

fmt.Println("Current temperature: ", h.Current())
fmt.Println("Target temperature: ", h.Target())
fmt.Println("Heating enabled: ", h.Enabled())
```
