package domain

import (
	"big2/domain/enum"
	"fmt"
	"reflect"
)

type Big2 struct {
	round                  int
	topPlay                CardPattern
	topPlayer              *IPlayer
	winner                 *IPlayer
	players                []*IPlayer
	deck                   *Deck
	makeCardPatternHandler *IMakeCardPatternHandler
}

func NewBig2(makeCardPatternHandler *IMakeCardPatternHandler, deck *Deck) *Big2 {
	return &Big2{
		round:                  1,
		deck:                   deck,
		makeCardPatternHandler: makeCardPatternHandler,
	}
}

func (big2 *Big2) GenerateHumanPlayer() {
	player := &IPlayer{
		Player: NewHumanPlayer(len(big2.players), big2.makeCardPatternHandler),
	}
	big2.players = append(big2.players, player)
	player.nameSelf()
}

func (big2 *Big2) GenerateAIPlayer() {
	player := &IPlayer{
		Player: NewAIPlayer(len(big2.players), big2.makeCardPatternHandler),
	}
	big2.players = append(big2.players, player)
	player.nameSelf()
}

func (big2 *Big2) playerDrawCards() {
	count := 0
	for {
		if len(big2.deck.Cards()) == 0 {
			break
		}
		deck := big2.deck
		card := deck.deal()
		p := big2.players[count%4]
		p.addCardIntoHand(card)
		count += 1
	}

}

func (big2 *Big2) Start() {

	big2.playerDrawCards()

	for big2.winner == nil {
		big2.takeRound()
	}

	fmt.Printf("遊戲結束，遊戲的勝利者為 %s\n", big2.winner.Name())
}

func (big2 *Big2) takeRound() {

	fmt.Println("新的回合開始了。")

	var player *IPlayer

	if big2.round == 1 { // 第一回合
		player = big2.getPlayerHasClub3()
		fmt.Printf("player %s has Club3\n", player.Name())
	} else {
		player = big2.topPlayer
		fmt.Printf("top player is %s\n", player.Name())
	}

	turn := 0
	passCount := 0
	for passCount < 3 {

		var cardPattern CardPattern

		for {
			cardPattern = player.takeTurn(turn)

			// 第一回合首輪玩家只能出包含梅花3的牌型
			if big2.round == 1 && turn == 0 && !cardPattern.containsClub3() {
				fmt.Println("出的牌型不包含梅花3")
				continue
			}

			// 如果有出牌(沒有 pass)，在有頂牌的狀況下，牌型必須和頂牌一樣，且比頂牌大
			if cardPattern != nil && big2.topPlay != nil {
				if reflect.TypeOf(big2.topPlay).String() != reflect.TypeOf(cardPattern).String() {
					fmt.Printf("牌型和頂牌不一樣，不可出該副牌, top Play: %s\n", big2.topPlay)
					continue
				}
				if !cardPattern.isBiggerThan(big2.topPlay) {
					fmt.Printf("沒有比頂牌大，不可出該副牌, top Play: %s\n", big2.topPlay)
					continue
				}
			}
			break
		}

		if cardPattern != nil {
			fmt.Printf("玩家 %s 打出了 %s %s\n", player.Name(), cardPattern.Name(), cardPattern.printCards())
			for _, card := range cardPattern.Cards() {
				player.removeCardFromHand(card)
			}
			// 打完牌後檢查玩家還有沒有剩餘的牌
			if len(player.Hand()) == 0 {
				big2.winner = player
				break
			}
			// 更新頂牌及頂牌玩家
			big2.topPlay = cardPattern
			big2.topPlayer = player
			fmt.Printf("top Play: %s %s\n", big2.topPlay.Name(), big2.topPlay.printCards())

			// 玩家出完新的牌後才開始起算接下來有幾個 pass
			passCount = 0
		} else {
			passCount += 1
			fmt.Printf("玩家 %s PASS.\n", player.Name())
		}

		nextPlayerId := (player.Id() + 1) % len(big2.players)
		player = big2.players[nextPlayerId]
		turn += 1
	}

	big2.topPlay = nil
	big2.round += 1

}

func (big2 *Big2) getPlayerHasClub3() *IPlayer {
	for _, player := range big2.players {
		for _, card := range player.Hand() {
			if card.Suit == enum.Club && card.Rank == enum.Three {
				return player
			}
		}
	}
	return nil
}
