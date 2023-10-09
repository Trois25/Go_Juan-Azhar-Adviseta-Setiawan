*** CI/CD
** Continuous Integration(CI) merupakan proses otomasi yang dilakukan dengan tujuan integrasi code dari sources yang berbeda dengan tujuan melakukan build atau melakukan test.

** Cycle dare cicd adalah melakuakn check pada perubahan -> fetch perubahan -> build -> test -> fail or success -> integrasikan ke server untuk mengirimkan notifikasi fail or succes.

** Continuous delivery/deployment, proses ini dilakukan selangkah lebih maju dari integration dan mengirimkan proses, alurnya akan otomatis melakukan deploy setiap melakukan build yang telah diverifikasi.

- buat folder sudah ketentuan ".github\workflows"
- tambahkan file nama_file.yml
-- 4 - 7 = ketika action bekerja saat push ke branch main
-- 10 - 22 docker itu melakukan install
-- 24 - 29 login dockerhub

- command secrets.nama_variable, diatur di setting github lalu Environtmen dari project. didalamnya tambahlah sesuai nama variable pada secrets yang didalam code
- membuat token docker hub di account setting, lalu security dan tambah new access token

-- 32 - 38 melakukan build dan push menggunakan Dockerfile dan di push ke repo yang telah dibuat

-- 45 - 51 melakukan connection ke SSH
-- secrete.host = ip public vm
-- secrete.username dari ssh vm
-- secrete.key = id_ecdsa

-- 53 : docker stop
-- 54 : menghapus container
-- 55 : menghapus image
-- 56 : pull dari docker hub
-- 57 : run image  yang diambil dari dockerhub