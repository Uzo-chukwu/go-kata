package main

import (
	"fmt"
)

func main() {
	questions := []struct {
		question string
		options  []string
		answer   int
	}{
		{"Which greenhouse gas is the primary contributor to global warming?", []string{"Oxygen", "Carbon dioxide", "Nitrogen", "Argon"}, 2},
		{"What is one argument skeptics use against global warming?", []string{"Climate change is a natural cycle.", "Human activities are the sole cause of climate change.", "Global temperatures have remained constant.", "Renewable energy is ineffective."}, 1},
		{"Which extreme weather event is linked to climate change?", []string{"Hurricanes", "Earthquakes", "Tornadoes", "Tsunamis"}, 1},
		{"What is the process where excess carbon dioxide is absorbed by oceans, leading to a decrease in pH levels?", []string{"Ocean acidification", "Lake salinization", "Sea alkalinization", "River neutralization"}, 1},
		{"What is a potential consequence of climate change?", []string{"Increased biodiversity", "More stable weather patterns", "Rising sea levels", "Decreased global temperatures"}, 3},
	}

	score := 0

	fmt.Println("🌍 Global Warming Quiz 🌍")
	for i, q := range questions {
		fmt.Printf("\nQuestion %d: %s\n", i+1, q.question)
		for j, option := range q.options {
			fmt.Printf("%d. %s\n", j+1, option)
		}

		var userAnswer int
		fmt.Print("Your answer (1-4): ")
		fmt.Scanln(&userAnswer)

		if userAnswer == q.answer {
			score++
		}
	}

	fmt.Printf("\nYour score: %d/5\n", score)

	switch score {
	case 5:
		fmt.Println("Excellent!")
	case 4:
		fmt.Println("Very good!")
	default:
		fmt.Println("Time to brush up on your knowledge of global warming.")
		fmt.Println("Check out these resources:")
		fmt.Println("- [Global Warming Quiz](https://www.proprofs.com/quiz-school/topic/global-warming)")
		fmt.Println("- [Climate Change Facts](https://quizax.com/trivia/geography/global-warming-quiz-geography-topics/)")
	}
}
