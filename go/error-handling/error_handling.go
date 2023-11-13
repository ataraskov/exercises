// Package erratum is a solution for exercise Error Handling on Exercism's Go Track.
package erratum

import "fmt"

// Use an opener on provided input
func Use(opener ResourceOpener, input string) (err error) {
	var res Resource

	for ok := true; ok; _, ok = err.(TransientError) {
		res, err = opener()
	}
	if err != nil {
		return err
	}
	// close resource
	defer func() {
		closeErr := res.Close()
		if closeErr != nil {
			err = closeErr
		}
		return
	}()

	// handle possible panic
	defer func() {
		r := recover()
		if r != nil {
			if err, ok := r.(FrobError); ok {
				fmt.Println("calling Defrob")
				res.Defrob(err.defrobTag)
			}
			err = r.(error)
		}
	}()
	res.Frob(input)

	return err
}
