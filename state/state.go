// Với mỗi trạng thái khác nhau sẽ có hành vi khác nhau

package main

import "fmt"

type StatePower int

const (
	NOPOWER = iota
	LOWPOWER
	FULLPOWER
)

type Speaker struct {
	StateCtx StatePower
}

func (s *Speaker) PlayMusic() {
	switch s.StateCtx {
	case NOPOWER:
		fmt.Println("Cant play")
	case LOWPOWER:
		fmt.Println("Play with save mode power")
	case FULLPOWER:
		fmt.Println("Play")
	}
}

func main() {
	speaker := Speaker{}
	speaker.PlayMusic()

	// After Chagre power
	speaker.StateCtx = FULLPOWER
	speaker.PlayMusic()

	// After 1 hour
	speaker.StateCtx = LOWPOWER
	speaker.PlayMusic()
}
