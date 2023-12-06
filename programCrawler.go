package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func crawlPrograms(conn *sql.DB, name string, submission string, platform string, bounty string, logo string) int {
	res := conn.QueryRow("SELECT * FROM Programs WHERE name=$1 AND platform=$2;", name, platform)
	dbProgram := new(Program)
	res.Scan(&dbProgram.id, &dbProgram.name, &dbProgram.platform, &dbProgram.submission, &dbProgram.bounty)
	data := make(map[string]string)
	if dbProgram.id == 0 {
		res = nil
		_, err := conn.Query("INSERT INTO Programs(name,platform,submission,bounty)VALUES($1, $2, $3, $4);", name, platform, submission, bounty)
		if err != nil {
			log.Fatal(err)
		}
		res := conn.QueryRow("SELECT * FROM Programs WHERE name=$1 AND platform=$2;", name, platform)
		res.Scan(&dbProgram.id, &dbProgram.name, &dbProgram.platform, &dbProgram.submission, &dbProgram.bounty)
		data["Title"] = "[+] New Program"
		data["Program"] = fmt.Sprintf("[+] %v", name)
		data["Platform"] = fmt.Sprintf("[+] %v", platform)
		data["Bounty"] = fmt.Sprintf("[+] %v", bounty)
		data["Submission"] = fmt.Sprintf("[+] %v", submission)
		data["Logo"] = fmt.Sprintf("[+] %v", logo)

		/*
			if Bounty == True{
				if !inactive { discord(data) }
			}
		*/
	} else {
		if dbProgram.bounty == false && bounty == "false" {
			data["Title"] = "[+] Change Program From VDP To BBP"
			data["Program"] = fmt.Sprintf("[+] %v", name)
			data["Platform"] = fmt.Sprintf("[+] %v", platform)
			data["Bounty"] = fmt.Sprintf("[+] %v", bounty)
			data["Submission"] = fmt.Sprintf("[+] %v", submission)
			data["Logo"] = fmt.Sprintf("[+] %v", logo)
			/*
				if Bounty == True{
					if !inactive { discord(data) }
				}
			*/
		} else if dbProgram.bounty == true && bounty == "false" {
			data["Title"] = "[-] Change Program From BBP To VDP"
			data["Program"] = fmt.Sprintf("[+] %v", name)
			data["Platform"] = fmt.Sprintf("[+] %v", platform)
			data["Bounty"] = fmt.Sprintf("[+] %v", bounty)
			data["Submission"] = fmt.Sprintf("[+] %v", submission)
			data["Logo"] = fmt.Sprintf("[+] %v", logo)
			/*
				if Bounty == True{
					if !inactive { discord(data) }
				}
			*/
		}

		res = nil
		_, err := conn.Query("UPDATE Programs SET bounty = $1 WHERE id = $2;", bounty, dbProgram.id)
		if err != nil {
			log.Fatal(err)
		}
	}
	return dbProgram.id
}
