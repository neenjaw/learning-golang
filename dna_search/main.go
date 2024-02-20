package main

import (
	"fmt"
	"slices"
)

type Nucleotide rune

const (
	A Nucleotide = 'A'
	C Nucleotide = 'C'
	G Nucleotide = 'G'
	T Nucleotide = 'T'
)

func parseNucleotide(r rune) (Nucleotide, error) {
	switch r {
	case 'A':
		return A, nil
	case 'C':
		return C, nil
	case 'G':
		return G, nil
	case 'T':
		return T, nil
	default:
		return 0, fmt.Errorf("invalid nucleotide: %c", r)
	}
}

func (g Nucleotide) compareTo(other Nucleotide) int {
	if g < other {
		return -1
	} else if g > other {
		return 1
	} else {
		return 0
	}
}

type Codon struct {
	Nucleotide1 Nucleotide
	Nucleotide2 Nucleotide
	Nucleotide3 Nucleotide
}

func (c Codon) compareTo(other Codon) int {
	if cmp := c.Nucleotide1.compareTo(other.Nucleotide1); cmp != 0 {
		return cmp
	} else if cmp := c.Nucleotide2.compareTo(other.Nucleotide2); cmp != 0 {
		return cmp
	} else {
		return c.Nucleotide3.compareTo(other.Nucleotide3)
	}
}

func parseCodon(s string) (Codon, error) {
	if len(s) != 3 {
		return Codon{}, fmt.Errorf("invalid codon: %s", s)
	}

	n1, err := parseNucleotide(rune(s[0]))
	if err != nil {
		return Codon{}, fmt.Errorf("invalid first nucleotide: %c", s[0])
	}
	n2, err := parseNucleotide(rune(s[1]))
	if err != nil {
		return Codon{}, fmt.Errorf("invalid second nucleotide: %c", s[1])
	}
	n3, err := parseNucleotide(rune(s[2]))
	if err != nil {
		return Codon{}, fmt.Errorf("invalid third nucleotide: %c", s[2])
	}

	return Codon{n1, n2, n3}, nil
}

func parseCodons(s string) ([]Codon, error) {
	if len(s)%3 != 0 {
		return nil, fmt.Errorf("invalid codons: %s", s)
	}
	codons := make([]Codon, len(s)/3)
	for i := 0; i < len(s); i += 3 {
		codon, err := parseCodon(s[i : i+3])
		if err != nil {
			return nil, err
		}
		codons[i/3] = codon
	}
	return codons, nil
}

func linearContains(codons []Codon, key Codon) bool {
	for _, codon := range codons {
		if codon.compareTo(key) == 0 {
			return true
		}
	}
	return false
}

func binaryContains(codons []Codon, key Codon) bool {
	// assumes codons is sorted
	low, high := 0, len(codons)-1
	for low <= high {
		mid := (low + high) / 2
		if cmp := codons[mid].compareTo(key); cmp < 0 {
			low = mid + 1
		} else if cmp > 0 {
			high = mid - 1
		} else {
			return true
		}
	}
	return false
}

func sortCodons(codons []Codon) []Codon {
	// sort the codons without mutating the original slice
	codonsCopy := make([]Codon, len(codons))
	copy(codonsCopy, codons)
	slices.SortStableFunc(codonsCopy, func(a, b Codon) int {
		return a.compareTo(b)
	})
	return codonsCopy
}

func main() {
	// get codons from gene string from STDIN
	var gene string
	fmt.Scan(&gene)
	codons, err := parseCodons(gene)

	if err != nil {
		fmt.Println(err)
		return
	}

	// print codons
	for _, codon := range codons {
		fmt.Printf("%c%c%c\n", codon.Nucleotide1, codon.Nucleotide2, codon.Nucleotide3)
	}

	codon_a := Codon{A, C, G}
	codon_b := Codon{G, A, T}

	fmt.Println(linearContains(codons, codon_a)) // true
	fmt.Println(linearContains(codons, codon_b)) // false

	sortedCodons := sortCodons(codons)
	fmt.Println(binaryContains(sortedCodons, codon_a)) // true
	fmt.Println(binaryContains(sortedCodons, codon_b)) // false
}
