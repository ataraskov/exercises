// Package erratum is a solution for exercise Error Handling on Exercism's Go Track.
package erratum

import "fmt"

// Use an opener on provided input
func Use(opener ResourceOpener, input string) error {
	var res Resource
	var err error

	for ok := true; ok; _, ok = err.(TransientError) {
		res, err = opener()
	}
	if err != nil {
		return err
	}
	// close resource
	defer func() {
		closeErr := res.Close()
		fmt.Println("closing res, err:", err, "closeErr:", closeErr)
		if closeErr != nil {
			err = closeErr
		}
	}()

	// handle possible panic
	defer func() {
		r := recover()
		fmt.Println("checking for panic, err:", r)
		if r != nil {
			if err, ok := r.(FrobError); ok {
				fmt.Println("calling Defrob")
				res.Defrob(err.defrobTag)
			}
			err = r.(error)
		}
	}()
	res.Frob(input)

	fmt.Println("end of func, err:", err)
	return err
}
