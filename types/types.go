package types

import "reflect"

type BalloonObject struct {
	Name       string
	Weight     string
	Balloons   string
	WeightType string
}

type Conversion struct {
	Pound       float64
	Ounce       float64
	UsTon       float64
	ImperialTon float64
	Microgram   float64
	Milligram   float64
	Gram        float64
	Kilogram    float64
	MetricTon   float64
}

func GetConversionField(c *Conversion, field string) float64 {
	r := reflect.ValueOf(c)
	f := reflect.Indirect(r).FieldByName(field)
	return f.Float()
}
