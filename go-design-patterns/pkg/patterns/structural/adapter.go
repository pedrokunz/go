package structural

type Target interface {
	Request() string
}

type Adaptee struct{}

func (a *Adaptee) SpecificRequest() string {
	return "Specific request from Adaptee"
}

type Adapter struct {
	adaptee *Adaptee
}

func NewAdapter(adaptee *Adaptee) *Adapter {
	return &Adapter{adaptee: adaptee}
}

func (a *Adapter) Request() string {
	return a.adaptee.SpecificRequest()
}

func (a *Adapter) Adapt() string {
	return "Adapted: " + a.Request()
}