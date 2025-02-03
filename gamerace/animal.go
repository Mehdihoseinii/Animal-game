package gamerace

//**** Dag ****//------------------------
type Dog struct {
	name  string
	speed int
	x     int
}

func NewDog(name string) *Dog {
	return &Dog{name, 3, 0}
}
func (c *Dog) Name() string {
	return c.name
}

func (c *Dog) Position() int {
	return c.x
}
func (c *Dog) Move() {
	c.x += c.speed
}

//**** Cat ****//*-----------------------
type Cat struct {
	name  string
	speed int
	x     int
}

func NewCat(name string) *Cat {
	return &Cat{name, 1, 0}
}
func (c *Cat) Name() string {
	return c.name
}

func (c *Cat) Position() int {
	return c.x
}
func (c *Cat) Move() {
	c.x += c.speed
}

//**** Owl ****//----------------------------

type Owl struct {
	name  string
	speed int
	x     int
}

func NewOwl(name string) *Owl {
	return &Owl{name, 2, 0}
}
func (c *Owl) Name() string {
	return c.name
}

func (c *Owl) Position() int {
	return c.x
}
func (c *Owl) Move() {
	c.x += c.speed
}
