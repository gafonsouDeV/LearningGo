package services

import (
	"context"
	"math/rand"
	"time"

	"github.com/gafonsouDeV/LearningGo/projects/fakeWeatherService/models"
	"github.com/gafonsouDeV/LearningGo/projects/fakeWeatherService/utils"
)

// option without context
//
// func FetchWeather(name string, result chan<- models.WeatherResult, done <-chan struct{}) {
// 	time.Sleep(utils.RandDelay())

// 	res := models.WeatherResult{Service: name, Temp: 20.0 + rand.Float32()*10}

// 	select {
// 	case <-done:
// 		return
// 	case result <- res:
// 		return
// 	}
// }

func FetchWithContext(ctx context.Context, name string, results chan<- models.WeatherResult) {
	select {
	case <-time.After(utils.RandDelay()):
	case <-ctx.Done():
		return
	}
	select {
	case <-ctx.Done():
		return
	case results <- models.WeatherResult{Service: name, Temp: 20.0 + rand.Float32()*10}:
	}
}
