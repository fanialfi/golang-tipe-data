# tipe data yang didukung oleh Golang

beberapa jenis tipe data yang ada di GO diantaranya tipe data `numeric` (desimal & non-desimal), string, boolean.

1. tipe data `numeric` non-desimal

secara umum ada 2 tipe data yang ada di dalam kategori ini yang perlu di ketahui, yaitu :

- `uint`, tipe data untuk bilangan cacah (bilangan positif)
- `int`, tipe data untuk bilangan bulat (bilangan negatif dan positif)

pembagian beberapa jenis tipe data di atas berdasarkan lebar cangkupan nilainya :

| tipe data | keterangan                                               |
| --------- | -------------------------------------------------------- |
| `uint8`   | 0 - 255                                                  |
| `uint16`  | 0 - 65535                                                |
| `uint32`  | 0 - 4294967295                                           |
| `uint64`  | 0 - 18446744073709551615                                 |
| `uint`    | sama dengan `uint32` atau `uint64` (tergantung nilainya) |
| `byte`    | sama dengan `uint8`                                      |
| `int8`    | -128 - +127                                              |
| `int16`   | -32768 - +32767                                          |
| `int32`   | -2147483648 - +2147483647                                |
| `int64`   | -9223372036854775808 - +9223372036854775807              |
| `int`     | sama dengan `int32` atau `int64` (tergantung nilainya)   |
| `rune`    | sama dengan `int32`                                      |

ketika membuat sebuah program, dianjurkan tidak sembarangan menggunakan tipe data variabel, karena bisa berdampak kepada alokasi memori.

2. tipe data `numeric` desimal

|tipe data| keterangan|
|----|----|
|`float32`| cangkupannya -3.4e+38 to 3.4e+38 atau -3.4 kali 10 pangkat 38 sampai 3.4 kali 10 pangkat 38|
|`float64`| cangkupannya -1.7e+308 to +1.7e+308 atau -1.7 kali 10 pangkat 308 sampai 1.7 kali 10 pangkat 308|


> secara default, jika tidak di tentukan jenis tipe data-nya, maka secara default go akan menggunakan `float64`.

3. tipe data boolean

merupakan tipe data yang berisi kebenaran (`false` atau `true`), sering di manfaatkan dalam seleksi kondisi dan perulanan.

4. tipe data string

merupakan tipe data yang berisi sekumpulan karakter, dan disimpan didalam tanda kurung dua ("), atau _backtics_ (``), jika menggunakan _backtics_ ketika memerlukan sebuah enter, tab, dll, tidak perlu lagi menggunakan karakter _escape sequence_ 
