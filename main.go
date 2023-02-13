package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

type applicant struct {
	company string
	plan    plan
	next    bool
}

type plan string

const (
	planPlaTinum plan = "platinum"
	planGold     plan = "gold"
	planSilver   plan = "silver"
	planBronze   plan = "bronze"
	planFree     plan = "free"
)

func run() error {

	cr := csv.NewReader(os.Stdin)
	records, err := cr.ReadAll()
	if err != nil {
		return err
	}

	records = records[1:] // skip header
	applicants := make(map[plan][]applicant, len(records))
	companies := make(map[string]bool, len(records))
	for _, record := range records {

		// skip duplicated company
		if companies[record[0]] {
			fmt.Fprintln(os.Stderr, record[0], "is duplicated")
			continue
		}
		companies[record[0]] = true

		a := applicant{
			company: record[0],
			plan:    plan(record[1]),
		}
		a.next, err = strconv.ParseBool(record[2])
		if err != nil {
			return err
		}

		applicants[a.plan] = append(applicants[a.plan], a)
	}

	seed := time.Now().UnixNano()
	rand.Seed(seed)
	fmt.Println("Seed is", seed)
	fmt.Println()

	time.Sleep(2 * time.Second)

	// Platinum "Go"ld sponsor
	fmt.Println(`==== Platinum "Go"ld sponsor ====`)
	rand.Shuffle(len(applicants[planPlaTinum]), func(i, j int) {
		applicants[planPlaTinum][i], applicants[planPlaTinum][j] = applicants[planPlaTinum][j], applicants[planPlaTinum][i]
	})
	printCompany(applicants[planPlaTinum][0].company, 1 * time.Second)
	printCompany(applicants[planPlaTinum][1].company, 1 * time.Second)
	for _, a := range applicants[planPlaTinum][2:] {
		if a.next {
			applicants[planGold] = append(applicants[planGold], a)
		}
	}
	fmt.Println()

	// "Go"ld sponsor
	fmt.Println(`==== "Go"ld sponsor ====`)
	rand.Shuffle(len(applicants[planGold]), func(i, j int) {
		applicants[planGold][i], applicants[planGold][j] = applicants[planGold][j], applicants[planGold][i]
	})
	printCompany(applicants[planGold][0].company, 1 * time.Second)
	printCompany(applicants[planGold][1].company, 1 * time.Second)

	for _, a := range applicants[planGold][2:] {
		if a.next {
			applicants[planSilver] = append(applicants[planSilver], a)
		}
	}
	fmt.Println()

	// Silver sponsor
	fmt.Println("==== Silver sponsor ====")
	rand.Shuffle(len(applicants[planSilver]), func(i, j int) {
		applicants[planSilver][i], applicants[planSilver][j] = applicants[planSilver][j], applicants[planSilver][i]
	})
	for _, a := range applicants[planSilver] {
		fmt.Println(a.company)
	}
	fmt.Println()

	// Bronze sponsor
	fmt.Println("==== Bronze sponsor ====")
	rand.Shuffle(len(applicants[planBronze]), func(i, j int) {
		applicants[planBronze][i], applicants[planBronze][j] = applicants[planBronze][j], applicants[planBronze][i]
	})
	for _, a := range applicants[planBronze] {
		fmt.Println(a.company)
	}
	fmt.Println()

	// Free sponsor
	fmt.Println("==== Free sponsor ====")
	rand.Shuffle(len(applicants[planFree]), func(i, j int) {
		applicants[planFree][i], applicants[planFree][j] = applicants[planFree][j], applicants[planFree][i]
	})
	for _, a := range applicants[planFree] {
		fmt.Println(a.company)
	}

	return nil
}

func printCompany(name string, d time.Duration) {
	var n int
	for _, c := range name {
		fmt.Printf("%c", c)
		n++
		if n < 6 {
			time.Sleep(d)
		}
	}
	fmt.Println()
	time.Sleep(d)
}
