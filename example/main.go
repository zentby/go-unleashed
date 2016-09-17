// Copyright 2015 The go-unleashed AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The basicauth command demonstrates using the unleashed.BasicAuthTransport,
// including handling two-factor authentication.  This won't currently work for
// accounts that use SMS to receive one-time passwords.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"go-unleashed/unleashed"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	fmt.Print("Unleashed API Key: ")
	username, _ := r.ReadString('\n')
	username = strings.TrimRight(username, " \n\r")

	fmt.Print("Unleashed API Secret: ")
	password, _ := r.ReadString('\n')
	password = strings.TrimRight(password, "\n\r ")

	client := unleashed.NewClient(username, password)
	products, _, err := client.Products.List(nil)

	if err != nil {
		panic(err)
	}
	fmt.Printf("\n%v\n", unleashed.Stringify(products))
}
