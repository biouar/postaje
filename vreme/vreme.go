// Paket vreme omogoča beleženje temperaturnih meritev vremenskih postaj in
// izračun osnovne statistike meritev.
//
// Postaje hranimo v slovarju, katerega ključ je oznaka postaje, vrednost pa struktura [Postaja]. 
// Ker so polja strukture neizvožena, postajo ustvarimo s funkcijo [DodajPostajo].
//
// # Primer uporabe
//
//	postaje := make(map[string]vreme.Postaja)
//	vreme.DodajPostajo(postaje, "LJ08", "Ljubljana FRI", "Ljubljana")
//	vreme.DodajMeritev(postaje, "LJ08", 12.1, -50, 50)
//	vreme.IzpisMeritev(postaje)
//	vreme.IzpisPorocila(postaje, 3)
package vreme

import (
	"fmt"
)

// Postaja predstavlja vremensko postajo z imenom, krajem in meritvami temperature.
type Postaja struct {
	ime     string
	kraj    string
	meritve []float64
}

// DodajPostajo v slovar postaje doda novo postajo z dano oznako, imenom in krajem.
// Če postaja s to oznako že obstaja, funkcija izpiše napako in slovarja ne spremeni.
func DodajPostajo(postaje map[string]Postaja, oznaka, ime, kraj string) {
	if _, obstaja := postaje[oznaka]; obstaja {
		fmt.Printf("Napaka: postaja z oznako %q že obstaja.\n", oznaka)
		return
	}
	postaje[oznaka] = Postaja{ime: ime, kraj: kraj}
}

// DodajMeritev postaji doda novo meritev. Meritev je veljavna, če velja minTemp <= temperatura <= maxTemp. 
// Če postaja ne obstaja ali meritev ni veljavna, funkcija izpiše napako in postaje ne spremeni.
func DodajMeritev(postaje map[string]Postaja, oznaka string, temperatura, minTemp, maxTemp float64) {
	postaja, obstaja := postaje[oznaka]
	if !obstaja {
		fmt.Printf("Napaka: postaja z oznako %q ne obstaja.\n", oznaka)
		return
	}
	if temperatura < minTemp || temperatura > maxTemp {
		fmt.Printf("Napaka: meritev ni veljavna.\n")
		return
	}

	postaja.meritve = append(postaja.meritve, temperatura)
	postaje[oznaka] = postaja // vrednosti v slovarju so kopije
}

// statistika vrne najnižjo, najvišjo in povprečno temperaturo postaje. 
// Če postaja ne obstaja ali ima manj kot stMeritev meritev, vrne ok == false.
func statistika(postaje map[string]Postaja, oznaka string, stMeritev int) (najnizja, najvisja, povprecna float64, ok bool) {
	postaja, obstaja := postaje[oznaka]
	if !obstaja || len(postaja.meritve) == 0 || len(postaja.meritve) < stMeritev {
		return 0, 0, 0, false
	}

	najnizja = postaja.meritve[0]
	najvisja = postaja.meritve[0]
	vsota := 0.0
	for _, m := range postaja.meritve {
		if m < najnizja {
			najnizja = m
		}
		if m > najvisja {
			najvisja = m
		}
		vsota += m
	}
	povprecna = vsota / float64(len(postaja.meritve))
	return najnizja, najvisja, povprecna, true
}

// IzpisMeritev izpiše oznake, imena, kraje in meritve vseh postaj.
// Primer izpisa:
//
//	MERITVE:
//	LJ08 - Ljubljana FRI (Ljubljana): [12.1 10.5 9.8 10.5]
func IzpisMeritev(postaje map[string]Postaja) {
	fmt.Println("MERITVE:")
	for oznaka, p := range postaje {
		fmt.Printf("%s - %s (%s): %v\n", oznaka, p.ime, p.kraj, p.meritve)
	}
}

// IzpisPorocila za vsako postajo izpiše ime, najnižjo, najvišjo in povprečno
// temperaturo ter opis glede na povprečje:
//
//   - Vroče, če je povprečje vsaj 25 °C,
//   - Toplo, če je povprečje vsaj 15 °C in manj kot 25 °C,
//   - Hladno, če je povprečje vsaj 5 °C in manj kot 15 °C,
//   - Mrzlo, če je povprečje manj kot 5 °C.
//
// Če ima postaja manj kot stMeritev meritev, funkcija izpiše "premalo meritev".
func IzpisPorocila(postaje map[string]Postaja, stMeritev int) {
	for oznaka, p := range postaje {
		najnizja, najvisja, povprecna, ok := statistika(postaje, oznaka, stMeritev)
		if !ok {
			fmt.Printf("%s: premalo meritev\n", p.ime)
			continue
		}

		var opis string
		switch {
		case povprecna >= 25:
			opis = "Vroče"
		case povprecna >= 15:
			opis = "Toplo"
		case povprecna >= 5:
			opis = "Hladno"
		default:
			opis = "Mrzlo"
		}

		fmt.Printf("%s: min %.1f °C, max %.1f °C, povprečje %.1f °C -> %s\n",
			p.ime, najnizja, najvisja, povprecna, opis)
	}
}
