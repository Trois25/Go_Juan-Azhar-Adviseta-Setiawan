** Command line adalah interface berbasis text yang biasanya berada di server side dengan eksekusi command lebih cepat, kuat, effective dan effisien.

** Mengapa menggunakan CLI?
* Memiliki granular control dari OS atau applikasi
* lebih cepat dalam menangani banyak operasi
* kemampuannya dalam menyimpan script untuk melakukan otomasi regular tasks
* digunakan untuk membantu dalam troubleshootingseperti jaringan atau connection issues seperti port dan IP

** Unix Shell
* ~ => current path
* $ => shell normal user
* # => shell root user

** Normal user hanya diizinkan memanipulasi file/folder pada path /home/username, sedangkan root dapat memanipulasi semua file/folder. untuk menggunakan command root secara sementara dapat digunakan command sudo.

** Popular Commands

** Directory
* cd => dilakukan untuk berpindah path
* mkdir => membuat directory (mkdir nama_folder)
* ls => menampilkan list file
* rm => menghapus file (rm file_to_copy.txt, atau menghapus folder dengan rmdir nama_folder, untuk menghapus folder atau file dengan isi maka ditambahkan -r ct : rm -r nama_file.txt)
* cp => copy file (cp file_to_copy.txt new_file.txt)
* mv => move file (mv source_file destination_folder/)
* ln => link shortcut yang mana dibagi menjadi soft dan hard (soft dengan command : ln -s path/file . apabila dihapus maka softlink tak dapat dibuka, dan hard dengan command : 
* pwd => perintah untuk melihat working directory sekarang

** Files
* touch => create file (touch nama_file)
* cat => melihat isi file (cat nama_file)
* vim,nano => edit file (vim nama_file/nano nama_file)
* chmod => memberikan permission pada file (chmod 777 nama_file, 7 disini mempresentasikan rwe yang mana merupakan binner)
* diff => membandingkan file (diff file_1 file_2)

** Network
* ping => melihat konektivitas internet
* ssh => media transport untuk koneksi ke remote server
* netstat => mengetahui apa yang terjadi didalam jaringan kita
* nmap => digunakan untuk analisis port
* ifconfig => digunakan untuk melihat ip list 
* wget => mengunduh suatu file dari luar (ct : wget url_unduhan)
* curl => sama seperti wget
* telnet => sama seperti ssh untuk melakukan remote komputer namun secara keamanan lebih rendah
* netcat => digunakan untuk monitoring jaringan

**Utility
* man => informasi perintah apa yang diinginkan (ct : man ls)
* env => (menambahkan variabel kedalam env : export key=value)
* echo => untuk menampilkan (untuk menampilkan dan menambahkan kedalam file caranya : echo "Username:juan" > .env dan untuk melanjutkan mengisi file kebaris selanjutnya gunakan >>. echo "Password:123" >> .env)
* date => fungsi untuk menampilkan waktu sekarang
* which => mengetahui lokasi program file
* watch => memonitoring program yang berjalan
* sudo => mengakses perintah root user
* history => melihat command apa saja yang pernah dijalankan
* grep => mengambil suatu kata pada suatu file (ct : grep "kata" nama_file.txt)
* locate => melakukan pencarian 

** Shell Script (fungsi yang menjembatani user dan kernel yang merupakan programming language berdasarkan syntax)

** ketika akan mendeploy server maka akan melakukan pull docker, running container, testing, push registery, connect ssh kedalam cloud, push image dan pull image dari registry baru di running.