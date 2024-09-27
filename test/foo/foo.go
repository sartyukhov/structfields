package foo

import (
	"fmt"

	"github.com/sartyukhov/structfieldschecker/test/models"
)

// nolint
func foo(child models.Child) {
	t := models.Test{
		// allrequired
		FamilyName: "Anderson",
		Child:      child,
		Child2:     &models.Child{
			// allrequired
		},
	}
	fmt.Println(t)
}
