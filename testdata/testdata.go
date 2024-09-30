package test

type Foo struct {
	Name       string
	FamilyName string
	Age        int
	Child      Child
	ChildPtr   *Child
}

type Child struct {
	Name string
}

func SomeFunc0() {
	_ = Foo{ // want "field Name is required but not initialized"
		// allrequired
		FamilyName: "Anderson",
		Age:        28,
		Child:      Child{},
		ChildPtr:   &Child{},
	}
}

func SomeFunc1() {
	_ = Foo{ // want "field Child is required but not initialized"
		// allrequired
		Name:       "Thomas",
		FamilyName: "Anderson",
		Age:        28,
		ChildPtr:   &Child{},
	}
}

func SomeFunc2() {
	_ = Foo{
		// allrequired
		Name:       "Thomas",
		FamilyName: "Anderson",
		Age:        28,
		Child:      Child{},
		ChildPtr:   &Child{ // want "field Name is required but not initialized"
			// allrequired
		},
	}
}

func SomeFunc3() {
	_ = Foo{
		// allrequired
		Name:       "Thomas",
		FamilyName: "Anderson",
		Age:        28,
		Child:      Child{},
		ChildPtr: &Child{
			// allrequired
			Name: "Who knows",
		},
	}
}
