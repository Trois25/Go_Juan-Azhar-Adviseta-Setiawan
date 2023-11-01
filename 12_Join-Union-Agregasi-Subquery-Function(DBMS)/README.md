** Database, merupakan sekumpulan data yang terorganisir.

* Model, tabel yang akan menyimpan data didalam database, relationship antar tabel ada One To One (1 user hanya memiliki 1 foto profile) , One To Many (1 user dapat melakukan banyka tweets) dan Many to Many (1 mahasiswa bisa memiliki banyak matkul dan 1 matkul bisa diambil banyak mahasiswa).

* Untuk mengimplementasikan hubungan diatas dapat digunakan Relational Database Management System (RDBMS), contoh software yang menggunakannya seperti MySQL.

* Perintah SQL dibagi menjadi Data Definition Language (DDL), Data Manipulation Language (DML) dana Data Control Language (DCL).

* DDL Statement -> CREATE DATABASE database_namel; USE database_name; CREATE TABLE table_name, DROP TABLE table_name, RENAME TABLE table_name.

* Membuat table menggunakan IS schema, CREATE TABLE table_name ( column 1 data_type PRIMARY KEY,column 2 data_type FOREIGN KEY, ... columnt data_type, PRIMARY KEY(one or more columns) ).

* Modify table schema dengan menggunakan ALTER TABLE table_name dan ADD COLUMN column_name data_type;

* Tipe data MySQL, Num, Huruf, Date.

* DML merupakan perintah yang digunakan untuk memanipulasi data dalam tabel dari suatu database, statement operation pada DML antara lain INSERT SELECT UPDATE dan DELETE.

* INSER EX : INSERT INTO Role(role_name,status) VALUES ('admin', 'active').

* SELECT EX : SELECT * FROM Role, SELECT role_name FROM ROLE WHERE status = 'active'

* UPDATE EX : UPDATE Role SET role_name = 'member' WHERE id = 1

* DELETE EX : DELETE FROM Role WHERE id = 1

* LIKE EX : SELECT user_id,message FROM tweets WHERE message LIKE 'H%'

* BETWEEN EX : SELECT user_id,message FROM tweets WHERE user_id BETWEEN 1 AND 5

* AND/OR EX : SELECT user_id,message FROM tweets WHERE message LIKE 'H%' OR user_id BETWEEN 1 AND 5

* AND/OR EX : SELECT user_id,message FROM tweets WHERE message LIKE 'H%' AND user_id BETWEEN 1 AND 5

* ORDER BY EX : SELECT user_id,message FROM tweets WHERE message LIKE 'H%' OR user_id BETWEEN 1 AND 5 ORDER BY id DESC (DESC dari terbesar ASC dari terkecil)

* LIMIT EX : SELECT user_id,message FROM tweets WHERE message LIKE 'H%' OR user_id BETWEEN 1 AND 5 ORDER BY id DESC LIMIT 2 (akan menampilkan hanya 2 data karena dibatasi limit 2)

* JOIN, menggabungkan 2 tabel yang mana dibagi menjadi 3 buah standar Inner Join (mengembalikan baris dari kedua tabel yang memiliki kesamaan), LEFT JOIN (mengembalikan seluruh baris dari tabel sebelah kiri yang dikenai kondisi ON dan hanya baris dari sebelah kanan yang memenuhi kondisi yang akan join), RIGHT JOIN (mengembalikan seluruh baris dari tabel sebelah kanan yang dikenai kondisi ON dan hanya baris dari sebelah kiri yang memenuhi kondisi yang akan join).

* INNER JOIN EX : SELECT t.message FROM users u INNER JOIN tweets t ON u.id = t.user_id;

* LEFT JOIN EX : SELECT u.username, t.message FROM users u LEFT JOIN tweets t ON u.id = t.user_id;

* RIGHT JOIN : SELECT u.username, t.message FROM users u RIGHT JOIN tweets t ON u.id = t.user_id;

* UNION digunakan untuk mengkombinasikan data dari 2 atau lebih hasil yang akan dijadikan menjadi 1 set hasil, syarat penggunaannya keduanya harus memiliki column yang sama.

* UNION EX : SELECT role_name FROM Role WHERE id = 1 UNION SELECT role_name FROM Role WHERE id = 2

* Fungsi Agregasi merupakan fungsi dimana nilai beberapa baris dikelompokkan bersama untuk membentuk nilai ringkasan tunggal (MIN, MAX, SUM, AVG, COUNT, HAVING)

* MIN (digunakan untuk mengeluarkan data terkecil) ex : SELECT MIN(id) AS is FROM users (akan menampilkan id terkecil dari tabel users)

* MAX (digunakan untuk mengeluarkan data terbesar) ex : SELECT MAX(id) FROM users (menampilkan nilai id terbesar dari tabel users)

* SUM (digunakan untuk mendapatkan jumlah total nilai dari sebuah data atau record di tabel) ex : SELECT SUM(fav_count) FROM tweets WHERE user_id = 1 (menampilkan jumlah total fav_count dari tweets id 1)

* AVG (digunakan untuk mencari nilai rata2 dari suatu data) ex : SELECT AVG(fav_count) FROM tweets WHERE user_id = 1 (menampilkan rata2 total fav_count dari tweets id 1)

* COUNT (digunakan untuk mencari jumlah dari sebuah data) ex : SELECT COUNT(1) FROM tweets WHERE user_id = 1 (menampilkan jumlah data dari tabel tweets dengan user_id 1)

* HAVING (digunakan untuk menyeleksi data berdasarkan kriteria tertentu, dimana kriteria berupa fungsi aggregat) ex : SELECT user_id FROM tweets GROUP BY user_id HAVING SUM(fav_count) > 2 (menampilkan data dari tabel tweets dengan total fav_count per user lebih dari 2)

* SUBQUERY / Inner query / Nested query adalah query lain dalam SQL yang digunakan untuk mengembalikan data yang akan digunakan dalam query utama sebagai syarat untuk lebih membatasi data yang akan diambil. Subquery dalam digunakan dengan SELECT, INSERT, UPDATE dan DELETE bersamaan dengan operator logical expression ( =, <, >, <=, >=, IN, BETWEEN, dll).

** peraturan SUBQUERY
* Harus tertutup dalam tanda kurung.
* ebuah subquery hanya dapat memiliki satu kolom pada klausa SELECT , kecuali beberapa kolom yang di query utama untuk subquery untuk membandingkan kolom yang dipilih.
* Subqueries yang kembali lebih dari satu baris hanya dapat digunakan dengan beberapa value operator , seperti operator IN .
* Daftar SELECT tidak bisa menyertakan referensi ke nilai - nilai yang mengevaluasi ke BLOB , ARRAY , CLOB , atau NCLOB .
* Sebuah subquery tidak dapat segera tertutup dalam fungsi set .

* Contoh : SELECT * FROM USERS WHERE id IN (SELECT user_id FROM tweets GROUP BY user_id); ( menampilkan data tabel users yang user_id nya ada pada tabel tweets)

* Function, sebuah kumpulan statement yang akan mengembalikan sebuah nilai balik pada pemanggilanya. Contoh : DELIMITER $$ CREATE FUNCTION sf_count_tweet_peruser(user_id_p int) RETURNS INT DETERMINISTIC BEGIN DECLARE total INT; SELECT COUNT (*) INTO total FROM tweets WHERE user_id - user_id_p AND type = 'tweets'; RETURN total; END$$ DELIMITER;
