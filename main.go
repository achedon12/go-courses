package main

import (
	"fmt"
	"time"
)

func main() {
	s := "Léo"
	now := time.Now()
	leoBirthday := time.Date(2003, time.May, 10, 00, 15, 0, 0, time.UTC)

	// Calcul des années
	leoAgeYears := now.Year() - leoBirthday.Year()
	if now.Month() < leoBirthday.Month() || (now.Month() == leoBirthday.Month() && now.Day() < leoBirthday.Day()) {
		leoAgeYears--
	}

	// Calcul des mois
	temp := leoBirthday.AddDate(leoAgeYears, 0, 0)
	leoAgeMonths := int(now.Month() - temp.Month())
	if now.Day() < temp.Day() {
		leoAgeMonths--
	}

	// Calcul des jours
	temp = leoBirthday.AddDate(leoAgeYears, leoAgeMonths, 0)
	leoAgeDays := now.Day() - temp.Day()

	fmt.Printf("Hello and welcome, %s!\n", s)
	fmt.Printf("Today is %s\n", now.Format("Monday, January 2, 2006"))
	fmt.Printf("Leo's Birthday is %s\n", leoBirthday.Format("Monday, January 2, 2006"))
	fmt.Printf("Leo's age is %d years old, %d month(s) and %d day(s)\n", leoAgeYears, leoAgeMonths, leoAgeDays)
}
