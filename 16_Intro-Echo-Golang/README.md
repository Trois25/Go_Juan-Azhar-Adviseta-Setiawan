** Pengenalan Echo

* Third party atau library merupakan kumpulan dari sebuat code yang memiliki fungsi yang dapat dipanggil oleh program lain, apabila kita kebanyakan menggunakan library maka kita akan keterbegantungan terhadap library sehingga ketika mereka melakukan update maka kita juga harus merombak code kita sesuai update dari third party tersebut.

* Echo adalah web framework golang dengan performa tinggi, extensible dan minimalist. Apabila dibandingkan oleh framework lain maka echo lebih cepat dalam operasinya, dan echo ini merupakan framework untuk mengurus REST API.

* Echo memiliki optimized router, middleware(banyak middleware yang dapat digunakan oleh Echo), data rendering(ketika memberikan response echo dapat menggunakan costume handler), scalable (dapat digunakan untuk skala besar atau kecil), data binding (Echo dapat lebih mudah untuk mendapatkan data request dari user),

* Echo merupakan framework yang minimalis namun extensible (tidak berukuran besar namun powerfull, namun dalam penggunaan ORM/db Echo masih belum dapat menghandlenya sehingga dibutuhkan library tambahan seperti GORM), echo tidak memiliki struktur sehingga dev lebih bebas dalam membangun strukturnya.

* Untuk menggunakan echo hal pertama yang harus dilakukan adalah melakukan init untuk go.mod dengan command "go mod init *nama_folder", lalu menambahkan echo kedalam go.mod dengan cara "go get -u github.com/labstack/echo/v4".

* Setelah melakukan configurasi diatas maka selanjutnya adalh membuat routing, pada func main bisa digunakan instance yang dideklarasikan dan membuat method yang akan digunakan lalu menerima 2 params yaitu url path dan fungsi routing.

* pada fungsi routing akan menggunakan parameter e echo.Context yang digunakan untuk menerima data dari client dan juga digunakna untuk memberikan response kepada client, fungsi controller akan diberikan kembalian error untuk report kepada user.

* lalu pada fungsi main dapat digunakan fungsi start dengan mendeklarasikan port yang akan digunakan, contoh : (e.start("
:8000"))