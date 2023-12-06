package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func crawlTargets(conn *sql.DB, name string, category string, scope bool, program_id int, logo string) {
	res := conn.QueryRow("SELECT * FROM Targets WHERE name=$1 AND type=$2 AND program_id=$3;", name, category, program_id)
	dbTarget := new(Target)
	res.Scan(&dbTarget.id, &dbTarget.name, &dbTarget.category, &dbTarget.scope, &dbTarget.program_id)
	if dbTarget.id == 0 {
		res = nil
		_, err := conn.Query("INSERT INTO Targets(name,category,scope,program_id)VALUES($1, $2, $3, $4) returning id;", name, category, scope, program_id)
		if err != nil {
			log.Fatal(err)
		}

		res = nil
		res := conn.QueryRow("SELECT * FROM Programs WHERE id=$1;", program_id)
		programDetails := new(Program)
		res.Scan(&programDetails.id, &programDetails.name, &programDetails.platform, &programDetails.submission, &programDetails.bounty)
		if scope == true {
			data := make(map[string]string)
			data["Title"] = "[+] New Scope."
			data["Target"] = fmt.Sprintf("%v", name)
			data["Type"] = fmt.Sprintf("%v", category)
			data["InScope"] = fmt.Sprintf("%v", scope)
			data["Program"] = fmt.Sprintf("%v", programDetails.name)
			data["Bounty"] = fmt.Sprintf("%v", programDetails.bounty)
			data["Platform"] = fmt.Sprintf("%v", programDetails.platform)
			data["Submission"] = fmt.Sprintf("%v", programDetails.submission)
			data["Logo"] = fmt.Sprintf("%v", logo)
			/*
				if Bounty == True{
					if !inactive { discord(data) }
				}
			*/
		}
	} else {
		res = nil
		res := conn.QueryRow("SELECT * FROM Programs WHERE id=$1 ;", dbTarget.program_id)
		programDetails := new(Program)
		res.Scan(&programDetails.id, &programDetails.name, &programDetails.platform, &programDetails.submission, &programDetails.bounty)
		if dbTarget.scope == true && scope == false {
			data := make(map[string]string)
			data["Title"] = "[-] Target is out of Scope Now"
			data["Target"] = fmt.Sprintf("%v", name)
			data["Type"] = fmt.Sprintf("%v", category)
			data["InScope"] = fmt.Sprintf("%v", scope)
			data["Program"] = fmt.Sprintf("%v", programDetails.name)
			data["Bounty"] = fmt.Sprintf("%v", programDetails.bounty)
			data["Platform"] = fmt.Sprintf("%v", programDetails.platform)
			data["Submission"] = fmt.Sprintf("%v", programDetails.submission)
			data["Logo"] = fmt.Sprintf("%v", logo)
			/*
				if Bounty == True{
					if !inactive { discord(data) }
				}
			*/
		} else if dbTarget.scope == false && scope == true {
			data := make(map[string]string)
			data["Title"] = "[+] Target is in Scope Now"
			data["Target"] = fmt.Sprintf("%v", name)
			data["Type"] = fmt.Sprintf("%v", category)
			data["InScope"] = fmt.Sprintf("%v", scope)
			data["Program"] = fmt.Sprintf("%v", programDetails.name)
			data["Bounty"] = fmt.Sprintf("%v", programDetails.bounty)
			data["Platform"] = fmt.Sprintf("%v", programDetails.platform)
			data["Submission"] = fmt.Sprintf("%v", programDetails.submission)
			data["Logo"] = fmt.Sprintf("%v", logo)
			/*
				if Bounty == True{
					if !inactive { discord(data) }
				}
			*/
		}

		res = nil
		_, err := conn.Query("UPDATE Programs SET bounty = $1 WHERE id = $2;", scope, dbTarget.id)
		if err != nil {
			log.Fatal(err)
		}
	}
}
