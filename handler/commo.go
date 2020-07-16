package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

/*
	WIP:
	implement conversion to USD
	implement aggregation
	implement jwt auth
	implement caching
*/

func getConvertionRate(c echo.Context) float64 {
	var price map[string]interface{}
	var client = &http.Client{}
	const urlRate = "https://free.currconv.com/api/v7/convert?q=IDR_USD&compact=ultra&apiKey=8090322b8ffdddb63d5a"
	req, err := http.NewRequest("GET", urlRate, nil)
	if err != nil {
		log.Fatalln("error_fetch_remote_url")
	}

	rsp, err := client.Do(req)
	if err != nil {
		log.Fatalln("error_fetch_remote_url")
	}
	defer rsp.Body.Close()

	err = json.NewDecoder(rsp.Body).Decode(&price)
	if err != nil {
		log.Fatalln("error_fetch_remote_url")
	}
	rate, ok := price["IDR_USD"].(float64)
	if !ok {
		log.Fatalln("error_casting_result")
	}
	return rate
}

func getCommodityData(c echo.Context) []map[string]interface{} {
	var commo []map[string]interface{}
	var client = &http.Client{}
	const urlFish = "https://stein.efishery.com/v1/storages/5e1edf521073e315924ceab4/list"
	req, err := http.NewRequest("GET", urlFish, nil)
	if err != nil {
		return nil
	}

	rsp, err := client.Do(req)
	if err != nil {
		return nil
	}
	defer rsp.Body.Close()

	err = json.NewDecoder(rsp.Body).Decode(&commo)
	if err != nil {
		return nil
	}

	return commo
}

// Fetch from another API
func Fetch(c echo.Context) error {
	// rate := getConvertionRate(c)

	commo := getCommodityData(c)
	log.Println(commo[0])
	price, ok := commo[0]["price"].(int64)
	if !ok {
		log.Fatalln("error_casting")
		return c.JSON(http.StatusInternalServerError, "error_casting")
	}
	log.Print(price)

	return c.JSON(http.StatusOK, commo)
}
