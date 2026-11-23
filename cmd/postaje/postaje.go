// Postaje je aplikacija, ki prikaže delovanje paketa [vreme]: ustvari
// nekaj vremenskih postaj, doda meritve ter izpiše meritve in poročilo.
//
// Delovanje nastavimo s tremi stikali:
//	--minTemp    najnižja dovoljena temperatura meritve v °C (privzeto -50)
//	--maxTemp    najvišja dovoljena temperatura meritve v °C (privzeto 50)
//	--stMeritev  najmanjše število meritev za izračun statistike (privzeto 3)
//
// Primer zagona:
//	postaje --minTemp=-30 --maxTemp=45 --stMeritev=2
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/biouar/postaje/vreme"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "postaje",
		Usage: "Beleženje temperaturnih meritev vremenskih postaj",
		Flags: []cli.Flag{
			&cli.FloatFlag{
				Name:  "minTemp",
				Usage: "najnižja dovoljena temperatura meritve v °C",
				Value: -50,
			},
			&cli.FloatFlag{
				Name:  "maxTemp",
				Usage: "najvišja dovoljena temperatura meritve v °C",
				Value: 50,
			},
			&cli.IntFlag{
				Name:  "stMeritev",
				Usage: "najmanjše število meritev, potrebnih za izračun statistike",
				Value: 3,
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			zagon(cmd.Float("minTemp"), cmd.Float("maxTemp"), int(cmd.Int("stMeritev")))
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

func zagon(minTemp, maxTemp float64, stMeritev int) {
	postaje := make(map[string]vreme.Postaja)

	podatki := []struct {
		oznaka, ime, kraj string
		meritve           []float64
	}{
		{"TO01", "Tolmin Center", "Tolmin", []float64{19.5, 17.1, 18.1, 19.4}},
		{"KP02", "Koper Market", "Koper", []float64{23, 24.5, 26.8}},
		{"LJ08", "Ljubljana FRI", "Ljubljana", []float64{12.1, 10.5, 9.8, 10.5}},
		{"TR01", "Triglav", "Triglav", []float64{-2.4}},
	}
	for _, p := range podatki {
		vreme.DodajPostajo(postaje, p.oznaka, p.ime, p.kraj)
		for _, t := range p.meritve {
			vreme.DodajMeritev(postaje, p.oznaka, t, minTemp, maxTemp)
		}
	}

	vreme.IzpisMeritev(postaje)
	fmt.Println()

	vreme.IzpisPorocila(postaje, stMeritev)
	fmt.Println()

	vreme.DodajMeritev(postaje, "KP02", 25.3, minTemp, maxTemp)  // veljavna
	vreme.DodajMeritev(postaje, "LJ08", 75.0, minTemp, maxTemp)  // prevelika
	vreme.DodajMeritev(postaje, "TR01", -60.2, minTemp, maxTemp) // premajhna
	vreme.DodajMeritev(postaje, "IO87", 10.0, minTemp, maxTemp)  // postaja ne obstaja

	vreme.DodajMeritev(postaje, "TR01", 1.8, minTemp, maxTemp)
	vreme.DodajMeritev(postaje, "TR01", 3.9, minTemp, maxTemp)
	fmt.Println()

	vreme.IzpisMeritev(postaje) // izpis po dodajanju
	fmt.Println()

	vreme.IzpisPorocila(postaje, stMeritev)
}
