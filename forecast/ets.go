package forecast

import (
	"context"
	"fmt"
	// "errors"

	"github.com/bradfitz/iter"
	"github.com/rocketlaunchr/dataframe-go"
)

// SimpleExponentialSmoothing method calculates
// and returns forecast for future m periods
//
//// s - dataframe.SeriesFloat64 object
//
// y - Time series data gotten from s.
// alpha - Exponential smoothing coefficients for level, trend,
//        seasonal components.
// m - Intervals into the future to forecast
//
// https://www.itl.nist.gov/div898/handbook/pmc/section4/pmc431.htm
// newvalue = smoothing * next + (1 - smoothing)*old value
// forecast[i+1] = St[i] + alpha * ϵt,
// where ϵt is the forecast error (actual - forecast) for period i.
func SimpleExponentialSmoothing(ctx context.Context, s *dataframe.SeriesFloat64, α float64, m int, r ...dataframe.Range) (*dataframe.SeriesFloat64, error) {

	if len(r) == 0 {
		r = append(r, dataframe.Range{})
	}

	start, end, err := r[0].Limits(len(s.Values))
	if err != nil {
		return nil, err
	}

	// TODO: add validation

	forecast := make([]float64, 0, m)
	var st float64
	for i := start; i < end+1; i++ {
		xt := s.Values[i]

		if i == start {
			st = xt
		} else {
			st = α*xt + (1-α)*st
		}
		fmt.Println("st", st)
	}

	// Now calculate forecast
	for range iter.N(m) {
		st = α*s.Values[end] + (1-α)*st
		forecast = append(forecast, st)
	}

	fdf := dataframe.NewSeriesFloat64("forecast", nil)
	fdf.Values = forecast

	return fdf, nil

	//////////

	// if len(r) == 0 {
	// 	r = append(r, dataframe.Range{})
	// }

	// start, end, err := r[0].Limits(len(s.Values))
	// if err != nil {
	// 	return nil, err
	// }
	// // inclusive of value at end index
	// y := s.Values[start : end+1]

	// // Validating arguments
	// if len(y) == 0 {
	// 	return nil, errors.New("value of y should be not null")
	// }

	// if m <= 0 {
	// 	return nil, errors.New("value of m must be greater than 0")
	// }

	// if m > len(y) {
	// 	return nil, errors.New("value of m can not be greater than length of y")
	// }

	// if (alpha < 0.0) || (alpha > 1.0) {
	// 	return nil, errors.New("value of Alpha should satisfy 0.0 <= alpha <= 1.0")
	// }

	// st := make([]float64, len(y))
	// forecast := make([]float64, m)

	// // Set initial value to first element in y
	// st[1] = y[0]

	// // start smoothing from the third element
	// for i := 2; i < len(y); i++ {

	// 	// Exiting on context error
	// 	if err := ctx.Err(); err != nil {
	// 		return nil, err
	// 	}

	// 	// simple exponential Smoothing
	// 	st[i] = alpha*y[i-1] + ((1.0 - alpha) * st[i-1])

	// 	// separating forecast from smoothing process
	// 	// forecast
	// 	for j := 0; j < m; j++ {
	// 		// 'pt' serves as reference point to start forecasting from from the y set of data passed in
	// 		pt := len(y) - m
	// 		// forecast[i+m] = st[i] + (alpha * (y[i] - st[i]))
	// 		forecast[j] = alpha*y[pt+j] + (1.0-alpha)*st[pt+j]
	// 	}

	// }

	// init := &dataframe.SeriesInit{}

	// seriesForecast := dataframe.NewSeriesFloat64("Forecast", init)

	// // Load forecast data into series
	// seriesForecast.Insert(seriesForecast.NRows(), forecast[:])

	// return seriesForecast, nil

}
