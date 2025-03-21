package main

func main() {
	cards := readFromFile("Test")
	cards.shuffle()

	cards.print()
}
