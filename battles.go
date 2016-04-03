package main

import (
	"sort"
	"math/rand"
)

func runBattle(players Players) {
	sort.Sort(players)

	DisplayInfo(players[0])
	DisplayInfo(players[1])
	DisplayInfo(players)
	round := 1
	numAlive := players.Len()
	playerAction := 0
	for {
		for x := 0; x < players.Len(); x++ {
			players[x].Evasion = 0
		}
		DisplayInfo("Combat round", round, "begins...")
		for x := 0; x < players.Len(); x++ {
			if players[x].Alive != true {
				continue
			}
			playerAction = 0
			if !players[x].Npc {
				DisplayInfo("DO you want to")
				DisplayInfo("\t1 - Run")
				DisplayInfo("\t2 - Evade")
				DisplayInfo("\t3 - Attack")
				GetUserInput(&playerAction)
			}
			if playerAction == 2 {
				players[x].Evasion = rand.Intn(15)
				DisplayInfo("Evasion set to:", players[x].Evasion)
			}
			tgt := selectTarget(players, x)
			if tgt != -1 {
				DisplayInfo("player: ", x, "target: ", tgt)
				attp1 := players[x].Attack() - players[tgt].Evasion
				if attp1 < 0 {
					attp1 = 0
				}
				players[tgt].Health = players[tgt].Health - attp1
				if players[tgt].Health <= 0 {
					players[tgt].Alive = false
					numAlive--
				}
				DisplayInfo(players[x].Name+" attacks and does", attp1, "points of damage with his", Weaps[players[x].Weap].Name, "to the ennemy.")
			}
		}
		if endBattle(players) || playerAction == 1 {
			break
		} else {
			DisplayInfo(players)
			round++
		}
	}
	DisplayInfo(players)
	DisplayInfo("Combat is over...")
	for x := 0; x < players.Len(); x++ {
		if players[x].Alive == true {
			DisplayInfo(players[x].Name + " is still alive!!!")
		}
	}
	DisplayInfo(players)
}

func selectTarget(players []Character, x int) int {
	y := x
	for {
		y = y + 1
		if y >= len(players) {
			y = 0
		}
		if (players[y].Npc != players[x].Npc) && players[y].Alive {
			return y
		}
		if y == x {
			return -1
		}
	}
	return -1
}

func endBattle(players []Character) bool {
	count := make([]int, 2)
	count[0] = 0
	count[1] = 0
	for _, pla := range players {
		if pla.Alive {
			if pla.Npc == false {
				count[0]++
			} else {
				count[1]++
			}
		}
	}
	if count[0] == 0 || count[1] == 0 {
		return true
	} else {
		return false
	}
}
