** System design, sebelum membangun sebuah sistem baiknya kita melakukan design terkait system yang akan kita kerjakan.

** Diagram, sebuah symbol yang mempresentasikan informasi design system yang akan disampaikan. tools yang dapat digunakan adalah draw.io, lucidchart,smartdraw dll.

** System design yang biasa digunakan adalah Flowchart, Use Case, ERD dan HLA.

** Flowchart, representasi sebuah diagram yang mempresentasikan sebuah alur. 

** Use Case diagram merupakan sebuah ringkasan dari sistem dan apa saja yang terdapat pada applikasi kita dan interaksi dari sistem kepada aktor/user. include dimaksudkan bahwa proses sebelumnya harus dikerjakan terlebih dahulu, sedangkan extends merupakan pilihan tambahan yang dapat dilakukan.

** ERD merupakan sebuah tabel yang memiliki entity apa saja yang dibutuhkan oleh applikasi dalam tabel database, lalu pada design ERD kita juga menentukan hubungan tiap tabel (1-1, 1-n, m-n).

** HLA (High Level Architecture), dibuat ketika kita sudah ada pandangan system design secara umum dan telah tau target dari applikasi. HLA merupakan gambaran paling tinggi ketika kita telah mengetahui teknologi dan bahasa pemrograman apa dalam sistem kita.

** ketika sistem masih kecil dan memiliki transaksi data masih dibawah 100, untuk transaksi yang besar kita harus memikirkan 3 aspek, 1. bagian terkecil apa yang dapat digunakan, 2. bagian tersebut akan bagaimana ketika telah banyak transaksi, 3. bagaimana optimasinya.

*** Karakteristik sistem terdistribusi :

*** Scalability (penting bagi sistem sederhana bagaimana sistem dapat berjalan dengan baik) dapat dilakukan dengan cara : ** Scaling yakni meningkatkan kebutuhan sistem sesuai dengan keadaan yang terjadi pada sistem, apabila server telah penuh maka dapat menyewa server lain untuk memenuhi kebutuhan request pada sistem (duplikasis server).
* Horizontal Scaling(scalling dengan cara meningkatkan angka/ability pada suatu server sesuai kebutuhan), dipilih ketika kita tidak dapat memprediksi dan ada kemungkinan akan naik dengan cepat serta ada kemungkinan akan dilakukan perubahan code untuk beradaptasi. (ketika butuh cepat dan besar)
* Vertical Scaling (melakukan duplikasi server untuk meningkatkan ability server), dipilih ketika memiliki kemungkinan untuk meningkatkan ability dan akan digunakan dalam jangka waktu yang lama serta tidak dilakukan perubahan code. (ketika memiliki waktu untuk analisis yang lama)

** Reliability, jangan sampai sistem gagal memberikan layanan bagi user maka dapat dilakukan dengan cara multi server agar sistem dapat selalu tersedia untuk memenuhi request user tanpa mengalami down.

** Availability, ketersedian dari sistem untuk diakses oleh user dan untuk sarannya dapat dilakukan multi server.

** Efficiency, untuk efisiensi sebuah sistem waktu tunggu maksimal adalah 2 detik untuk user menunggu response dari sistem.

** Serviceablity or Manageability, sistem harus dapat di maintenance sehingga dalam code harus sudah clean sehingga dapat lebih mudah dalam maintenance sistem.

*** Job/ Work Queue, ketika user melakukan pembayaran didalam sistem maka akan berjalan di work queue sehingga client tidak harus menunggu dan dapat melakukan kegiatan lain dalam sistem dan apabila sudah selesai maka user akan diberikan notifikasi.

*** Load Balancing, merupakan komponen penting yang membantu dalam pembagian trafic diantara server sehingga dapat menggunakan server secara maksimal. dapat diberikan ketika user mengakses Front End sehingga agar berjalan smooth dapat diberikan load balancer diantaranya, front end akan mengakses API sehingga ketika akan mengakses API ke server dapat diberikan load balancer, ketika mengakses database server juga dapat diberikan load balancer.

*** Monolithic dan Microservices
** Monolithic, ketika semua service pada project dijadikan 1 server. ketika transaksi tidak besar maka monolithic dapat menanganinya dengan baik. apabila kita akan melakukan upgred pada 1 service yang banyak digunakan oleh user maka semua service akan terupgrade dan membutuhkan cost lebih banyak.

** Microservices digunakna ketika transaksi dari user besar sehingga agar tidak terjadi overload / down maka digunakan microservices, sehingga service akan dipecah-pecah menjadi sesuai kebutuhannya. apabila terjadi banyak requesst pada salah satu service maka kita dapat hanya meningkatkan satu service itu saja tersebut. dan untuk komunikasinya menggunakan jalur REST atau menggunakan broker.

Sql dan nosql

Sql (relational database/terstruktur) no sql (tabel tidak memiliki hubungan/tak terstruktur)

Manfaat sql
- dirancang untuk segala keperluan
- memiliki standar yang jelas
- memiliki banyak tool

4 prinsip didalam sql
- atomicity, transaksk terjadi semua atau tidak sama sekali
- consistency, data tertulis merupakan data valid ditentukan berdasarkan aturan tertentu
- Isolation, pada saat terjadi request yang bersamaan maka memastikan bahwa transaksi dieksekusi secara sekuensial
- Durability, transaksi yg tersimpan akan tetap tersimpan


No Sql, merupakan pandangan baru terkait database. DBMS yang menyediakan mekanisme yg lebih fleksibel dibandingkan dengan model RDBMS

No sql digunakan untuk menghindari :
- effort pada sifat transaksi ACID
- kompleksitas SQL
- design schema didepan
- transaction ditangani oleh aplikasi

Kelebihan no sql adalah schema less, fast development, etc.

No sql digunakan ketika membutuhkan skema fleksibel, ACID tidak diperlukan, data loging (json), data sementara (cache)

No sql tidak dapat digunakan pada data finansial, data transaksi, business critical, ACID - compliant.

Caching, merupakan penyimpanan data sementara yang mana akan menyimpan data req sehingga ketika menerima req yang sama maka tidak perlu mengakses db lagi. Cache akan disimpan didalam Ram.

Database replication
Redudancy and replication (replication database back up yang mana saling terhubung dengan database utama untuk mengantisipasi kerusakan pada database utama sehingga dapat langsung menggunakan database replication)

Database indexing, merupakan cara untuk mengoptimalkan performa dsri database dengan meminimalisir banyaknya akses kedalam disk keika melakukan query.