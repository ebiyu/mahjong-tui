package game

import (
	"math/rand"
	"sort"
	"time"
)

type Game struct {
	yama      []string
	kawa      [4][]string // e.g. {"1m", "0p", "1z"}
	tehai     [4][]string // e.g. {"1m", "0p", "1z"}
	point     [4]int      // e.g. 25000
	kyokuKaze int         // 0: east, 1: south
	kyokuNum  int         // 1-index
	honba     int         // 0-index
	kyotaku   int         // e.g. 1000
}

func Init() *Game {
	game := &Game{
		yama:  generateYama(),
		tehai: [4][]string{{}, {}, {}, {}},
		kawa: [4][]string{
			{"1p", "2p", "3p", "4p", "5p", "6p", "7p", "8p"},
			{"1m", "2m", "3m", "4m", "5m", "6m", "7m", "8m"},
			{"1s", "2s", "3s", "4s", "5s", "6s", "7s", "8s"},
			{"1z", "1z", "2z", "2z", "3z", "3z", "4z", "5z"},
		},
		point:     [4]int{25000, 25000, 25000, 25000},
		kyokuKaze: 0,
		kyokuNum:  1,
		honba:     0,
		kyotaku:   0,
	}

	chicha := 0
	game.haipy().ripai(chicha).tsumo(chicha)

	return game
}

func (game *Game) Tehai() [4][]string {
	return game.tehai
}

func (game *Game) Kawa() [4][]string {
	return game.kawa
}

func (game *Game) Kill(player int, pieIndex int) {
	game.kawa[player] = append(game.kawa[player], game.tehai[player][pieIndex])
}

func generateYama() []string {
	yama := []string{
		"1m", "1m", "1m", "1m",
		"2m", "2m", "2m", "2m",
		"3m", "3m", "3m", "3m",
		"4m", "4m", "4m", "4m",
		"5m", "5m", "5m", "0m",
		"6m", "6m", "6m", "6m",
		"7m", "7m", "7m", "7m",
		"8m", "8m", "8m", "8m",
		"9m", "9m", "9m", "9m",
		"1s", "1s", "1s", "1s",
		"2s", "2s", "2s", "2s",
		"3s", "3s", "3s", "3s",
		"4s", "4s", "4s", "4s",
		"5s", "5s", "5s", "0s",
		"6s", "6s", "6s", "6s",
		"7s", "7s", "7s", "7s",
		"8s", "8s", "8s", "8s",
		"9s", "9s", "9s", "9s",
		"1p", "1p", "1p", "1p",
		"2p", "2p", "2p", "2p",
		"3p", "3p", "3p", "3p",
		"4p", "4p", "4p", "4p",
		"5p", "5p", "5p", "0p",
		"6p", "6p", "6p", "6p",
		"7p", "7p", "7p", "7p",
		"8p", "8p", "8p", "8p",
		"9p", "9p", "9p", "9p",
		"1z", "1z", "1z", "1z", // east
		"2z", "2z", "2z", "2z", // south
		"3z", "3z", "3z", "3z", // west
		"4z", "4z", "4z", "4z", // north
		"5z", "5z", "5z", "5z", // hatsu
		"6z", "6z", "6z", "6z", // haku
		"7z", "7z", "7z", "7z", // chun
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(yama), func(i, j int) { yama[i], yama[j] = yama[j], yama[i] })
	return yama
}

func (game *Game) haipy() *Game {
	for player := 0; player < 4; player++ {
		for i := 0; i < 13; i++ {
			game.tsumo(player)
		}
	}
	return game
}

func (game *Game) tsumo(player int) *Game {
	pai := game.yama[0]
	game.tehai[player] = append(game.tehai[player], pai)
	game.yama = game.yama[1:]
	return game
}

func (game *Game) ripai(player int) *Game {
	pai := game.tehai[player]
	game.tehai[player] = sortPai(pai)
	return game
}

func sortPai(pai []string) []string {
	sort.Slice(pai, func(i int, j int) bool {
		if pai[i][1] < pai[j][1] {
			return true
		} else if pai[i][1] == pai[j][1] {
			l := pai[i][0]
			r := pai[j][0]
			// 0p == r5p
			if l == '0' {
				if r == '5' {
					// "r5 < 5" is false bacause 5 < r5
					return false
				}
				return '5' < r
			} else if r == '0' {
				if l == '5' {
					// "5 < r5" is true bacause 5 < r5
					return true
				}
				return l < '5'
			} else {
				return l < r
			}
		}
		return false
	})
	return pai
}
