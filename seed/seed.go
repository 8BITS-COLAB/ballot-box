package seed

import (
	"github.com/8BITS-COLAB/ballot-box/candidate"
	"github.com/jaswdr/faker"
)

var f = faker.New()

func Up() {
	var cs []candidate.Candidate

	for i := 0; i < 10; i++ {
		cs = append(cs, candidate.Candidate{
			Name:  f.Person().FirstName(),
			Party: f.Lorem().Word(),
		})
	}

	for _, c := range cs {
		candidate.New(c.Name, c.Party)
	}
}
