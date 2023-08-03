package main

import (
	"C3M1H1/domain"
	"time"
)

func main() {

	pewDiePie := domain.NewChannel("PewDiePie")
	waterBallClub := domain.NewChannel("水球軟體學院")

	waterBall := domain.NewWaterBall()
	fireBall := domain.NewFireBall()

	waterBallClub.Subscribe(waterBall)
	pewDiePie.Subscribe(waterBall)

	waterBallClub.Subscribe(fireBall)
	pewDiePie.Subscribe(fireBall)

	waterBallClub.Upload(domain.NewVideo("C1M1S2", "這個世界正是物件導向的呢！", 4*time.Minute))
	pewDiePie.Upload(domain.NewVideo("Hello guys", "Clickbait", 30*time.Second))

	waterBallClub.Upload(domain.NewVideo("C1M1S3", "物件 vs. 類別", 1*time.Minute))
	pewDiePie.Upload(domain.NewVideo("Minecraft", "Let’s play Minecraft", 30*time.Minute))
}
