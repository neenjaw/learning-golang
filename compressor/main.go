package main

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/bits-and-blooms/bitset"
)

func main() {
	compressed := compressGene("ACGT")
	fmt.Println(compressed)

	// print in binary representation
	fmt.Printf("%b\n", compressed)

	decompressed := decompressGene(compressed)
	fmt.Println(decompressed)
}

func compressGene(gene string) bitset.BitSet {
	var b bitset.BitSet

	gene_len := len(gene)
	for i := 0; i < gene_len; i++ {
		first_bit := uint(i * 2)
		second_bit := uint(i*2 + 1)
		r := rune(gene[i])
		switch unicode.ToUpper(r) {
		case 'A':
			b.Clear(first_bit)
			b.Clear(second_bit)
		case 'C':
			b.Clear(first_bit)
			b.Set(second_bit)
		case 'G':
			b.Set(first_bit)
			b.Clear(second_bit)
		case 'T':
			b.Set(first_bit)
			b.Set(second_bit)
		default:
			panic("Invalid nucleotide")
		}
	}

	return b
}

func decompressGene(b bitset.BitSet) string {
	var gene strings.Builder

	for i := uint(0); i < b.Len(); i += 2 {
		// get the 2 bits
		first_bit := b.Test(uint(i))
		second_bit := b.Test(uint(i + 1))

		// convert the bits to a nucleotide
		switch {
		case !first_bit && !second_bit:
			gene.WriteRune('A')
		case !first_bit && second_bit:
			gene.WriteRune('C')
		case first_bit && !second_bit:
			gene.WriteRune('G')
		case first_bit && second_bit:
			gene.WriteRune('T')
		}
	}

	return gene.String()
}
