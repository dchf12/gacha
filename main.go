package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/dchf12/gacha/gacha"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	p := gacha.NewPlayer(10, 100)
	// ※本当はハンドラ間で競合が起きるのでNG
	play := gacha.NewPlay(p)

	http.HandleFunc("/draw", func(w http.ResponseWriter, r *http.Request) {
		if play.Draw() {
			fmt.Fprintln(w, play.Result())
		}

		if err := play.Err(); err != nil {
			fmt.Fprintln(w, err.Error(), http.StatusInternalServerError)

			return
		}

		fmt.Fprintln(w, "残り:", p.DrawableNum())
	})

	return http.ListenAndServe(":8080", nil)
}
