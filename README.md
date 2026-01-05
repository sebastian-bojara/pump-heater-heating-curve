# Heat Pump Heating Curve

A CLI application in Go for generating a heating curve table for heat pumps.

## Description

The application calculates and displays a table showing the relationship between heating temperature and outdoor temperature. The heating curve is linear - the lower the outdoor temperature, the higher the required heating temperature.

### Parameters

- `-min-ext` - Minimum outdoor temperature in °C (default: -20)
- `-max-ext` - Maximum outdoor temperature in °C (default: 15)
- `-min-heat` - Minimum heating temperature in °C (default: 25)
- `-max-heat` - Maximum heating temperature in °C (default: 50)

The temperature step is fixed at 1°C.

### Examples

Run with default parameters:

```bash
make run
```

Run with custom parameters:
```bash
make run ARGS="-min-ext -20 -max-ext 15 -min-heat 20 -max-heat 40"
```

## Example Output

```
=== Heat Pump Heating Curve ===
Outdoor temperature range: -20°C to 15°C
Heating temperature range: 25°C to 50°C

┌──────────────┬──────────────┐
│ Outdoor (°C) │ Heating (°C) │
├──────────────┼──────────────┤
│          -20 │           55 │
│          -19 │           54 │
├──────────────┼──────────────┤
│          -18 │           53 │
│          -17 │           52 │
├──────────────┼──────────────┤
...
│           13 │           22 │
│           14 │           21 │
├──────────────┼──────────────┤
│           15 │           20 │
└──────────────┴──────────────┘
```
