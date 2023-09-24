* ORM(Object Relational Mapping) yang mana akan membantu dalam mengkonversi data dari data tabel menjadi struck/slice

** ORM yang akan digunakan adalah gorm

** keuntungan ORM :
* Less repetitive querry
* Melakukan fetch data secara otomatis untuk tersimpan kedalam variabel objek
* ketika melakukan validasi data akan lebih mudah karena yang akan diisi didalam struck dapat terlihat 
* memiliki fitur cache query

** kelemahan ORM :
* harus menggunakan library pihak ke-3 dan menambahkan cost didalam program
* dapat terjadi eksekusi query yang tidak diinginkan
* ketika menggunakan ORM maka kita tergantung pada pihak ke-3 dan terkadang terdapat sebuah fiture yang belum tersedia di ORM yang kita gunakan

** GORMm adalah ORM berbasis SQL

* cara melakukan koneksi dari aplikasi dengan database adalah dengan cara : "<uusername>:<password>@/<db_name>?charset=utf8&parseTime=True&loc=Local"

* database migration penting karena ketika kita melakukan update dalam database maka migrate dapat dijadikan tracking untuk melihat perubahan didalam tabel. didalam GORM ada fitur otomatis db migration.

** mengapa menggunakan db migretion?
* lebih simpel ketika melakukan update db
* dalam melakukan rollback lebih simple
* dapat melakukan track pada perubahan tabel
* dapat melihat siapa yang melakukan update didalam migration
* applikasi akan selalu kompetible dengan db

* Menggunakan GORM
** instalasi dengan command "go get -u gorm.io/gorm"
** install driver agar GORM dapat mengakses sql dengan command "go get -u gorm.io/driver/mysql"
** cara mealkukan koneksi kedalam database adalah : 
* dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
  db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
** setelah melakukan connection kedalam database selanjutnya membuat struct yang diawali dengan huruf besar agar dapat diakses dari luar file dan melakukan migration

** menggunakan MVC
* folder config diisi file untuk melakukan configurasi didalam projek, sebagai contoh adalah migration dan membuat db
* folder controller berisi controller yang akan dibuat/fungsi dari API
* folder model berisi struct dari tabel
* folder route berisi route/ path API yang akan dibuat