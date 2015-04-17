# HeatMiser

Golang library for HeatMiser Wifi Thermostat

## Usage

```
import "github.com/nigelramsay/heatmiser"

h := HeatMiser{baseURL: "http://192.168.1.7"}
h.Perform()

fmt.Println("Actual: ", h.Current())
fmt.Println("Target: ", h.Target())
fmt.Println("Enabled: ", h.Enabled())
```
