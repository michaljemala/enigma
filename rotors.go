package enigma

import (
	"strings"
)

// Rotor is the device performing letter substitutions inside
// the Enigma machine. Rotors can be put in different positions,
// swapped, and replaced; they are also rotated during the encoding
// process, following the machine configuration. As a result, there
// are billions of possible combinations, making brute-forcing attacks
// on Enigma unfeasible.
type Rotor struct {
	Sequence string
	Offset   int
	Ring     int
	Notch    map[rune]bool
}

// NewRotor is a constructor taking a mapping string and a notch position
// that will trigger the rotation of the next rotor.
func NewRotor(mapping string, notch []rune) *Rotor {
	notchMap := map[rune]bool{}
	for _, char := range notch {
		notchMap[char] = true
	}
	return &Rotor{mapping, 0, 0, notchMap}
}

// Step through the rotor, performing the letter substitution depending
// on the offset and direction.
func (r *Rotor) Step(letter *rune, invert bool) {
	number := (ToInt(*letter) - r.Ring + r.Offset + 26) % 26
	if invert {
		number = strings.IndexRune(r.Sequence, ToChar(number))
	} else {
		number = ToInt(rune(r.Sequence[number]))
	}
	*letter = ToChar((number + r.Ring - r.Offset + 26) % 26)
}

// RotorConfig sets the initial configuration for a rotor: ID from
// the pre-defined list, a starting position (A-Z), and a ring setting (1-26).
type RotorConfig struct {
	ID    string
	Start rune
	Ring  int
}

// Rotors match the original Enigma configurations, including the
// notches. "Beta" and "Gamma" are additional rotors used in M4
// at the 4th position.
var Rotors = map[string]Rotor{
	"I":     *NewRotor("EKMFLGDQVZNTOWYHXUSPAIBRCJ", []rune{'Q'}),
	"II":    *NewRotor("AJDKSIRUXBLHWTMCQGZNPYFVOE", []rune{'E'}),
	"III":   *NewRotor("BDFHJLCPRTXVZNYEIWGAKMUSQO", []rune{'V'}),
	"IV":    *NewRotor("ESOVPZJAYQUIRHXLNFTGKDCMWB", []rune{'J'}),
	"V":     *NewRotor("VZBRGITYUPSDNHLXAWMJQOFECK", []rune{'Z'}),
	"VI":    *NewRotor("JPGVOUMFYQBENHZRDKASXLICTW", []rune{'Z', 'M'}),
	"VII":   *NewRotor("NZJHGRCXMYSWBOUFAIVLPEKQDT", []rune{'Z', 'M'}),
	"VIII":  *NewRotor("FKQHTLXOCBJSPDZRAMEWNIUYGV", []rune{'Z', 'M'}),
	"Beta":  *NewRotor("LEYJVCNIXWPBQMDRTAKZGFUHOS", []rune{}),
	"Gamma": *NewRotor("FSOKANUERHMBTIYCWLQPZXVGJD", []rune{}),
}

// Reflector is used to reverse a signal inside the Enigma: the current
// goes from the keys through the rotors to the reflector, then it is
// reversed and goes through the rotors again in the opposite direction.
type Reflector struct {
	Sequence string
}

// Reflect is a method for reversing the Enigma signal in a reflector:
// it is just a simple substitution, essentially.
func (r *Reflector) Reflect(letter *rune) {
	*letter = ToChar(strings.IndexRune(r.Sequence, *letter))
}

// Reflectors in the list are pre-loaded with historically accurate data
// from Enigma machines. Use "B-Thin" and "C-Thin" with M4 (4 rotors).
var Reflectors = map[string]Reflector{
	"A":      Reflector{"EJMZALYXVBWFCRQUONTSPIKHGD"},
	"B":      Reflector{"YRUHQSLDPXNGOKMIEBFZCWVJAT"},
	"C":      Reflector{"FVPJIAOYEDRZXWGCTKUQSBNMHL"},
	"B-Thin": Reflector{"ENKQAUYWJICOPBLMDXZVFTHRGS"},
	"C-Thin": Reflector{"RDOBJNTKVEHMLFCWZAXGYIPSUQ"},
}
