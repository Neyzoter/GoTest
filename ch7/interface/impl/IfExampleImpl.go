package impl

type Apple struct {
	Name string
}

func (apple Apple) GetName() string {
	return apple.Name
}

func (apple *Apple) SetName(v string) {
	apple.Name = v
}