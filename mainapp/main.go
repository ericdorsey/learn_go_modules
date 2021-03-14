package main

import (
    "fmt"
    "fake.com/unit_conversion/unit_conversion"
    //"unit_conversion"
)

func main() {
    //oneMeterToFeet := unit_conversion.MetersToFeet(1)
    //fmt.Printf("%v\n", oneMeterToFeet)
    unit_conversion.FeetToMetersPretty(1)
    unit_conversion.FeetToMetersPretty(2)
    fmt.Println()
    unit_conversion.MetersToFeetPretty(1)
    unit_conversion.MetersToFeetPretty(2)
}
