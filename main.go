package main

import (
	"fmt"
	"sort"
	"strings"
)

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func main() {
	fmt.Println("Go Test Example!")
}
