/* berikut adalah tugas ke 2 yaitu bmi
1. Simpan massa dan tinggi badan Mark dan John dalam variabel.
2. Hitung kedua BMI mereka menggunakan rumus (kamu bahkan dapat menerapkan kedua versi)
3. Buat variabel Boolean 'markHigherBMI' yang berisi informasi tentang apakah Mark memiliki BMI lebih tinggi dari John.

*/

package main

import "fmt"

// Buat struct untuk menyimpan data berat dan tinggi badan
type Person struct {
	Name   string
	Weight float64 // kg
	Height float64 // m
}

// Metode untuk menghitung BMI
func (p *Person) hitungBMI() float64 {
	return p.Weight / (p.Height * p.Height)
}

// Metode untuk membandingkan BMI
func (p *Person) bandingkanBMI(other *Person) bool {
	return p.hitungBMI() > other.hitungBMI()
}

func main() {
	// penggunaan data mark dan john menggunakan struct Person
	mark := Person{Name: "Mark", Weight: 78.0, Height: 1.69}
	john := Person{Name: "John", Weight: 92.0, Height: 1.95}

	//hitung  dan bandingkan BMI untuk data 1
	markHigherBMI1 := mark.bandingkanBMI(&john)

	//menampilkan hasil data 1
	fmt.Printf("Data 1:\n")
	fmt.Printf("BMI %s: %.2f\n", mark.Name, mark.hitungBMI())
	fmt.Printf("BMI %s: %.2f\n", john.Name, john.hitungBMI())
	fmt.Printf("%s memiliki BMI lebih tinggi dari %s : %t\n \n ", mark.Name, john.Name, markHigherBMI1)

	//data baru untuk data 2
	mark2 := Person{Name: "Mark", Weight: 95.0, Height: 1.88}
	john2 := Person{Name: "John", Weight: 85.0, Height: 1.76}

	//hitung dan bandingkan data BMI untuk data 2
	markHigherBMI2 := mark.bandingkanBMI(&john2)

	//menampilkan hasil data 1
	fmt.Printf("Data 2:\n")
	fmt.Printf("BMI %s: %.2f\n", mark2.Name, mark2.hitungBMI())
	fmt.Printf("BMI %s: %.2f\n", john2.Name, john2.hitungBMI())
	fmt.Printf("%s memiliki BMI lebih tinggi dari %s : %t\n \n ", mark2.Name, john2.Name, markHigherBMI2)
}
