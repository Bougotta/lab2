package main

import (
	"bufio"
	"log"
	"os"
	"time"
)

func msg(s string) {
	log.Println(s)
}

func main() {
	// таймер, который вызывает функцию один раз через указанное время
	timeoutTimer, err := NewTimeout(msg, 2*time.Second)
	if err != nil {
		log.Fatal(err)
	}

	// таймер, который вызывает функцию с указанной периодичностью
	intervalTimer, err := NewInterval(msg, 2*time.Second)
	if err != nil {
		log.Fatal(err)
	}

	timeoutTimer.Run("tick")
	intervalTimer.Run("tock")
	log.Println("timers running")

	// ждем ввода от пользователя
	log.Println("press enter or ctrl+c to stop interval timer")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()

	intervalTimer.Stop()
}
