package erratum

const testVersion = 2

func Use(o ResourceOpener, input string) (err error) {
	var resource Resource
	resource, err = o()
	for err != nil {
		switch err.(type) {
		case TransientError:
			resource, err = o()
		default:
			return err
		}
	}
	defer resource.Close()
	defer func() {
		if r := recover(); r != nil {
			switch x := r.(type) {
			case FrobError:
				err = x
				resource.Defrob(x.defrobTag)
			case error:
				err = x
			}
		}
	}()

	resource.Frob(input)
	return err
}
