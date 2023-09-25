***Middleware, sebuah entitas yang dipasangkan didalam server yang ketika server melakukan req dan mengrim balasan respon maka akan dipasangkan layer yang berisi fungsi kusus untuk komunikasi data.

** contoh third party middleware : Negroni, Echo, Interpoce, Alice. dan yang akan digunaakn adalah basic autch , jwt dan logger.

** terdapat 2 jenis setup middleware pada echo yaitu:  
* echo#Pre() => akan dieksekusi sebelum root masuk kedalam controller
* echo#Use() => middleware yang akan dieksekusi setelah proses pada router

** Pada BE diusahakan menggunakan protokol https dikarenakan kita akan mengolah data maka dengan menggunakan https dapat mengurangi kemungkinan data kita diretas oleh hacker

** Log Middleware, merupakan middleware untuk mencatat informasi apa saja yang terjadi didalam server.

** Auth Middleware, merupakan proses dalam mengidentifikasi user. Dilakukan untuk mengecek apakah sebuah user itu memiliki izin untuk mengakses api kita atau tidak.

** Basic authentication adalah sebuah proses autentikasi rest api dengan mengirimkan uname dan password melalui header, aturan dalam basic aut adalah mengirim key auth dan valuenya basic + base64encode.

** JWT Middleware, pada jwt kita tidak perlu mengakses db sehingga kita dapat mengecek kevalid-an data dari user berdasarkan token user. JWT dibagi menjadi 3 bagian yaitu encode header berisi algoritma apa yang digunakan dan tipe, pada bagian kedua berupa encode dari payload data dan bagian terakhir adalah gabungan antara token pertama dan kedua lalu ditambahkan secret key yang telah kita set.

* untuk melakukan cek pada postman menggunakan JWT maka menggunakan keyword Bearer + token