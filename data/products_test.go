package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "Coca-Cola",
		Price: 2.50,
		SKU:   "abc-abc-abc",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
