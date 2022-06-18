package consts

import (
	"fmt"
	"strings"
)

type Item struct {
	ID   int
	Name string
}

func NewItem(id int, name string) *Item {
	return &Item{ID: id, Name: name}
}

func CreateLookup(itemSets ...[]*Item) map[int]string {
	m := make(map[int]string)
	for _, items := range itemSets {
		for _, i := range items {
			m[i.ID] = i.Name
		}
	}
	return m
}

func Search(s string, itemSets ...[]*Item) (result []string) {
	s = strings.ToLower(s)
	for _, items := range itemSets {
		for _, i := range items {
			if strings.Index(strings.ToLower(i.Name), s) != -1 {
				result = append(result, fmt.Sprintf("%d - %s", i.ID, i.Name))
			}
		}
	}
	return
}
