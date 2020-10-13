package main

import "MigrationSurgery/service"

var s service.SurgeryService

func main() {
	s.Migrate()
}
