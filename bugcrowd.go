package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func bugcrowd(conn *sql.DB) {
	BugcrowdURL := "https://raw.githubusercontent.com/Osb0rn3/bugbounty-targets/main/programs/bugcrowd.json"
	response, err := http.Get(BugcrowdURL)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var data []map[string]interface{}
	err = json.Unmarshal(body, &data)
	for _, program := range data {
		name := program["name"].(string)
		submission := fmt.Sprintf("https://bugcrowd.com%v", program["report_path"])
		platform := "bugcrowd"
		logo := program["logo"].(string)
		targets := program["target_groups"].([]interface{})
		bounty := "false"
		if _, exists := program["min_rewards"]; exists {
			bounty = "true"
		} else {
			bounty = "false"
		}
		PK := crawlPrograms(conn, name, submission, platform, bounty, logo)
		for _, target := range targets {
			if target.(map[string]interface{})["in_scope"].(bool) == true {
				for _, asset := range target.(map[string]interface{})["targets"].([]interface{}) {
					target_name := asset.(map[string]interface{})["name"]
					category := asset.(map[string]interface{})["category"]
					scope := true
					crawlTargets(conn, target_name.(string), category.(string), scope, PK, logo)
				}
			}
		}
	}
	fmt.Println("[+] Bugcrowd Programs Updated. ✅")
}
