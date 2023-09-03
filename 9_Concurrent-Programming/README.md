**Concurrency in Golang

* Sequential vs parallel vs concurrent. * Sequential program akan menyelesaikan task sebelumnya terlebih dahulu baru dapat melanjutkan untuk mengerjakan task selanjutnya (dalam waktu load task akan dijumlahkan setiap waktunya, yang mana akan menampilkan hasil setelah setiap task selesai dikerjakan).

* Parallel program dapat menyelesaikan multiple task secara bersamaan dan membutuhkan mesin multi-core (dalam waktu load task akan menggunakan waktu task terlama). 

* Concurent progam dapat menyelesaikan multiple task dan akan dikerjakan secara independen dan mungkin muncul bersamaan (dalam waktu load task akan langsung dimunculkan hasil dari task seetelah task selesai dieksekusi).

* Concurrency pada golang dapat melakukan parallel program yang mengambil keuntungan dari mesin dengan multiple processors.

** Goroutines feature
* Concurrent execution (Goroutines).
* Synchronization and messaging (Channels).
* Multi-way concurrent control (Select).

* Goroutines merupakan fungsi/methods yang berjalan secara concurrently(independent) dengan fungsi/methods lainnya. Cost dalam penggunaan Goroutine lebih kecil dibandingkan thread.

**Gomaxprocs, digunakan untuk mengontrol jumlah thread sistem operasi yang tersedia untuk goroutine execution untuk beberapa program go tertentu.

** GOMAXPROCS, digunakan untuk mengontrol banyaknya operasi sistem thread dan tersedia untu goroutine execution untuk program go tertentu.

** Channel & Select
* channel adalah komunikasi objek menggunakan antara goroutines yang dapat berkomunikasu satu sama lain.

** Race condition terjadi ketika 2 threads mengakses memory diwaktu yang sama, yang mana salah satunya menulis. kondisi ini terjadi karena unsynchronized access to shared memory.
* untuk menangani data race dapat dilakukan (WaitGroups, Channels Blocking, Mutex)

* Blocking With WaitGroups merupakan cara yang secara langsung memperbaiki data race, ia akan memblokir akses baca hingga operasi tulis selesai.

* Channel Blocking mengizinkan goroutines kita untuk sinkronkan satu sama lain sejenak.

* Mutex adalah ketika kita tidak peduli terkait perintah dari membaca dan menulis, kita hanya meminta agar hal tersebut tidak terjadi secara bersamaan. apabila terdengar seperti use case, maka kita akan mempertimbangkan untuk menggunakan mutex.