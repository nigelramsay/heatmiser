# HeatMiser

Golang library for HeatMiser Wifi Thermostat

## Usage

```
package main

import "github.com/nigelramsay/heatmiser"
import "fmt"

func main() {
  h := heatmiser.New("http://192.168.1.7")

  fmt.Println("Current temperature: ", h.Current())
  fmt.Println("Target temperature: ", h.Target())
  fmt.Println("Heating enabled: ", h.Enabled())
}
```

## Compatibility

This library is compatible with the following firmware revisions:

- version 1.3

## Authentication

The wifi thermostat uses Javascript to enforce the authentication, which
essentially means that there is no authentication. 
