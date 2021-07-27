package erratum

func Use(o ResourceOpener, s string) error {

	var r Resource
	var e error
	var t bool = true
	for ; t; r, t, e = open(o) {
	}
	if e != nil {
		return e
	}

	return frobbing(r, s)
}

func open(o ResourceOpener) (Resource, bool, error) {
	r, err := o()
	if err != nil {
		if _, ok := err.(TransientError); ok {
			return nil, true, err
		}
		return nil, false, err
	}
	return r, false, nil
}

func frobbing(r Resource, s string) (err error) {
	defer func() {
		if c := recover(); c != nil {
			if ferr, ok := c.(FrobError); ok {
				r.Defrob(ferr.defrobTag)
			}
			if res, ok := c.(error); ok {
				err = res
			}
		}
		r.Close()
	}()

	r.Frob(s)

	return nil
}
