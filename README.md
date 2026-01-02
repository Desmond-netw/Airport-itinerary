# Airport Itinerary-prettifier App

## Intro.
A command line tool, which reads a text-based itinerary from a file (input), processes the text to make it customer-friendly, and writes the result to a new file (output).

The Itinerary - prettifier cli app automatically ;
- Converts **airport codes (IATA & ICAO )** to full airport or city names.
- Formats **ISO date &  Zulu time** into readable formats.
- Cleans **extral whitespace and line breaks** .
- ouputs **color-highlighted results** (optional feature)

## Features

Converting ariport codes : 
- `#HKN` => kimbe Airport
- `##AYMD` => Madang Airport
- `*#LHR` => London (city name)

Formats ISO date/time
- `D(2025-10-13T09:00+02:00)` → **13 Oct 2025**
- `T12(2025-10-13T08:45+01:00)` → **08:45AM (+01:00)**
- `T24(2025-10-13T08:45+01:00)` → **08:45 (+01:00)**

Cleans text whitespace 
- It changes `\r`, `\v`, `\f` to `\n`
- Cleaning excessive white and blank spaces after each col

##  Usage

### CLI Arugments  to passed
go run . ./PATH/INPUT.TXT  ./PATH/OUTPUT.TXT .AIRPORT-LOOKUP.CSV

```bash
 go run . ./input.txt  ./output.txt ./airport-lookup.csv 
```
- '#Optional feature'
        If you want to print the colorful output to ther terminal 
        use ther `print flag (-p)` to perfom that operation
```bash
go run . -p ./input-DT.txt ./output-DT.txt ./airport-lookup.csv
```
* NOTE:
        |./input.txt = inputfile path|
        |./output.txt = ouputfile path |
        |./airport-lookup = airportlookup file path |
### Example 

**Run**

- Download the application to your local machine from github
- ``cd`` [ application directory]
- run  ``go run . ./input.txt ./ouput.txt ./csv/airport-lookup.csv`` 

| Input test | Output |
|-----------|--------|
| Trip Details<br>From: #HKN<br>To: ##AYMD<br>Departure: D(2025-10-13T12:30Z)<br>Arrival: T12(2025-10-13T20:45+01:00)<br><br>Your flight from #LAX to the destination Airport ##EGLL. Departure time is D(2025-10-13T12:30Z) and the Arrival time is T12(2025-10-13T20:45+01:00)<br>Enjoy your flight. | Trip Details<br>From: Kimbe Airport<br>To: Madang Airport<br>Departure: 13 Oct 2025<br>Arrival: 08:45PM (+01:00)<br><br>Your flight from Los Angeles International Airport to the destination Airport London Heathrow Airport. Departure time is 13 Oct 2025 and the Arrival time is 08:45PM (+01:00)<br>Enjoy your flight. |

## Error Handling
| Situation | Message |
|------------|----------|
| Missing arguments | `itinerary usage: go run . ./input.txt ./output.txt ./airport-lookup.csv` |
| Input file missing | `Input not found.` |
| CSV file missing | `Airport lookup not found.` |
| CSV malformed | `Airport lookup malformed` |
| Other unknown error | `Unable to write to output file.` |

## Project Compeleting and Learning Outcomes

This project teaches:
- File I/O (reading and writing files)
- CSV parsing and validation
- Regular expressions in Go
- String manipulation and cleanup
- Time formatting with time zones
- Command-line argument handling with `flag`
- ANSI terminal color formatting
- Clean Go project structure