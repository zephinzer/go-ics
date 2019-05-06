package ics

type Person struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Dir   string `json:"dir"`
}

func (person Person) String() string {
	p := "CN=" + person.Name
	if len(person.Dir) > 0 {
		p += ";DIR=\"" + person.Dir + "\""
	}
	if len(person.Email) > 0 {
		p += ":mailto:" + person.Email
	}
	return p
}
