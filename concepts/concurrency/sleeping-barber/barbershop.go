package main

import (
	"time"

	"github.com/fatih/color"
)

type BarberShop struct {
	ShopCapacity    int
	HairCutDuration time.Duration
	NumberOfBarbers int
	BarbersDoneChan chan bool
	ClientsChan     chan string
	Open            bool
}

func (b *BarberShop) addBarber(barber string) {
	b.NumberOfBarbers++

	go func() {
		isSleeping := false
		color.Yellow("%s goes to the waiting room to check for clients", barber)

		for {
			// if there are no clients, the barber goes to sleep
			if len(b.ClientsChan) == 0 {
				color.Yellow("there is nothing to do, so %s takes a nap", barber)
				isSleeping = true
			}

			client, shopOpen := <-b.ClientsChan

			if shopOpen {
				if isSleeping {
					color.Yellow("%s wakes %s up", client, barber)
					isSleeping = false
				}
				// cut hair
				b.cutHair(barber, client)
			} else {
				// shop is closed, so send the barber home (and close its goroutine)
				b.sendBarberHome(barber)
				return
			}
		}
	}()
}

func (b *BarberShop) cutHair(barber, client string) {
	color.Green("%s is cutting %s's hair", barber, client)
	time.Sleep(b.HairCutDuration)
	color.Green("%s finished cutting %s's hair", barber, client)
}

func (b *BarberShop) sendBarberHome(barber string) {
	color.Cyan("%s is going home", barber)
	b.BarbersDoneChan <- true
}

func (b *BarberShop) closeShopForDay() {
	color.Cyan("Closing shop for the day.")

	close(b.ClientsChan)
	b.Open = false

	for range b.NumberOfBarbers {
		<-b.BarbersDoneChan
	}

	close(b.BarbersDoneChan)

	color.Green("---------------------------------------------------------------------")
	color.Green("The barbershop is now closed for the day, and everyone has gone home.")
}
