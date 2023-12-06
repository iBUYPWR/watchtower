package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func hackerone(conn *sql.DB) {
	BugcrowdURL := "https://raw.githubusercontent.com/Osb0rn3/bugbounty-targets/main/programs/hackerone.json"
	response, err := http.Get(BugcrowdURL)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var data []map[string]map[string]interface{}
	err = json.Unmarshal(body, &data)
	for _, program := range data {
		handle := program["attributes"]["handle"]
		name := program["attributes"]["name"].(string)
		submission := fmt.Sprintf("https://hackerone.com/%v", handle)
		submissionState := program["attributes"]["submission_state"]
		platform := "hackerone"
		logo := program["attributes"]["profile_picture"].(string)
		targets := program["relationships"]["structured_scopes"].(map[string]interface{})["data"].([]interface{})
		bounty := "false"
		if submissionState == "paused" {
			bounty = "false"
		} else {
			bounty = strconv.FormatBool(program["attributes"]["offers_bounties"].(bool))
		}
		if name == "Agoric" {
			bounty = "true"
		}
		PK := crawlPrograms(conn, name, submission, platform, bounty, logo)
		for _, target := range targets {
			attributes := target.(map[string]interface{})["attributes"].(map[string]interface{})
			title := attributes["asset_identifier"]
			category := attributes["asset_type"]
			scope := false
			if attributes["eligible_for_bounty"].(bool) == true && attributes["eligible_for_submission"].(bool) == true {
				scope = true
			} else {
				scope = false
			}
			crawlTargets(conn, title.(string), category.(string), scope, PK, logo)
		}
	}
	fmt.Println("[+] Hackerone Programs Updated. ✅")
}
