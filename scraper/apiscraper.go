package main

import (
	"image"
	"net/http"
	"strconv"
)

const apiKey string = "AIzaSyBMqdtriyHGTZHOsD2x-EJzqsK3N9PlPC4"

func getImage(location string, w int, h int, zoom int) (image.Image, error) {
	URL := makeURLFromParams(location, w, h, zoom)
	response, err := http.Get(URL)
	handleErr(err)

	return readImageFromResponse(response)

}

func readImageFromResponse(response *http.Response) (image.Image, error) {
	defer response.Body.Close()
	img, _, err := image.Decode(response.Body)
	return img, err
}

func makeURLFromParams(location string, w int, h int, zoom int) string {
	baseURL := "https://maps.googleapis.com/maps/api/staticmap?"
	sizeString := strconv.Itoa(w) + "x" + strconv.Itoa(h)
	return baseURL + "center=" + location +
		"&zoom=" + strconv.Itoa(zoom) +
		"&size=" + sizeString +
		"&maptype=satellite" +
		"&key=" + apiKey

}
