# tristarget

Small tool I created for Go learning purposes.
Allows you to fetch departure data for your stop of choice.
Works only in Gdansk. Using data from [Otwarte dane ZTM w Gdańsku](https://ztm.gda.pl/ztm/otwarte-dane,a,3099)

## Installation
```bash
go install https://github.com/PawelReich/tristarget.git@latest
```

## How to find my stop?
```bash
❯ tristarget find -name "brama oliwska"
1541: Brama Oliwska 04
1542: Brama Oliwska 03
2005: Brama Oliwska 01
2006: Brama Oliwska 02
```

## How to get departure data?
```bash
❯ tristarget departures -stop 2005 -limit 1 -route 3
3 21:42:53
```

