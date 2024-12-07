package functions

func toplavariadic(sayilar ...int) int { // biriden fazala sayısı bellli olmayan tane int değişkeni yazabilirz denir ve go bunları bir array gibi alır
	toplam := 0

	for i := 0; i < len(sayilar); i++ {
		toplam = toplam + sayilar[i]
	}

	return toplam

}
