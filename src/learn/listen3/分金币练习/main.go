package main

var (
	count = 50
	users = []string{
		"Matthew",
		"Sarah",
		"Augustus",
		"Heidi",
		"Emilie",
		"Peter",
		"Giana",
		"Adriano",
		"Aaron",
		"Elizabeth",
	}
	distribution := make(map[string]int,len(users))
)

func dispatchCoin() {
	for _, name := range users {
		for _, c := range name {
			switch c {
			case

			}
		}
	}
}

func main() {
	dispatchCoin()
}
