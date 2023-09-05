** Clean Code
* Clean code adalah istilah untuk kode yang mudah dibaca, dipahami dan diubah oleh programmer

* Clean Code baik diterapkan untuk kolaborasi antar programmer, feature development dan faster development.

** Karakteristik Clean Code : 
* Mudah dipahami, 
* mudah dieja dan dan dicari, 
* singkat namun mendeskripsikan konteks,
* konsisten, hindari penambahan konteks yang tidak perlu, 
* komentar pada fungsi penting pada code, 
* good function, 
* gunakan konvensi, 
* formatting, saran : lebar baris code 80 - 120 character, satu class 300 - 500 baris, baris code yang berhubungan saling berdekatan, dekatkan fungsi dengan peamnggilnya, deklarasi variabel berdekatan dengan penggunannya, perhatikan indentasi, menggunakan prettier/formatter.

** Prinsip clean code adalah KISS (Keep It So Simple), hindari membuat fungsi yang dibuat untuk melakukan A, sekaligus memodifikasi B, mengecek fungsi C, dst.

**Tips KISS
* Fungsi atau class harus kecil.
* Fungsi dibuat untuk melakukan satu tugas saja.
* Jangan gunakan terlalu banyak argumen pada fungsi.
* Harus diperhatikan untuk mencapai kondisi yang seimbang,kecil dan jumlahnya minimal.

**  DRY (Don't Repeat Yourself), duplikasi code terjadi karena sering copas dan untuk menghindarinya buatlah fungsi yang dapat digunakan secara berulang-ulang.

** Refaktoring, merupakan proses restruturisasi kode yang dibuat dengan mengubah struktur internal tanpa mengubah perilaku eksternal dan merupakan cara untuk mencapai prinsip KISS dan DRY.

** Teknik Refactoring
* Membuat sebuah abstraksi
* Memecah kode dengan fungsi/class
* Perbaiki penamaan dan lokasi kode
* Deteksi kode yang memiliki duplikasi