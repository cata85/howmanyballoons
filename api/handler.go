package api

import (
	"strconv"

	"github.com/cata85/balloons/types"
)

var gramsPerBalloon = 14.0
var conversions *types.Conversion

func Calculate(weight string, weightType string) string {
	if conversions == nil {
		conversions = new(types.Conversion)
		conversions.Pound = 0.00220462
		conversions.Ounce = 0.035274
		conversions.UsTon = 1.1023e-6
		conversions.ImperialTon = 9.8421e-7
		conversions.Microgram = 1e+6
		conversions.Milligram = 1000
		conversions.Gram = 1.0
		conversions.Kilogram = 0.001
		conversions.MetricTon = 1e-6
	}
	weightConversion := types.GetConversionField(conversions, weightType)

	weightOriginal, _ := strconv.ParseFloat(weight, 64)
	weightGrams := weightOriginal / float64(weightConversion)
	balloons := weightGrams / gramsPerBalloon
	return strconv.FormatFloat(balloons, 'f', -1, 64)
}
