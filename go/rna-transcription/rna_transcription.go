// Package strand provides a function to convert DNA to RNA
package strand

// ToRNA returns the RNA complement of a DNA strand
func ToRNA(dna string) string {
	translate := map[rune]rune{
		'G': 'C',
		'C': 'G',
		'T': 'A',
		'A': 'U',
	}

	rna := make([]rune, len(dna))
	for i, n := range dna {
		rna[i] = translate[n]
	}

	return string(rna)
}
