package behavioral

type Observer interface {
	Update(data interface{})
}

type Subject struct {
	observers []Observer
}

func (s *Subject) Attach(observer Observer) {
	s.observers = append(s.observers, observer)
}

func (s *Subject) Notify(data interface{}) {
	for _, observer := range s.observers {
		observer.Update(data)
	}
}