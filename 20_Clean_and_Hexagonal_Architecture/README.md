* Clean architecture, arsitektur yg baik adalah ketika kita memisahkan kebutuhan secara layer dan membuat modular dan dapat dilakukan scalable, maintanable dan test application.

* MVC and MVVM
** dalam mvc maka akan menjalankan :
* user interface with a view akan mengakses layer View
* view alerts controller of particullar event dari user akan mengakses layer Controller
* controller akan update model dan mengupdate layer Model
* Model alerts view that it has change akan mengembalikannya kelayer View 
* view grabs model data and updates itself ke layer Model

** Hexagonal architecture merupakan awal mula dari layer architecture yang dibagi 3 yaitu API(semua yang berhubungan dengan API berada disini), Domain dan SPI(yang berhubungan dengan database berada di SPI)

** The constraint before designing the clean architecture are :
* Independent of Frameworks
* Testable (dapat dilakukan tes bahkan jika tanpa dependensinya)
* Independent of UI (kita dapat merubah UI tanpa merubah business logic)
* Independent of Database (business layer tidak terlalu terikat dengan database, sehingga kita dapat mengganti ke database lain dengan mudah)
* Independent of any external (business logic sebaiknya tidak terikat dengan external)

** The benefit
* struktur yang standar, sehingga akan mempermudah untuk menemukan alur pada projek
* melakukan develop dengan cepat untuk jangka waktu yang panjang
* mocking dependencies becomes trivial in unit tests
* easy switching from prototypes to proper solutions (ct : changing in-memory storage to an SQL database).

** We will define classic 3-layer architecture (we could have more layers)
* Entities layer - berisi business object yang merefleksikan bisnis kita.
* Use case (domain layer/services) berisi business logic.
* controller (present data ke layar, dan disinilah adanya framework)
* drivers / data layers - memanage data ke third party seperti database.

** Domain Driven Design (DDD), dengan menggunakan DDD kita dapat mengelompokkan yang mana yang dapat dijadidak entity baru.


* Context Golang
** context merupakan package yang membawa deadline, cancellation signal, atau value lain pada request/permintaan API.
* untuk mengimplementasikan context kita dapat membuat root context dengan fungsi background (var ctx = context.Background()), fungsi background akan membalikkan root context dimana kita dapat menggunakannya untuk operasi lain.

** Context with value, context value akan sering kita lihat pada middleware. kita dapat mengirim data dari midleware(contoh user_id dari token) ke handler menggunakan context, contoh : ctx = context.WithValue(ctx, "key", "value").

** Context with cancelation
* normal request : request -> service -> database
* user cancel request : request -> service -> (user cancel request, tetapi data masih tetap diteruskan ke database) -> database
 * apabila tanpa menggunakan context cancellation maka data tidak diteruskan kedatabase ketika pada alur service -> database.

* MVC to Clean Code
** MVC, tujuan dari penggunaan clean code adalah kode kita lebih modular, scallable, dan maintainable.
* Modular dalam artian kita bisa dengan mudah mengganti dipedensi satu ke dipendensi lain.
* Scallable dalam artian kita dapat dengan mudah menambahkan feature baru dan lain sebagainya.
* Mainainable dalam artian kita dapat dengan mudah memperbaiki issue bilamana terdapat issue pada kode kita.

** Opsi migrasi
* pertahankan desain sekarang dengan memisahkan dependensi.
* Pertahankan desainnya tetapi kita pindahkan kodenya kedalam suatu layer.
* Ubah desainnya dan pisahkan dependensi.

** Struktur code
* Controller : berisi kode yang berhubungan langsung ke user (interface layer).
* Repository : berisi kode yang berhubungan langsung ke database (interface layer).
* Usecase : berisi bisnis logik yang dipakai.