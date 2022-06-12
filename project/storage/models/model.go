package models

import "strings"

type Link struct {
	Long  string
	Short string
}

func MakeLink(long string) Link {
	return Link{
		Long:  long,
		Short: longToShort(long),
	}
}

func longToShort(long string) string {
	s := strings.Split(long, "/")
	short := s[0] + "//" + s[2] + "/"
	return short
}
