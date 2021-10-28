package solid

import (
	"fmt"
	"strings"
)

var entryCount = 0

type Journal struct {
	entries []string
}

func (j *Journal) String() string{
	return strings.Join(j.entries, "\n")
}

func (j *Journal) AddEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d: %s", entryCount, text)
	j.entries = append(j.entries, entry)
	return entryCount
}

func (j *Journal) RemoveEntry(index int){
	// ...
}

func SingleResponsibilityPrinciple(){
	fmt.Println("/* SOLID: Single Responsibility Principle */")
	j := Journal{}
	j.AddEntry("Learned Python")
	j.AddEntry("Learned Javascript")
	j.AddEntry("Learned Java")
	j.AddEntry("Hate Java")
	j.AddEntry("Learned Golang")

	fmt.Println(j.String())
}
// Separation of Concern should detail that
// a file has a singular responsibility and
// not act as a god object