package main

import (
	"big2/domain"
)

func main() {

	makeCardPatternHandler := domain.NewIMakeCardPatternHandler(domain.NewMakeSingleCardPatternHandler(
		domain.NewIMakeCardPatternHandler(domain.NewMakePairCardPatternHandler(
			domain.NewIMakeCardPatternHandler(domain.NewMakeStraightCardPatternHandler(
				domain.NewIMakeCardPatternHandler(domain.NewMakeFullHouseCardPatternHandler(nil)),
			)),
		)),
	))

	deck := domain.NewDeck()
	deck.Shuffle()
	//deck := domain.NewDeckFromCardsInput("D[7] C[A] S[6] S[4] S[A] S[J] C[10] C[K] D[4] H[9] D[J] D[K] C[7] H[8] C[3] S[K] S[3] D[2] C[8] C[4] H[Q] C[J] D[3] D[6] D[9] D[A] H[6] S[7] H[7] C[6] H[3] C[Q] H[J] H[10] S[9] D[10] C[9] D[8] D[Q] H[5] H[K] C[5] D[5] C[2] S[10] S[5] S[Q] S[8] H[2] H[4] H[A] S[2]")
	big2 := domain.NewBig2(makeCardPatternHandler, deck)

	big2.GenerateHumanPlayer()
	big2.GenerateHumanPlayer()
	big2.GenerateHumanPlayer()
	big2.GenerateHumanPlayer()

	big2.Start()

}
