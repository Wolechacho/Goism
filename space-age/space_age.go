package space

const orbitalperiod int = 31557600

// Planet is of type string
type Planet string

// Age calculate the given age in different planet
func Age(seconds float64, planet Planet) float64 {
	calculatedAge := 0.0
	switch planet {
	case "Earth":
		calculatedAge = seconds / float64(orbitalperiod)
	case "Mercury":
		calculatedAge = seconds / (float64(orbitalperiod) * 0.2408467)
	case "Venus":
		calculatedAge = seconds / (float64(orbitalperiod) * 0.61519726)
	case "Mars":
		calculatedAge = seconds / (float64(orbitalperiod) * 1.8808158)
	case "Jupiter":
		calculatedAge = seconds / (float64(orbitalperiod) * 11.862615)
	case "Saturn":
		calculatedAge = seconds / (float64(orbitalperiod) * 29.447498)
	case "Uranus":
		calculatedAge = seconds / (float64(orbitalperiod) * 84.016846)
	case "Neptune":
		calculatedAge = seconds / (float64(orbitalperiod) * 164.79132)
	}
	return calculatedAge

}
