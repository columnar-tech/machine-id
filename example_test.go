package machineid_test

import (
	"fmt"

	machineid "github.com/columnar-tech/machine-id"
)

func Example() {
	id, err := machineid.ID()
	if err != nil {
		panic(err)
	}
	fmt.Println(id)
}

func Example_protectedID() {
	id, err := machineid.ProtectedID()
	if err != nil {
		panic(err)
	}

	fmt.Println(id)
}
