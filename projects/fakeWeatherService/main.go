package main

import (
	"context"
	"fmt"
	"time"

	"github.com/gafonsouDeV/LearningGo/projects/fakeWeatherService/models"
	"github.com/gafonsouDeV/LearningGo/projects/fakeWeatherService/services"
)

func main() {
	weatherServices := []string{"SvcA", "SvB", "SvC"}
	result := make(chan models.WeatherResult)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for _, service := range weatherServices {
		go services.FetchWithContext(ctx, service, result)
	}

	select {
	case res := <-result:
		fmt.Println("Weather Service:", res)
		cancel()
	case <-time.After(2 * time.Second):
		fmt.Println("Timeout")
	}

}
