package main

import "migrationSurgery/service"

var s service.SurgeryService

func main() {
	s.Migrate()
}
