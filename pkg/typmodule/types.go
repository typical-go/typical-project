package typmodule

// Provider responsible to provide dependency
type Provider interface {
	Provide() []interface{}
}

// Preparer responsible to prepare
type Preparer interface {
	Prepare() []interface{}
}

// Destroyer responsible to destruct dependency
type Destroyer interface {
	Destroy() []interface{}
}

// Actionable responsible to provide action
type Actionable interface {
	Action() interface{}
}

// Validator responsible to validate the struct
type Validator interface {
	Validate() error
}

// IsProvider return true if object implementation of provider
func IsProvider(obj interface{}) (ok bool) {
	_, ok = obj.(Provider)
	return
}

// IsPreparer return true obj implement Preparer
func IsPreparer(obj interface{}) (ok bool) {
	_, ok = obj.(Preparer)
	return
}

// IsDestroyer return true if object implementation of destructor
func IsDestroyer(obj interface{}) (ok bool) {
	_, ok = obj.(Destroyer)
	return
}

// IsActionable return true if object is actionable
func IsActionable(obj interface{}) bool {
	_, ok := obj.(Actionable)
	return ok
}

// IsValidator return true if object is actionable
func IsValidator(obj interface{}) bool {
	_, ok := obj.(Validator)
	return ok
}
