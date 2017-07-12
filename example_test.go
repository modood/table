// Copyright 2017 modood. All rights reserved.
// license that can be found in the LICENSE file.

package table_test

import (
	"fmt"

	"github.com/modood/table"
)

type House struct {
	Name  string
	Sigil string
	Motto string
}

func Example() {
	s := []House{
		{"Stark", "direwolf", "Winter is coming"},
		{"Targaryen", "dragon", "Fire and Blood"},
		{"Lannister", "lion", "Hear Me Roar"},
	}

	t := table.Table(s)

	fmt.Println(t)
}
