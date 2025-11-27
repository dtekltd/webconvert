package utils

import "strconv"

func IntWithDefault(val int, defaultVal int) int {
	if val != 0 {
		return val
	}
	return defaultVal
}

func FloatWithDefault(val float64, defaultVal float64) float64 {
	if val != 0 {
		return val
	}
	return defaultVal
}

func ParseInt(s string, defaultValue int) int {
	if s == "" {
		return defaultValue
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		return defaultValue
	}
	return i
}

func ParseFloat(s string, defaultValue float64) float64 {
	if s == "" {
		return defaultValue
	}
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return defaultValue
	}
	return f
}
