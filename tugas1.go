/* tugas 1 yaitu dibawah ini
1. Hitung skor rata-rata untuk setiap tim, menggunakan data uji di bawah ini.
2. Bandingkan skor rata-rata tim untuk menentukan pemenang kompetisi, dan cetak ke konsol. Jangan lupa bahwa bisa ada hasil seri, jadi uji juga untuk itu (seri berarti mereka memiliki skor rata-rata yang sama).
3. Bonus 1: Sertakan persyaratan untuk skor minimum 100. Dengan aturan ini, sebuah tim hanya menang jika memiliki skor lebih tinggi dari tim lain, dan pada saat yang sama skor minimal 100 poin. Petunjuk: Gunakan operator logika untuk menguji skor minimum, serta beberapa blok else-if.
4. Bonus 2: Skor minimum juga berlaku untuk seri! Jadi seri hanya terjadi ketika kedua tim memiliki skor yang sama dan keduanya memiliki skor lebih besar atau sama dengan 100 poin. Jika tidak, tidak ada tim yang memenangkan trofi.
*/

package main

// hitungRataRata adalah func menghitung rata rata dai slice integer.
func hitungRataRata(scores []int) int {
	if len(scores) == 0 {
		return 0
	}
	total := 0
	for _, score := range scores {
		total += score
	}
	return total / len(scores)
}

// tentukanPemenang adalah func menentukan pemenang dari hasil skor rata - rata dan nilai minimum.
func tentukanPemenang(rataRata1, rataRata2, minScore int) string {
	if rataRata1 >= minScore && rataRata1 > rataRata2 {
		return "Lumba - Lumba"
	} else if rataRata2 >= minScore && rataRata2 > rataRata1 {
		return "Koala"
	} else if rataRata1 == rataRata2 && rataRata1 >= minScore {
		return "Seri"
	}

}

func main() {
	// Data Utama
	lumbaLumba := []int{96, 108, 89}
	koala := []int{88, 91, 110}

	// Data Bonus
	bonusData := []struct {
		LumbaLumba []int
		Koala      []int
	}{
		{[]int{97, 112, 101}, []int{109, 95, 123}}, //Bonus 1
		{[]int{97, 112, 101}, []int{109, 95, 106}}, //Bonus 2
	}

	
}
