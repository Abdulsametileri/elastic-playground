package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	host := "http://localhost:9200"
	language := "tr"
	slug := "test2"

	type queryAndGroupId struct {
		Query           string `json:"query"`
		SolrPageGroupID int    `json:"solrPageGroupId"`
	}

	addedDoc := queryAndGroupId{
		Query:           "query gelicek buraya",
	}

	bd := fmt.Sprintf(`{"query":"%s", "solrPageGroupId":"%d"}`, addedDoc.Query, addedDoc.SolrPageGroupID)

	client := http.Client{}

	queryUrl := fmt.Sprintf("%s/lucene_query-%s/_doc/", host, language)
	req, _ := http.NewRequest(http.MethodPut, fmt.Sprintf("%s%s", queryUrl, slug), strings.NewReader(bd))
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode >= 300 {
		body, _ := ioutil.ReadAll(res.Body)
		log.Fatalf("add query with lang %s and slug %s to es request respond with status %d with body %+v", language, slug, res.StatusCode, string(body))
	}

	log.Println("başarılı")
}
