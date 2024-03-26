// package main

// import (
// 	"fmt"
// )

// // disini merupakan struct dengan nama pemnang dan
// // ada data, data1, dan data2 sebagai penampung sebuah field
// type pemenang struct {
// 	data, data1, data2 int
// }

// /*
// disini merupakan method dengan nama rata untuk menghitung rata rata
// dari semua data set
// */
// func (h pemenang) rata() int {
// 	return (h.data + h.data1 + h.data2) / 3
// }

// /*
// disini merupaka function perbadingan untuk membandingkan hasil yang suda di lakukan di
// method rata()
// */
// func perbandingan(lumba pemenang, koala pemenang) string {
// 	lumbarata := lumba.rata()
// 	koalarata := koala.rata()
// 	if lumbarata > koalarata {
// 		return "yang menang lumba"
// 	} else if lumbarata < koalarata {
// 		return "yang menang koala"
// 	} else {
// 		return "hasilnya imbang"
// 	}
// }

// /*
// disini merupaka function perbadingan1 untuk membandingkan hasil yang suda di lakukan di
// method rata() tapi untuk data bonus 1
// */
// func perbandingan1(lumba1 pemenang, koala1 pemenang) string {
// 	lumbarata1 := lumba1.rata()
// 	koalarata1 := koala1.rata()
// 	if lumbarata1 >= 100 || koalarata1 >= 100 {
// 		if lumbarata1 > koalarata1 {
// 			return "yang menang lumba"
// 		} else if lumbarata1 < koalarata1 {
// 			return "yang menang koala"
// 		} else {
// 			return "hasilnya imbang"
// 		}
// 	} else {
// 		return "tidak ada yang menang"
// 	}
// }

// /*
// disini merupaka function perbadingan2 untuk membandingkan hasil yang suda di lakukan di
// method rata() tapi untuk data bonus 2
// */
// func perbandingan2(lumba2, koala2 pemenang) string {
// 	lumbarata2, koalarata2 := lumba2.rata(), koala2.rata()

// 	if lumbarata2 == koalarata2 && lumbarata2 >= 100 && koalarata2 >= 100 {
// 		return "hasilnya imbang"
// 	} else if lumbarata2 >= 100 && koalarata2 >= 100 {
// 		switch {
// 		case lumbarata2 > koalarata2:
// 			return "yang menang lumba"
// 		case lumbarata2 < koalarata2:
// 			return "yang menang koala"
// 		}
// 	}
// 	return "tidak ada yang memenuhi skor minimum"
// }

// func main() {
// 	//data set 1
// 	lumba := pemenang{data: 96, data1: 108, data2: 89}
// 	koala := pemenang{data: 88, data1: 91, data2: 110}

// 	//data data bonus 1
// 	lumba1 := pemenang{data: 97, data1: 112, data2: 101}
// 	koala1 := pemenang{data: 109, data1: 95, data2: 123}

// 	//data data bonus 2
// 	lumba2 := pemenang{data: 97, data1: 112, data2: 101}
// 	koala2 := pemenang{data: 109, data1: 95, data2: 106}

// 	//hasil data set 1
// 	fmt.Println("hasil rata rata", lumba.rata())
// 	fmt.Println("hasil rata rata", koala.rata())
// 	fmt.Println("hasilnya dimennagkan oleh", perbandingan(lumba, koala))

// 	//hasil data bonus 1
// 	fmt.Println("hasil rata rata", lumba1.rata())
// 	fmt.Println("hasil rata rata", koala1.rata())
// 	fmt.Println("hasilnya dimennagkan oleh", perbandingan1(lumba1, koala1))

// 	//hasil data bonus 2
// 	fmt.Println("hasil rata rata", lumba2.rata())
// 	fmt.Println("hasil rata rata", koala2.rata())
// 	fmt.Println("hasilnya dimennagkan oleh", perbandingan2(lumba2, koala2))
// }
