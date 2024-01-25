package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Player struct {
	id            int
	dice          []int
	score         int
	totalDice     int
	remainingDice int
}

func NewPlayer(id, totalDice int) *Player {
	return &Player{
		id:        id,
		dice:      make([]int, 0),
		score:     0,
		totalDice: totalDice,
	}
}

func (p *Player) throwDice(numDice int) {
	p.dice = make([]int, numDice)
	for i := 0; i < numDice; i++ {
		p.dice[i] = rand.Intn(6) + 1
	}
}

func (p *Player) evaluateDice() {
	i := 0
	for i < len(p.dice) {
		if p.dice[i] == 6 {
			p.score++
			p.dice = append(p.dice[:i], p.dice[i+1:]...)
		} else if p.dice[i] == 1 {
			if i > 0 {
				p.dice[i], p.dice[i-1] = p.dice[i-1], p.dice[i]
				i--
			}
		}
		i++
	}
}

func (p *Player) winner(players []*Player) *Player {
	winner := players[0]
	for _, currentPlayer := range players {
		if currentPlayer.score > winner.score {
			winner = currentPlayer
		}
	}
	return winner
}

func (p *Player) hasDiceRemaining() bool {
	return len(p.dice) > 0
}

func playGame(players []*Player) string {
	round := 1
	output := ""

	for len(players) > 1 {
		output += fmt.Sprintf("====================\nRounde %d lempar dadu:\n", round)

		for _, player := range players {
			if len(player.dice) == 0 {
				player.throwDice(player.totalDice)
			}
			player.throwDice(len(player.dice))
			output += fmt.Sprintf("          Pemain #%d (%d):%v\n", player.id, player.score, player.dice)
		}

		for _, player := range players {
			player.evaluateDice()
			output += "Setelah evaluasi: \n   "
			if player.hasDiceRemaining() {
				output += fmt.Sprintf(" Pemain #%d (%d):%v\n", player.id, player.score, player.dice)
			} else {
				output += fmt.Sprintf(" Pemain #%d (%d):_,(Berhenti bermain karena tidak memiliki dadu)\n", player.id, player.score)
				players = removePlayer(players, player)
			}
		}

		round++
		output += "======================\n"
	}

	output += fmt.Sprintf("Game berakhir karena hanya pemain #%d yang memiliki dadu.\n", players[0].id)

	winnerPlayer := players[0].winner(players)
	output += fmt.Sprintf("Pemain #%d adalah pemenang dengan skor %d.\n", winnerPlayer.id, winnerPlayer.score)

	return output
}

func removePlayer(players []*Player, player *Player) []*Player {
	index := -1
	for i, p := range players {
		if p == player {
			index = i
			break
		}
	}
	if index != -1 {
		players = append(players[:index], players[index+1:]...)
	}
	return players
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var totalPemain, totalDadu int
	fmt.Print("Masukkan total pemain: ")
	fmt.Scan(&totalPemain)
	fmt.Print("Masukkan total dadu: ")
	fmt.Scan(&totalDadu)

	if totalDadu >= 2 && totalPemain >= 2 {
		players := make([]*Player, totalPemain)
		for i := 0; i < totalPemain; i++ {
			players[i] = NewPlayer(i+1, totalDadu)
		}
		output := playGame(players)
		fmt.Println(output)
	} else {
		fmt.Println("Total pemain dan total dadu harus lebih besar atau sama dengan 2")
	}
}
