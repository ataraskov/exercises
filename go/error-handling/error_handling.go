// Package erratum is a solution for exercise Error Handling on Exercism's Go Track.
package erratum

// Use an opener on provided input
func Use(opener ResourceOpener, input string) (err error) {
	var res Resource

	// open resource
	for ok := true; ok; _, ok = err.(TransientError) {
		res, err = opener()
		// there should be some limiter
	}
	if err != nil {
		return err
	}
	defer res.Close()

	// handle possible panic
	defer func() {
		if r := recover(); r != nil {
			if err, ok := r.(FrobError); ok {
				res.Defrob(err.defrobTag)
			}
			err = r.(error)
		}
	}()
	res.Frob(input)

	return err
}
