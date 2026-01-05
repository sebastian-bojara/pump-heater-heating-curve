package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Define flags for input parameters
	minExtTemp := flag.Int("min-ext", -20, "Minimum outdoor temperature (°C)")
	maxExtTemp := flag.Int("max-ext", 15, "Maximum outdoor temperature (°C)")
	minHeatTemp := flag.Int("min-heat", 25, "Minimum heating temperature (°C)")
	maxHeatTemp := flag.Int("max-heat", 50, "Maximum heating temperature (°C)")

	flag.Parse()

	// Validate parameters
	if *minExtTemp >= *maxExtTemp {
		fmt.Fprintln(os.Stderr, "Error: Minimum outdoor temperature must be less than maximum")
		os.Exit(1)
	}

	if *minHeatTemp >= *maxHeatTemp {
		fmt.Fprintln(os.Stderr, "Error: Minimum heating temperature must be less than maximum")
		os.Exit(1)
	}

	// Display parameters
	fmt.Println("=== Heat Pump Heating Curve ===")
	fmt.Printf("Outdoor temperature range: %d°C to %d°C\n", *minExtTemp, *maxExtTemp)
	fmt.Printf("Heating temperature range: %d°C to %d°C\n\n", *minHeatTemp, *maxHeatTemp)

	// Generate table
	fmt.Println("┌──────────────┬──────────────┐")
	fmt.Println("│ Outdoor (°C) │ Heating (°C) │")
	fmt.Println("├──────────────┼──────────────┤")

	rowCount := 0
	for extTemp := *minExtTemp; extTemp <= *maxExtTemp; extTemp++ {
		heatTemp := calculateHeatingTemp(extTemp, *minExtTemp, *maxExtTemp, *minHeatTemp, *maxHeatTemp)
		fmt.Printf("│ %12d │ %12d │\n", extTemp, heatTemp)
		rowCount++
		if rowCount%2 == 0 && extTemp+1 <= *maxExtTemp {
			fmt.Println("├──────────────┼──────────────┤")
		}
	}

	fmt.Println("└──────────────┴──────────────┘")
}

// calculateHeatingTemp calculates the heating temperature based on the heating curve
// The curve is linear: lower outdoor temperature = higher heating temperature
func calculateHeatingTemp(extTemp, minExtTemp, maxExtTemp, minHeatTemp, maxHeatTemp int) int {
	// Linear interpolation
	// When extTemp = minExtTemp, return maxHeatTemp
	// When extTemp = maxExtTemp, return minHeatTemp
	ratio := float64(extTemp-minExtTemp) / float64(maxExtTemp-minExtTemp)
	heatTemp := float64(maxHeatTemp) - ratio*float64(maxHeatTemp-minHeatTemp)

	return int(heatTemp + 0.5) // Round to nearest integer
}
