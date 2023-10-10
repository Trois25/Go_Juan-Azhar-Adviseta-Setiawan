
Link Success deploy
https://github.com/Trois25/latihan_ci-cd_Golang_Altera

*  CI/CD
** Continuous Integration(CI) merupakan proses otomasi yang dilakukan dengan tujuan integrasi code dari sources yang berbeda dengan tujuan melakukan build atau melakukan test.

** Cycle dare cicd adalah melakuakn check pada perubahan -> fetch perubahan -> build -> test -> fail or success -> integrasikan ke server untuk mengirimkan notifikasi fail or succes.

** Continuous delivery/deployment, proses ini dilakukan selangkah lebih maju dari integration dan mengirimkan proses, alurnya akan otomatis melakukan deploy setiap melakukan build yang telah diverifikasi.

- buat folder sudah ketentuan ".github\workflows"
- tambahkan file nama_file.yml
- isinya copy dari mas fakhry :
-- 3 - 6 = ketika action bekerja saat push ke branch main
-- 9 - 20 docker itu melakukan install
-- 22 - 26 login dockerhub

- command secrets.nama_variable, diatur di setting github lalu Environtmen dari project. didalamnya tambahlah sesuai nama variable pada secrets yang didalam code
- membuat token docker hub di account setting, lalu security dan tambah new access token

-- 28 - 34 melakukan build dan push menggunakan Dockerfile dan di push ke repo yang telah dibuat

-- 40 - 46 melakukan connection ke SSH
-- secrete.host = ip public vm
-- secrete.username dari ssh vm
-- secrete.key = id_ecdsa

-- 48 : docker stop
-- 49 : menghapus container
-- 50 : menghapus image
-- 51 : pull dari docker hub
-- 52 : run image  yang diambil dari dockerhub