package seed

import "github.com/8BITS-COLAB/ballot-box/candidate"

func Up() {
	candidates := []candidate.Candidate{
		{
			Name:  "John Doe",
			Party: "Democrat",
		},
		{
			Name:  "Jane Doe",
			Party: "Democrat",
		},
		{
			Name:  "John Smith",
			Party: "Republican",
		},
		{
			Name:  "Jane Smith",
			Party: "Republican",
		},
	}

	for _, c := range candidates {
		candidate.New(c.Name, c.Party)
	}
}
