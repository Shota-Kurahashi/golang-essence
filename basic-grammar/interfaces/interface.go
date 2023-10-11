package interfaces

type speaker interface {
	Speak() error
}

type dog struct{}

func (d *dog) Speak() error {
	println("wan")
	return nil
}

type cat struct{}

func (c *cat) Speak() error {
	println("nya")
	return nil
}

func doSpeak(s speaker) error {
	return s.Speak()
}

func Main() {
	dog := dog{}
	doSpeak(&dog)

	cat := cat{}
	doSpeak(&cat)
}
