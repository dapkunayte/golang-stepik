package main

import "fmt"

type Warrior struct {
	On    bool
	Ammo  int
	Power int
}

func (w *Warrior) Shoot() bool {
	var result bool
	if w.On == false || w.Ammo == 0 {
		result = false
	} else if w.Ammo > 0 {
		w.Ammo--
		result = true
	}
	return result
}

func (w *Warrior) RideBike() bool {
	var result bool
	if w.On == false || w.Power == 0 {
		result = false
	} else if w.Power > 0 {
		w.Power--
		result = true
	}
	return result
}

func main() {
	warrior := &Warrior{false, 1, 1}
	fmt.Println(warrior.Shoot())
	fmt.Println(warrior.RideBike())
}
