# postaje

Paket vreme se uporablja za beleženje temperaturnih meritev vremenskih postaj.

Primer uporabe:

```go
postaje := make(map[string]vreme.Postaja)
vreme.DodajPostajo(postaje, "LJ08", "Ljubljana FRI", "Ljubljana")
vreme.DodajMeritev(postaje, "LJ08", 12.1, -50, 50)
vreme.IzpisMeritev(postaje)
vreme.IzpisPorocila(postaje, 3)
```

Modul vsebuje tudi ukazno aplikacijo `postaje` (`cmd/postaje`), ki prikaže
delovanje paketa ; požene se jo z:

```bash
go run ./cmd/postaje --minTemp=-30 --maxTemp=45 --stMeritev=2
```

- `--minTemp`: najnižja dovoljena temperatura meritve (privzeto -50 °C)
- `--maxTemp`: najvišja dovoljena temperatura meritve (privzeto 50 °C)
- `--stMeritev`: najmanjše število meritev za izračun statistike (privzeto 3)
