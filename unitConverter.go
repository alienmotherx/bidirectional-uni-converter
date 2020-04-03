package main

import (
	"fmt"
	"math"
)

// BI-DIRECTIONAL UNIT CONVERTER

type Converter struct{}

type (
	Feet         float64
	Centimeter   float64
	Minutes      float64
	Seconds      float64
	Milliseconds float64
	Celsius      float64
	Fahrenheit   float64
	Radian       float64
	Degree       float64
	Kilogram     float64
	Pounds       float64
	Kelvin       float64
	Meter        float64
	Kilometer    float64
	Mile         float64
	Yard         float64
	Ounce        float64
	Karat        float64
)

//Converter for Feet and Centimeter
func (cvr Converter) FeetToCentimeter(c Feet) Centimeter {
	return Centimeter(c * Feet(30.48))
}

func (cvr Converter) CentimeterToFeet(c Centimeter) Feet {
	return Feet(c / Centimeter(30.48))
}

//Converter for Minutes and Seconds
func (cvr Converter) MinutesToSeconds(c Minutes) Seconds {
	return Seconds(c * Minutes(60))
}

func (cvr Converter) SecondsToMinutes(c Seconds) Minutes {
	return Minutes(c / Seconds(60))
}

//Converter for Seconds and Milliseconds
func (cvr Converter) SecondsToMilliseconds(c Seconds) Milliseconds {
	return Milliseconds(c * Seconds(1000))
}

func (cvr Converter) MillisecondsToSeconds(c Milliseconds) Seconds {
	return Seconds(c / Milliseconds(1000))
}

//Converter for Celius and Fehreanheit
func (cvr Converter) CelsiusToFehrenheit(c Celsius) Fahrenheit {
	return Fahrenheit((c * Celsius(9/5)) + Celsius(32))
}

func (cvr Converter) FehrenheitToCelsius(c Fahrenheit) Celsius {
	return Celsius((c - Fahrenheit(32)) * Fahrenheit(5/9))
}

//Converter for Radian and Degree
func (cvr Converter) RadianToDegree(c Radian) Degree {
	return Degree(c * (Radian(180) / math.Pi))
}

func (cvr Converter) DegreeToRadian(c Degree) Radian {
	return Radian((c * math.Pi) / Degree(180))
}

//Converter for Kilogram and Pounds
func (cvr Converter) KilogramToPounds(c Kilogram) Pounds {
	return Pounds(c * Kilogram(2.0462))
}

func (cvr Converter) PoundsToKilogram(c Pounds) Kilogram {
	return Kilogram(c / Pounds(2.0462))
}

//Converter for Celsius and Kelvin
func (cvr Converter) CelsiusToKelvin(c Celsius) Kelvin {
	return Kelvin(c + Celsius(273.15))
}

func (cvr Converter) KelvinToCelsius(c Kelvin) Celsius {
	return Celsius(c - Kelvin(273.15))
}

//Converter for Meter and Feet
func (cvr Converter) MeterToFeet(c Meter) Feet {
	return Feet(c * Meter(3.28084))
}

func (cvr Converter) FeetToMeter(c Feet) Meter {
	return Meter(c / Feet(3.28084))
}

//Converter for Meter and Kilometer
func (cvr Converter) MeterToKilometer(c Meter) Kilometer {
	return Kilometer(c / Meter(1000))
}

func (cvr Converter) KilometerToMeter(c Kilometer) Meter {
	return Meter(c * Kilometer(1000))
}

//Converter for Mile and Yard
func (cvr Converter) MileToYard(c Mile) Yard {
	return Yard(c * Mile(1760))
}

func (cvr Converter) YardToMile(c Yard) Mile {
	return Mile(c / Yard(1760))
}

//Converter for Ounce and Karat
func (cvr Converter) OunceToKarat(c Ounce) Karat {
	return Karat(c * Ounce(141.748))
}

func (cvr Converter) KaratToOunce(c Karat) Ounce {
	return Ounce(c / Karat(1760))
}

func main() {
	cvr := Converter{}
	fmt.Println("Feet To Centimeter:", cvr.FeetToCentimeter(100))
	fmt.Println("Centimeter To Feet:", cvr.CentimeterToFeet(60))
	fmt.Println("Minutes To Seconds:", cvr.MinutesToSeconds(50))
	fmt.Println("Seconds To Minutes:", cvr.SecondsToMinutes(120))
	fmt.Println("Seconds to Milliseconds:", cvr.SecondsToMilliseconds(6090))
	fmt.Println("Milliseconds To Seconds:", cvr.MillisecondsToSeconds(13000))
	fmt.Println("Fehrenheit To Celsius:", cvr.FehrenheitToCelsius(345))
	fmt.Println("Celsius To Fehrenheit:", cvr.CelsiusToFehrenheit(234))
	fmt.Println("Radian To Degree:", cvr.RadianToDegree(243))
	fmt.Println("Degree To Radian:", cvr.DegreeToRadian(180))
	fmt.Println("Kilogram To Pounds:", cvr.KilogramToPounds(678))
	fmt.Println("Pounds To Kilogram:", cvr.PoundsToKilogram(333))
	fmt.Println("Celsius To Kelvin:", cvr.CelsiusToKelvin(880))
	fmt.Println("Kelvin To Celsius:", cvr.KelvinToCelsius(768))
	fmt.Println("Meter To Feet:", cvr.MeterToFeet(1000))
	fmt.Println("Feet To Meter:", cvr.FeetToMeter(10))
	fmt.Println("Meter To Kilometer:", cvr.MeterToKilometer(5009))
	fmt.Println("Kilometer To Meter:", cvr.KilometerToMeter(789))
	fmt.Println("Mile To Yard:", cvr.MileToYard(100))
	fmt.Println("Yard To Mile:", cvr.YardToMile(15))
	fmt.Println("Ounce To Karat:", cvr.OunceToKarat(60))
	fmt.Println("Karat To Ounce:", cvr.KaratToOunce(24))
}
