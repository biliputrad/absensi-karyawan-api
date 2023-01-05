package main

import (
	"absensi-karyawan-api/config"
	"absensi-karyawan-api/service/activity"
	attendance "absensi-karyawan-api/service/attendance"
	"absensi-karyawan-api/service/division"
	"absensi-karyawan-api/service/role"
	"absensi-karyawan-api/service/user"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

func main() {
	// Load configuration
	c, configErr := config.LoadConfig(".")
	if configErr != nil {
		log.Fatal("[CONFIGURATION] Can not load configuration file.")
	}

	// Database init
	dsn := fmt.Sprintf(
		"host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=%v",
		c.DbHost, c.DbUsername, c.DbPassword, c.DbName, c.DbPort, c.DbTz,
	)

	db, dbConErr := gorm.Open(postgres.Open(dsn), &gorm.Config{}, &gorm.Config{
		Logger: logger.Default.LogMode(config.GetDBLogLevel(c.DbLogLevel)),
	})
	if dbConErr != nil {
		log.Fatal("[DATABASE] Database connection failed.")
	}

	log.Println("[DATABASE] Database connection success.")

	//Auto Migrate Database
	log.Println("Registering table..")
	err := config.Attendance(db)
	if err != nil {
		log.Fatal("[DATABASE] Database attendance cant migrate")
	}

	err = config.Activity(db)
	if err != nil {
		log.Fatal("[DATABASE] Database activity cant migrate")
	}

	err = config.User(db)
	if err != nil {
		log.Fatal("[DATABASE] Database user cant migrate")
	}

	err = config.Role(db)
	if err != nil {
		log.Fatal("[DATABASE] Database role cant migrate")
	}

	err = config.Division(db)
	if err != nil {
		log.Fatal("[DATABASE] Database division cant migrate")
	}

	log.Println("[DATABASE] Database migrate success.")

	//setup router
	router := config.SetupRouter(c)

	////grouping route
	activityGroup := router.Group("/activity")
	attendanceGroup := router.Group("/attendance")
	userGroup := router.Group("/user")
	roleGroup := router.Group("/role")
	divisionGroup := router.Group("/division")

	////implement grouping route
	activity.RouteActivity(db, activityGroup)
	attendance.RouteAttendance(db, attendanceGroup)
	user.RouteUser(db, userGroup)
	role.RouteRole(db, roleGroup)
	division.RouteDivision(db, divisionGroup)

	router.Run()
}
