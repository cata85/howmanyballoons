package api

import (
	"math"
	"os"
	"strconv"

	helper "github.com/cata85/balloons/helpers"
	"github.com/cata85/balloons/types"
	"github.com/gorilla/sessions"
)

var user *types.User
var balloonObject *types.BalloonObject
var store *sessions.CookieStore

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
	if weight == "" {
		return ""
	}
	weightConversion := types.GetConversionField(conversions, weightType)

	weightOriginal, _ := strconv.ParseFloat(weight, 64)
	weightGrams := weightOriginal / float64(weightConversion)
	balloons := weightGrams / gramsPerBalloon
	return strconv.FormatInt(int64(math.Ceil(balloons)), 10)
}

func Initialize() {
	if user == nil {
		user = new(types.User)
	}
	if store == nil {
		config := helper.Config()
		gorillaConfig := config["gorilla"]
		store = sessions.NewCookieStore([]byte(os.Getenv(helper.String(helper.Get(gorillaConfig, "key")))))
	}
}
