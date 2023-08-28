**String

* menggunakan string diperlukan import "strings"
* menghitung panjang string dapat digunakan keyword len(string)
* string juga dapat dilakukan operasi comparison
* untuk membandingkan apakah strung tersebut termasuk bagian string lainnya. dapat digunakan keyword contain => var:=strings.Contains(val1,val2)
* mengambil value dari string lain dapat dilakukan dengan cara => var:=value[index:index]
* untuk mengganti isi string dapat digunakan keyword strings.replace()

* insert digunakan untuk menambahkan string kepada var string yang telah tersedia

** Advance Function
*Variadic function, memanggil function dengan jumlah parameter yang berbeda dann dapat dilakukan dengan membuat temporary slice
* Anonymous Function, yaitu function tanpa nama yang digunakan untuk mengumpulkan/mengelompokkan code
* Closure sebuah tindak lanjut dari anonymous function yang ketika dipakai diluar maka akan dilakukakn reference
* Defer Function merupak penjabaran dari anonymous yang mana kelompok code akan di run di akhir, untuk pengeluaran dari defer menggunakan konsep stack LIFO

** Pointer, merupakan variabel yang dapat menyimpan memorry address variabel lain 

** Struct, memanggil object didalam golang yang berisi kumpulan dari variabel atau fungsi

** Method, method merupakan fungsi yang terikat oleh sebuah type (bisa berupa struct atau tipe yang lain). mengapa menggunakan method? method lebih baik dalam membuat OOP style dalam golang, method membantu untuk menghindari name conflict, memanggil method lebih mudah untuk dibaca dan dipahami dibanding fungsi.