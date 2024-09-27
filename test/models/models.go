package models

type Test struct {
	Name       string
	Age        int
	FamilyName string
	Child      Child
	Child2     *Child
}

type Child struct {
	Name string
}
