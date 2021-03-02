package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"test-go-server/models"
	"time"

	"github.com/gin-gonic/gin"
)

func GetValueByObject(c *gin.Context) {

	utime := int(time.Now().Unix())
	fmt.Println(utime)

	type Data struct {
		Main  int `json:"Main"`
		Sub   int `json:"Sub"`
		Value int `json:"Value"`
	}

	type Response struct {
		Utime    int    `json:"utime"`
		Response []Data `json:"data"`
	}
	//res1 := Response{}
	data1 := Data{}
	main := c.Param("main")

	q := c.Request.URL.Query()
	sub := q.Get("sub")

	// Open our jsonFile
	jsonFile, err := os.Open("./data/Format_" + main + ".json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened Format_101.json")
	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var formats models.Format

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &formats)

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	if len(sub) > 0 {
		// ... process it, will be the first (only) if multiple were given
		// note: if they pass in like ?param1=&param2= param1 will also be "" :|
		sub_i, err := strconv.Atoi(sub)
		main_i, err := strconv.Atoi(main)

		fmt.Println("Res1")

		fmt.Println("Res")
		if err != nil {
			fmt.Println(err)
		}
		for i := range formats.Format {
			if formats.Format[i].Id == sub_i {
				fmt.Println(formats.Format[i])
				data1.Main = main_i
				data1.Sub = sub_i
				data1.Value = formats.Format[i].Value
				var data_arr []Data
				data_arr = append(data_arr, data1)

				jsonResponse := Response{
					Utime:    utime,
					Response: data_arr,
				}

				if err != nil {
					fmt.Println(err)
				}
				c.JSON(http.StatusOK, jsonResponse)

			}
		}
	} else {
		fmt.Println("No Sub")
		c.JSON(http.StatusOK, formats)
	}

}

func GetFormatByObject(c *gin.Context) {

	utime := int(time.Now().Unix())
	fmt.Println(utime)

	type Data struct {
		Main   int `json:"Main"`
		Sub    int `json:"Sub"`
		Format int `json:"Format"`
	}

	type Response struct {
		Utime    int    `json:"utime"`
		Response []Data `json:"data"`
	}
	//res1 := Response{}
	data1 := Data{}
	main := c.Param("main")

	q := c.Request.URL.Query()
	sub := q.Get("sub")

	// Open our jsonFile
	jsonFile, err := os.Open("./data/Format_" + main + ".json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened Format_101.json")
	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var formats models.Format

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &formats)

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	if len(sub) > 0 {
		// ... process it, will be the first (only) if multiple were given
		// note: if they pass in like ?param1=&param2= param1 will also be "" :|
		sub_i, err := strconv.Atoi(sub)
		main_i, err := strconv.Atoi(main)

		fmt.Println("Res1")

		fmt.Println("Res")
		if err != nil {
			fmt.Println(err)
		}
		for i := range formats.Format {
			if formats.Format[i].Id == sub_i {
				fmt.Println(formats.Format[i])
				data1.Main = main_i
				data1.Sub = sub_i
				data1.Format = formats.Format[i].Type
				var data_arr []Data
				data_arr = append(data_arr, data1)

				jsonResponse := Response{
					Utime:    utime,
					Response: data_arr,
				}

				if err != nil {
					fmt.Println(err)
				}
				c.JSON(http.StatusOK, jsonResponse)

			}
		}
	} else {
		fmt.Println("No Sub")
		c.JSON(http.StatusOK, formats)
	}

}
