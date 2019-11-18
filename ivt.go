package main

import (
	"fmt"
	"lang.yottadb.com/go/yottadb"
	"os"
)

func main() {
	var buffertary1 yottadb.BufferTArray
	var errstr yottadb.BufferT
	tptoken := yottadb.NOTTP

	err := buffertary1.TpST(tptoken, &errstr, func(tptoken uint64, errstr *yottadb.BufferT) int32 {

		err := yottadb.SetValE(tptoken, nil, "Go World", "^test", []string{"world"})
		if err != nil {
			panic(err)
		}

		// Retrieve the value that was set
		r, err := yottadb.ValE(tptoken, nil, "^test", []string{"world"})
		if err != nil {
			panic(err)
		}
		if r != "Go World" {
			panic("Value not what was expected; did someone else set something?")
		}
		fmt.Println("Insert OK")
		// Set a few more nodes so we can iterate through them
		err = yottadb.SetValE(tptoken, nil, "Go Middle Earth", "^test", []string{"world"})
		if err != nil {
			panic(err)
		}
		// Retrieve the value that was set
		r, err = yottadb.ValE(tptoken, nil, "^test", []string{"world"})
		if err != nil {
			panic(err)
		}
		if r != "Go Middle Earth" {
			panic("Value not what was expected; did someone else set something?")
		}
		fmt.Println("Update OK")

		var cur_sub = ""
		for true {
			cur_sub, err = yottadb.SubNextE(tptoken, nil, "^test", []string{cur_sub})
			if err != nil {
				error_code := yottadb.ErrorCode(err)
				if error_code == yottadb.YDB_ERR_NODEEND {
					break
				} else {
					panic(err)
				}
			}
			yottadb.DeleteE(tptoken, nil, 1, "^test", []string{cur_sub})
		}

		f, err := yottadb.DataE(tptoken, nil, "^test", []string{"world"})
		if err != nil {
			panic(err)
		}
		if f != 0 {
			panic("Data was found ! Delete incomplete")
		}
		fmt.Println("Delete OK")

		return yottadb.YDB_OK
	}, "TEST")
	if err != nil {
		panic(err)
	}
	os.Exit(0)
}
