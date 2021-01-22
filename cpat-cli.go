package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

/** Getting Started
	Rapid7 - Simple CLI tool in golang: https://blog.rapid7.com/2016/08/04/build-a-simple-cli-tool-with-golang/
	Consuming a REST API in golang: https://tutorialedge.net/golang/consuming-restful-api-with-go/
*/

/**
	GET request
*/
func restRequest() {
	// Make an HTTP GET request
	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(responseData))
}

/**
	POST request
	ref: https://medium.com/@masnun/making-http-requests-in-golang-dd123379efe7
*/
func postRequest() {
	requestBody, err := json.Marshal(map[string]string{
		"someKey": "someValue",
		"anoKey": "otherValue"
	})

	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post("http://url", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))
}

func main() {
	// call our function, which then makes a REST request
	restRequest()
}

/*
	Company

	http://localhost:{port}/api/Mongo/Company/Get/{id?}

	http://localhost:{port}/api/Mongo/Company/Page
	-> [body: PageRequest pageDock]

	http://localhost:{port}/api/Mongo/Company/Insert
	-> [body: Company data]

	http://localhost:{port}/api/Mongo/Company/Update/{id?}
	-> [body: Company data]

	http://localhost:{port}/api/Mongo/Company/PartialUpdate/{id?}
	-> [body: JsonMergePatchDocument<Company> patch]

	http://localhost:{port}/api/Mongo/Company/Remove/{id?}
*/

/*
	Device

	http://localhost:{port}/api/Mongo/Device/Get/{id?}

	http://localhost:{port}/api/Mongo/Device/Page
	-> [body: PageRequest pageDock]

	http://localhost:{port}/api/Mongo/Device/Insert
	-> [body: Device data]

	http://localhost:{port}/api/Mongo/Device/Update/{id?}
	-> [body: Device data]

	http://localhost:{port}/api/Mongo/Device/PartialUpdate/{id?}
	-> [body: JsonMergePatchDocument<Device> patch]

	http://localhost:{port}/api/Mongo/Device/Remove/{id?}
*/

/*
	Location

	http://localhost:{port}/api/Mongo/Location/Get/{id?}

	http://localhost:{port}/api/Mongo/Location/Page
	-> [body: PageRequest pageDock]

	http://localhost:{port}/api/Mongo/Location/Insert
	-> [body: Location data]

	http://localhost:{port}/api/Mongo/Device/Update/{id?}
	-> [body: Location data]

	http://localhost:{port}/api/Mongo/Location/PartialUpdate/{id?}
	-> [body: JsonMergePatchDocument<Location> patch]

	http://localhost:{port}/api/Mongo/Location/Remove/{id?}
*/

/*
	Person

	http://localhost:{port}/api/Mongo/Person/Get/{id?}

	http://localhost:{port}/api/Mongo/Person/Page
	-> [body: PageRequest pageDock]

	http://localhost:{port}/api/Mongo/Person/Insert
	-> [body: Person data]

	http://localhost:{port}/api/Mongo/Device/Update/{id?}
	-> [body: Person data]

	http://localhost:{port}/api/Mongo/Person/PartialUpdate/{id?}
	-> [body: JsonMergePatchDocument<Person> patch]

	http://localhost:{port}/api/Mongo/Person/Remove/{id?}
*/

/*
	Target

	http://localhost:{port}/api/Mongo/Target/Get/{id?}

	http://localhost:{port}/api/Mongo/Target/Page
	-> [body: PageRequest pageDock]

	http://localhost:{port}/api/Mongo/Target/Insert
	-> [body: Target data]

	http://localhost:{port}/api/Mongo/Device/Update/{id?}
	-> [body: Target data]

	http://localhost:{port}/api/Mongo/Target/PartialUpdate/{id?}
	-> [body: JsonMergePatchDocument<Target> patch]

	http://localhost:{port}/api/Mongo/Target/Remove/{id?}
*/
