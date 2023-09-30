* Unit Testing
** software testing adalah proses untuk menganalisa suatu software untuk melihat perbedaan antara feat yang tersedua dan yang akan dibuat, tujuan software testing adalah memastikan software saat ini telah berjalan dengan baik.

** Tujuan testing
* Preventing Regression, ketika software yang telah jadi menjadi rusak karena penambahan feat yang salah.
* confidence level in refactoring, dengan adanya unit testing kita akan lebih percaya diri dengan refactoring kita.
* improve code design
* code documentation
* scaling the team

** Lelvel testing
* UI, dilakukan pada tampilan aplikasi untuk melihat flow aplikasi
* Integration, melakukan testing terhadap feat spesifik
* Unit, melakukan testing dalam unit terkecil didalam aplikasi contohnya adalah testing terhadap fungsi atau method

** Framework, adalah sebuah kerangka kerja yang digunakan dalam unit testing. dengan menggunakan framework kita dapat fokus membuat testing untuk aplikasi kita dan tidak perlu memikirkan terkait library yang harus digunakan. Framework ini terspesifikasi oleh bahasa pemrograman apa yang digunakan jadi ketika melakukan unit testing akan lebih spesifik.

** Struktur, struktur dalam unit testing adalah strategi dalam meletakkan file testing dalam code based application. 2 pattern yg biasa digunakan dalam unit testing :
1. centralize, meletakkan file didalam 1 folder. keuntungannya adalah directory akan lebih clean.
2. menaruh file testing disamping setiap file yang mau di test. keuntungannya file test akan lebih mudah ditemukan.

** Komponen penting didalam unit testing :
* Test file, merupakan file dimana code kita dibuat yang berisi test suites
* Test suites adalah collection test case
* Test fixtures adalah script untuk memastikan environtment yang digunakan tetap konsisten
* Test case, kumpulan scenario yang dijalankan untuk melakukan test dalam program

** Runner merupakan sebuah aplikasi yang didesign untuk menjalankan test. tools yang terdapat didalam runners :
* feature untuk menjalankan file test
* watch mode, yakni feature yang melakukan test secara otomatis
* memilih runner yang lebih cepat

** Mocking, merupakan sebuah imitasi yang melakukan dependencies kedalam api lain maka testing akan terpengaruh dan tak akan sesuai maka kita akan membuat objek tiruan berdasarkan api yang digunakan shingga program tidak mengalami ketergantungan.

** Testing harus independent karena kita harus fokus kepada code yang kita buat.

** Hal yang tidak disarankan didalam unit testing :
* Third party modul
* Database
* Third party api atau sistem external lain
* object class yang harus dilakukan test ditempat lain

** Testing Coverage adalah alat ukur untuk menunjukkan source code mana saja yang telah dijalankan.

** Bagian yang diexecute oleh coverage antara lain :
* Function 
* Statement
* Branch
* Lines

** Format report dapat berbentuk :
* CLI
* XML
* HTML
*LCOV

** selain menggunakan coverage test kita dapat menggunakan sonarQube

** Simple steps to test
* Create a new test file in go dengan nama file library_test.go lalu untuk lokasi file di folder yang sama package sama atau package beda
* Tulis fungsi test dengan nama fungsi func TestNamafungsi(t *testing.T)
* pahami apa yang anda buat dengan membuat lebih dari 1 test case dan ikuti aturan testing dalam go
* untuk menjalankannya dapat menggunakan command : "go test tes.go tes_test.go -cover" yang mana pada kasus diatas tes adalah file yang akan dilakukan tes lalu tes_test merupakan file unit test

* Testing Scenario (apa yang mau dilakukan test) dan Test Case (bagaimana cara kita akan melakukan test)
** Testing scenario(high level test case) adalah gambaran umum terhadap apa yang mau dites
** Test case adalah kumpulan langkah2 uji positif dan negatif dalam test scenario, contoh :
* Scenario : check login functionality
* test case 1 : masukkan email dan password yang valid
* test case 2 : masukkan valid email dan invalid password
* test case 3 : masukkan invalid email dan valid password
* test case 4 : masukkan email dan password yang invalid