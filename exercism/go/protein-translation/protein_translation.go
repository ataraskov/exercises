// Package proten contains solution for Protein Translation exercise
// from Exercism's Go Track
package protein

import (
	"errors"
)

var (
	ErrInvalidBase = errors.New("invalid base")
	ErrStop        = errors.New("stop")
)

var codons = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

// FromRNA returns proteins array for a given rna sequence
func FromRNA(rna string) ([]string, error) {
	res := []string{}
	for i := 0; i <= len(rna)-3; i += 3 {
		protein, err := FromCodon(rna[i : i+3])
		switch {
		case err == ErrStop:
			return res, nil
		case err != nil:
			return nil, err
		default:
			res = append(res, protein)
		}
	}

	return res, nil
}

const stop = "STOP"

// FromCodon returns protein for a given codon
func FromCodon(codon string) (string, error) {
	if protein, ok := codons[codon]; ok {
		if protein == stop {
			return "", ErrStop
		}
		return protein, nil
	}
	return "", ErrInvalidBase
}
