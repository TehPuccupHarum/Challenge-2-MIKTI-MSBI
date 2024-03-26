/* berikut adalah tugas ke 2 yaitu bmi
1. Simpan massa dan tinggi badan Mark dan John dalam variabel.
2. Hitung kedua BMI mereka menggunakan rumus (kamu bahkan dapat menerapkan kedua versi)
3. Buat variabel Boolean 'markHigherBMI' yang berisi informasi tentang apakah Mark memiliki BMI lebih tinggi dari John.

*/

package main

// Buat struct untuk menyimpan data berat dan tinggi badan
type Person struct {
	Name string
	Weight float64 // kg
	Height float64 // m
}

func main() {
	// penggunaan data mark dan john menggunakan struct Person
	mark := Person{Name: "Mark", Weight: 78.0, Height: 1.69}
	john := Person{Name: "John", Weight: 92.0, Height: 1.95}
}
