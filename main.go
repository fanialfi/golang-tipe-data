/*
  program file ini menjelaskan beberapa tipe data yang didukung oleh golang
  seperti tipe data numeric non-desimal, dan numeric desimal
  tipe data boolean
  tipe data string
  nilai nil dan zero value
*/

package main

import (
	"fmt"
	"reflect"
)

func numNotDecUint(){
  // cara menghitungnya 2^exponent - 1
  var uint8 uint8 = 255; // jarak jangkauannya dari 0 - 255,sama dengan byte
  var uint16 uint16 = 65535; // jarak jangkauannya dari 0 - 6553
  var uint32 uint32 = 4294967295 // atau 2^32 - 1 =  4294967295
  var uint64 uint64 = 18446744073709551615 // 2^64 - 1 = 18446744073709551615
  var uintIni uint = 18446744073709551615 // nilai nya sama dengan uint32 atau uint64

  fmt.Printf("variabel uint8 memiliki isi %d dan berjenis %s\n",uint8,reflect.TypeOf(uint8))
  fmt.Printf("variabel uint16 memiliki isi %d dan berjenis %s\n",uint16,reflect.TypeOf(uint16))
  fmt.Printf("variabel uint32 memiliki isi %d dan berjenis %s\n",uint32,reflect.TypeOf(uint32))
  fmt.Printf("variabel uint64 memiliki isi %d dan berjenis %s\n",uint64,reflect.TypeOf(uint64))
  fmt.Printf("variabel uintIni memiliki isi %d dan berjenis %s\n",uintIni,reflect.TypeOf(uintIni))
}

func numNotDecInt(){
  // cara menghitungnya -2^exponent sampai (2^exponent - 1)
  var intDelapan int8 = -128 // -2^7 sampai 2^7 - 1 = -128 - 127
  var intEnamBelas int16 = 32767 // -2^15 sampai 2^15 - 1 = -32768 - 32767
  var intTigaDua int32 = 2147483647 // -2^31 sampai 2^31 - 1 = (-2147483648, 2147483647)
  var intEnamEmpat int64 = 9223372036854775807 // -2^63 sampai 2^63 - 1 = (-9223372036854775808, 9223372036854775807)
  var intIni int = 30000000000 // tipenya sama dengan int32 atau int64 tergantung nilainya

  fmt.Printf("variabel intDelapan memiliki isi %d dan berjenis %s\n",intDelapan,reflect.TypeOf(intDelapan))
  fmt.Printf("variabel intEnamBelas memiliki isi %d dan berjenis %s\n",intEnamBelas,reflect.TypeOf(intEnamBelas))
  fmt.Printf("variabel intTigaDua memiliki isi %d dan berjenis %s\n",intTigaDua,reflect.TypeOf(intTigaDua))
  fmt.Printf("variabel intEnamEmpat memiliki isi %d dan berjenis %s\n",intEnamEmpat,reflect.TypeOf(intEnamEmpat))
  fmt.Printf("variabel intIni memiliki isi %d dan berjenis %s\n",intIni,reflect.TypeOf(intIni))
}

func tanpaDecType(){
  // jika sebuah variabel berisi bilangan, baik bilangan negatif ataupun positif akan berjenis int
  var satu = -100;
  var dua = 100

  fmt.Printf("variabel satu memiliki isi %d dan berjenis %s\n",satu,reflect.TypeOf(satu))
  fmt.Printf("variabel dua memiliki isi %d dan berjenis %s\n",dua,reflect.TypeOf(dua))
}

func float(){
  var flt = 2.16 // secara default akan menggunakan tipe data float64
  var flt32 float32 = 100.10

  fmt.Printf("variabel flt berisi %.4f dan berjenis %s\n",flt,reflect.TypeOf(flt))
  // fmt.Printf("variabel flt32 berisi %.4f dan berjenis %s\n",flt32,reflect.TypeOf(flt32))
  fmt.Printf("variabel flt32 berisi %.4f dan berjenis %T\n",flt32,flt32)
}

func main(){
  fmt.Println("Tipe data non desimal uint")
  numNotDecUint()
  fmt.Println("================================")
  fmt.Println("Tipe data non desimal int")
  numNotDecInt()
  fmt.Println("=================================")
  fmt.Println("tanpa deklarasi tipe data")
  tanpaDecType()
  fmt.Println("==================================")
  fmt.Println("==================================")
  float()
}

/*
  template %d di atas digunakan untuk memformat tipe data berjenis non-desimal
  template %f di atas digunakan untuk memformat tipe data berjenis desimal dengan 6 digit di belakang
  jika ingin membatasi jumplah digit desimal yang dihasilkan, bisa dikontrol dengan menggunakan %.nf dimana n adalah jumplah number-nya, seperti kasus di atas diberikan nilai 4, maka digit desimal-nya hanya berisi 4
  template format specifier %T merepresentasikan type data dari value, format specifier ini berada didalam package fmt, jadi jika ingin menggunakannnya, harus import package fmt terlebih dahulu

  lebih lanjut mengenai format specifier di golang, bisa dibaca di https://pkg.go.dev/fmt
*/
