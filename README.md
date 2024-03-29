# South American Qualifiers CLI

south-american-qualifiers-cli is a cli program to be aware of the qualifiers for the FIFA World Cup 2026 in South America, written in Go.  
Data taken from [CONMEBOL API](https://github.com/fcoagz/conmebol)

## Features

* Show the current positions table
* Show next matches for all the teams 
* Show previous matches for all the teams
* See the next matches for a specified team
* See the previous matches for a specified team

## Installation

Use the `go build` command and move the binary into your PATH.

```bash
go build main.go
```

## Usage

#### Show commands
Use `-h` or `-help` for showing the options

```bash
south-american-qualifiers-cli -h
```

#### Positions Table
Use `-table` for showing the current positions table

```bash
south-american-qualifiers-cli -table
```

![alt text](media/table.png)

#### Specified Team Fixture
Use `-fixture` and the team name to see the fixture (case sensitive)

```bash
south-american-qualifiers-cli -fixture <<Team-Name>>
```

![alt text](media/fixture.png)

#### Specified Team Previous Matches
Use `-previous` and the team name to see the previous matches (case sensitive)

```bash
south-american-qualifiers-cli -previous <<Team-Name>>
```

![alt text](media/previous.png)

#### Teams Previous Matches
`-show-previous` for showing the previous matches

```bash
south-american-qualifiers-cli -show-previous
```

![alt text](media/show-previous.png)

#### Teams Fixture
Use `-show-fixture` for showing the next matches

```bash
south-american-qualifiers-cli -show-fixture
```

![alt text](media/show-fixture.png)

