package LinkData

import (
	"fmt"
	"strings"
)

type LinkData struct {
	Link, Date string
}

func (l LinkData) String() string {
	d := strings.TrimSpace(l.Date)
	return strings.TrimSpace(l.Link) + "|" + strings.Replace(d, "/", "-", -1)
}

// NewLinkData creates a LinkData from a string in the format "link|date"
func NewLinkData(input string) LinkData {
	split := strings.Split(input, "|")
	fmt.Println(split)
	return LinkData{
		Link: split[0],
		Date: split[1],
	}
}