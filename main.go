package main

import "fmt"

// ANSI escape code (colors)
const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Cyan   = "\033[36m"
	Gray   = "\033[37m"
	Bold   = "\033[1m"
)

func main() {
	for {
		// Accept input from a user
		var value int
		fmt.Print("Enter number (decimal): ")
		fmt.Scanf("%d", &value)
		// Operation => (Clear + shift right)

		// Get the first 5 bits
		teamOne := value & 0b0000_0000_0001_1111
		if teamOne > 17 {
			fmt.Printf("%s \n", "the team one exceeded the 5 bits range")
			continue
		}

		// Get the second 5 bits
		teamTwo := (value & 0b0000_0011_1110_0000) >> 5
		if teamTwo > 17 {
			fmt.Printf("%s \n", "the team two exceeded the 5 bits range")
			continue
		}

		// Get the two match state bits
		matchState := (value & 0b0000_1100_0000_0000) >> 10

		// Get the one bit of the team stadium
		teamStadium := (value & 0b0001_0000_0000_0000) >> 12

		// Get the two championship bits
		championship := (value & 0b0110_0000_0000_0000) >> 13

		// Get the one referee bit
		referee := (value & 0b1000_0000_0000_0000) >> 15

		// Then your Printf looks much cleaner:
		fmt.Printf("[%s vs %s] • [%s] • [%s] • [%s] • [Ref: %s]\n",
			colorize(bringTeam(teamOne), Cyan),
			colorize(bringTeam(teamTwo), Cyan),
			colorize(bringMatchState(matchState), Red),
			colorize(bringStadium(teamStadium), Gray),
			colorize(bringChamp(championship), Yellow),
			colorize(bringReferee(referee), Gray),
		)
	}
}

func colorize(text string, colorCode string) string {
	return colorCode + text + "\033[0m"
}

func bringReferee(referee int) string {
	switch referee {
	case 0:
		return "local referee"
	case 1:
		return "international referee"
	default:
		return "No such referee"
	}
}

func bringChamp(champ int) string {
	switch champ {
	case 0:
		return "Saudi pro league"
	case 1:
		return "Saudi King's cup"
	case 2:
		return "Saudi crown prince cup"
	case 3:
		return "Saudi super cup"
	default:
		return "No such championship"
	}
}

func bringStadium(stadium int) string {
	switch stadium {
	case 0:
		return "team one stadium"
	case 1:
		return "team two stadium"
	default:
		return "No such stadium"
	}
}

func bringMatchState(state int) string {
	switch state {
	case 0:
		return "team one won"
	case 1:
		return "team two won"
	case 2:
		return "tie between team one and two"
	case 3:
		return "match canceled"
	default:
		return "No such status"
	}
}

func bringTeam(team int) string {
	switch team {
	case 0:
		return "Al-Nassr"
	case 1:
		return "Al-Hilal"
	case 2:
		return "Al-Ahli"
	case 3:
		return "Al-Qadisiyah"
	case 4:
		return "Al Taawon"
	case 5:
		return "Al-Ittihad"
	case 6:
		return "Al-Ettifaq"
	case 7:
		return "NEOM"
	case 8:
		return "Al-Hazm"
	case 9:
		return "Al Khaleej"
	case 10:
		return "Al-Fayha"
	case 11:
		return "Al Shabab"
	case 12:
		return "Al-Fateh"
	case 13:
		return "Al Kholood"
	case 14:
		return "Damac"
	case 15:
		return "Al Riyadh"
	case 16:
		return "Al Okhdood"
	case 17:
		return "Al Najma"
	default:
		return "No such team"
	}
}
