// main.go
package main

func main() {
	config := loadConfig()
	conn := database(config.Database)
	initTables(conn)
	hackerone(conn)
	bugcrowd(conn)
	conn.Close()
}
