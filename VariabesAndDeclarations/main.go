package main

import (
   "fmt"

   _ "github.com/briancoveney/stringutil"
   "github.com/briancoveney/stringutil"
)

func main() {

   fmt.Println(stringutil.Reverse("!oG, olleH"))

   fmt.Println(stringutil.MyName)

   // var power int = 9000

   // Go has a shortcut variabe devlaration operator, := which ca infer type:

   power := 9000

   fmt.Printf("It's over %d\n", power)

   pow := getPower()
   fmt.Printf("It's over %d\n", pow)
}

// Also works well with functions:
func getPower() int {
   return 9001
}

