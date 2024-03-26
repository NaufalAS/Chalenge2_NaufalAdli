package main

import (
	"fmt"
)

/*
	disini merupakan struct dengan nama person dan

ada berat dan tinggi sebagai penampung sebuah field
*/
type person struct {
	berat  int
	tinggi float64
}

// method untuk menghitung BMI
func (h person) bmi() float64 {
	return float64(h.berat) / (float64(h.tinggi) * float64(h.tinggi))
}

// method untuk menghitung BMi dan hasil boolean
func CompareBMI(mark person, john person) bool {
	return mark.bmi() > john.bmi()
}

func main() {
	// Inisialisasi objek data1 dengan nilai berat dan tinggi
	beratmark1 := person{berat: 78, tinggi: 1.69}
	beratjhon1 := person{berat: 92, tinggi: 1.95}
	berat := CompareBMI(beratjhon1, beratmark1)

	beratmark2 := person{berat: 95, tinggi: 1.88}
	beratjhon2 := person{berat: 85, tinggi: 1.76}
	berat2 := CompareBMI(beratjhon2, beratmark2)

	// Hitung BMI dan cetak hasilnya
	fmt.Println("BMImark1:", beratmark1.bmi())
	fmt.Println("BMIjohn1:", beratjhon1.bmi())
	fmt.Println("apakah mark BMI lebih tinggi dari jhon?", berat)

	fmt.Println("BMImark2:", beratmark2.bmi())
	fmt.Println("BMIjohn2:", beratjhon2.bmi())
	fmt.Println("apakah mark BMI lebih tinggi dari jhon?", berat2)

}
