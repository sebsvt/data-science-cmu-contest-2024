package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(parseDate("2024-01-01"))
}

func parseDate(s string) time.Time {
	t, _ := time.Parse("2006-01-02", s) // Adjust format if needed
	return t
}
