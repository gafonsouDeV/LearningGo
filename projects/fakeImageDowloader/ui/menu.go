package ui

import (
	"time"

	"sync"

	"math/rand"
	"strconv"

	"github.com/gafonsouDeV/LearningGo/projects/fakeImageDowloader/utils"
)

func Menu() {
	println("How many images do you want to download?")
	wg := &sync.WaitGroup{}

	imagesToDownload := utils.ReadNumberInput()

	images := make([]string, imagesToDownload)

	// creation of the array of images
	for index := 0; index < imagesToDownload; index++ {
		imageName := "Image" + strconv.Itoa(index+1) + ".png"
		images[index] = imageName
	}

	println("Starting downloads...")
	for _, image := range images {
		wg.Add(1)
		go download(image, wg)
	}

	wg.Wait()
	println("All images dowloaded")
}

func download(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	println("Dowloading " + url + "...")
	sleepTime := rand.Intn(3)
	time.Sleep(time.Duration(sleepTime) * time.Second)
	println(url + " dowloaded.")
}
