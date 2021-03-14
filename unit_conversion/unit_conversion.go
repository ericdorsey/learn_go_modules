package unit_conversion

import (
    "fmt"
)

type Feet float64
type Meter float64
//type Yard float64
//type Pound float64
//type Kilogram float64

// Go Programming Language; Type Declaration Methods p. 40
// type Feet method
func (feet Feet) String() string {
    return fmt.Sprintf("%f feet", feet) 
}

//type Meter method
func (meter Meter) String() string {
    return fmt.Sprintf("%f meter(s)", meter)
}

// Converts feet to meters
func FeetToMeters(feet Feet) Meter {
    meters_result := feet / 3.28
    // Conversion to Meter, not a function call
    return Meter(meters_result)
}

// Prints a user friendly Feet to Meters conversion
func FeetToMetersPretty(feet Feet) {
    meters_result := FeetToMeters(feet)
    if feet < 2 {
        fmt.Printf("%.0f feet is equal to %.2f meters\n", feet, meters_result)
    } else {
        fmt.Printf("%.0f feet is equal to %.2f meters\n", feet, meters_result)
    }
}

// Converts meters to feet
func MetersToFeet(meter Meter) Feet {
    feet_result := meter * 3.28
    return Feet(feet_result)
}

// Prints a user friendly Meters to Feet conversion
func MetersToFeetPretty(meter Meter) {
    feet_result := MetersToFeet(meter)
    if meter < 2 {
        fmt.Printf("%.0f meter is equal to %.2f feet\n", meter, feet_result)
    } else {
        fmt.Printf("%.0f meters is equal to %.2f feet\n", meter, feet_result)
    }
}
