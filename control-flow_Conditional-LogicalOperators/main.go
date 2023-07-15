package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//SEQUENCE - stuff runs in order
	fmt.Println("1st to run")
	fmt.Println("2nd")
	x := 34 // 3rd
	y := 41 // 4th
	fmt.Printf("x = %v \ny = %v\n", x, y)

	//CONDITIONAL - IF or Switch statements
	if x < 42 {
		fmt.Println("Less than the meaning of life")
	}
	// ----------------------------------------------------------------------------------------------------------------------------------------------
	/*
	   "If" statements specify the conditional execution of two branches--
	    according to the value of a boolean expression.
	   If the expression evaluates to true,
	   the "if" branch is executed, otherwise, if present, the "else" branch is executed.
	*/

	if x < 42 {
		fmt.Println("Less than the meaning of life")
	} else {
		fmt.Println("Equal to or Greater than the meaning of life")
	}

	if x < 42 {
		fmt.Println("Less than the meaning of life")
	} else if x == 42 { // else if as many times as needed, other statements only 1 of each.
		fmt.Println("Equal to the meaning of life")
	} else {
		fmt.Println("Greater than the meaning of life")
	}

	// ADENDUM: The expression may be preceded by a simple statement,
	// which executes before the expression is evaluated.

	if z := x * rand.Intn(x); z >= y {
		fmt.Printf("z is %v which is GREATER THAN OR EQUAL TO y which is %v", z, y)
	} else {
		fmt.Printf("z is %v which is LESS THAN y which is %v", z, y)
	}

	/*
		COMMA OK:
		In this example, if tz is present, seconds will be set appropriately and ok will be true;
		if not, seconds will be set to zero and ok will be false.
		Here's a function that puts it together with a nice error report:

		func offset(tz string) int {
		    if seconds, ok := timeZone[tz]; ok {
		        return seconds
		    }
		    log.Println("unknown time zone:", tz)
		    return 0
		}*/

	// ----------------------------------------------------------------------------------------------------------------------------------------------

	if x < 42 && y < 42 { // && requires both to be true
		fmt.Println("Both are less than the meaning of life")
	}

	if x < 42 || x < 30 {
		fmt.Println("x is close to the meaning of life")
	}

	if x != 42 && y != 10 {
		fmt.Println("x is not 42 and y is not 10")
	}

	/*
		Logical operators
		apply to boolean values and yield a result of the same type as the operands.
		The right operand is evaluated conditionally.

			    &&    conditional AND    p && q  is  "if p then q else false"
			    ||    conditional OR     p || q  is  "if p then true else q"
			    !     NOT                !p      is  "not p"
	*/
	// ----------------------------------------------------------------------------------------------------------------------------------------------

}
