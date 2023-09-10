-- DML

-- 1. insert

-- a. insert 5 operators pada table operators
INSERT INTO operators (id, name, created_at, updated_at)
VALUES
( 1, 'Decade',CURRENT_TIME(),CURRENT_TIME() ),
( 2, 'Build',CURRENT_TIME(),CURRENT_TIME() ),
( 3, 'Ryuki',CURRENT_TIME(),CURRENT_TIME() ),
( 4, 'Faiz',CURRENT_TIME(),CURRENT_TIME() ),
( 5, 'Gaim',CURRENT_TIME(),CURRENT_TIME() );

-- b. insert 3 product type 
INSERT INTO product_types (id, name, created_at, updated_at)
VALUES
( 1, 'Hobi',CURRENT_TIME(),CURRENT_TIME() ),
( 2, 'Pakaian',CURRENT_TIME(),CURRENT_TIME() ),
( 3, 'Smarthphone',CURRENT_TIME(),CURRENT_TIME() );

-- c. insert 2 product dengan product type id = 1, dan operators id = 3
INSERT INTO products (id, product_type_id, operator_id, code, name, status, created_at, updated_at)
VALUES
( 1, 1, 3, "GUNPLA","Gundam Aerial HG 1/144",400,CURRENT_TIME(),CURRENT_TIME() ),
( 2, 1, 3, "FIGURE","Nekomata Okayu Relax Time Figure",400,CURRENT_TIME(),CURRENT_TIME() );

-- d. insert 3 product dengan product type id = 2, dan operators id = 1
INSERT INTO products (id, product_type_id, operator_id, code, name, status, created_at, updated_at)
VALUES
( 3, 2, 1, "HODIE","Hodie RRQ",400,CURRENT_TIME(),CURRENT_TIME() ),
( 4, 2, 1, "KEMEJA","Kemeja hitam pria",400,CURRENT_TIME(),CURRENT_TIME() ),
( 5, 2, 1, "Batik","Batik Jambi motif Angso Duo",400,CURRENT_TIME(),CURRENT_TIME() );

-- e. Insert 3 product dengan product type id = 3, dan operators id = 4
INSERT INTO products (id, product_type_id, operator_id, code, name, status, created_at, updated_at)
VALUES
( 6, 3, 4, "TWS","TWS Neko",400,CURRENT_TIME(),CURRENT_TIME() ),
( 7, 3, 4, "EARPHONE","Earphone Jack 10D",400,CURRENT_TIME(),CURRENT_TIME() ),
( 8, 3, 4, "CASE","Custome case",400,CURRENT_TIME(),CURRENT_TIME() );

-- f. Insert product description pada setiap product
INSERT INTO products_descriptions (id, description, created_at, updated_at)
VALUES
( 1, "GUNPLA Aerial dengan ukuran 1/144 ori BANDAI",CURRENT_TIME(),CURRENT_TIME() ),
( 2, "Action Figure kolaborasi pemandian air panas dan member HOLOLIVE",CURRENT_TIME(),CURRENT_TIME() ),
( 3, "Hodie original RRQ ",CURRENT_TIME(),CURRENT_TIME() ),
( 4, "Kemeja Hitam polos pria untuk segala umur",CURRENT_TIME(),CURRENT_TIME() ),
( 5, "Batik khas jambi dengan motif dua angsa",CURRENT_TIME(),CURRENT_TIME() ),
( 6, "Tws dengan bluetoth 5g dan jarak 10m",CURRENT_TIME(),CURRENT_TIME() ),
( 7, "Earphone dengan jack audio yang sangat baik dan tahan lama",CURRENT_TIME(),CURRENT_TIME() ),
( 8, "Custome case dapat dilakukan dengan mengirim gambar ke EMAIL dan melakukan konfirmasi dengan admin",CURRENT_TIME(),CURRENT_TIME() );

-- g. Insert 3 payment methods
INSERT INTO payment_methods (id, name, status, created_at, updated_at)
VALUES
( 1, "Bank BSI", 400, CURRENT_TIME(),CURRENT_TIME() ),
( 2, "DANA", 400, CURRENT_TIME(),CURRENT_TIME() ),
( 3, "Shopeepay", 400, CURRENT_TIME(),CURRENT_TIME() );

-- h. Insert 5 user pada tabel user
INSERT INTO users (id, name, status, dob, gender, created_at, updated_at)
VALUES
( 1, 'juan', 400, '2003-01-25', 'M', CURRENT_TIME(),CURRENT_TIME() ),
( 2, 'fathir', 400, '2007-11-09', 'M', CURRENT_TIME(),CURRENT_TIME() ),
( 3, 'setiawan', 400, '1982-05-09', 'M', CURRENT_TIME(),CURRENT_TIME() ),
( 4, 'sari', 400, '1980-11-06', 'W',CURRENT_TIME(),CURRENT_TIME() ),
( 5, 'tiara', 400, '2002-07-16', 'W',CURRENT_TIME(),CURRENT_TIME() );

-- i. Insert 3 transaksi di masing-masing user. (soal berlanjut ke soal 1.j)
INSERT INTO transactions (id, user_id, payment_method_id, status, total_qty, total_price, created_at, updated_at)
VALUES
(1, 1, 1, 400, 2, 500000, CURRENT_TIME(),CURRENT_TIME() ),
(2, 1, 1, 400, 1, 400000, CURRENT_TIME(),CURRENT_TIME() ),
(3, 1, 1, 400, 5, 100000, CURRENT_TIME(),CURRENT_TIME() ),
(4, 2, 2, 400, 1, 250000, CURRENT_TIME(),CURRENT_TIME() ),
(5, 2, 2, 400, 1, 400000, CURRENT_TIME(),CURRENT_TIME() ),
(6, 2, 2, 400, 1, 75000, CURRENT_TIME(),CURRENT_TIME() ),
(7, 3, 3, 400, 5, 500000, CURRENT_TIME(),CURRENT_TIME() ),
(8, 3, 3, 400, 4, 490000, CURRENT_TIME(),CURRENT_TIME() ),
(9, 3, 3, 400, 10, 650000, CURRENT_TIME(),CURRENT_TIME() ),
(10, 1, 4, 400, 2, 750000, CURRENT_TIME(),CURRENT_TIME() ),
(11, 1, 4, 400, 3, 70000, CURRENT_TIME(),CURRENT_TIME() ),
(12, 1, 4, 400, 4, 100000, CURRENT_TIME(),CURRENT_TIME() ),
(13, 3, 1, 400, 6, 200000, CURRENT_TIME(),CURRENT_TIME() ),
(14, 3, 1, 400, 4, 100000, CURRENT_TIME(),CURRENT_TIME() ),
(15, 3, 1, 400, 2, 50000, CURRENT_TIME(),CURRENT_TIME() );

-- j. Insert 3 product di masing-masing transaksi


-- 2. Select

-- a. Tampilkan nama user / pelanggan dengan gender Laki-laki / M.
SELECT name FROM users WHERE gender = 'M';

-- b. Tampilkan product dengan id = 3.
SELECT * FROM product WHERE id = 3;

-- c. Tampilkan data pelanggan yang created_at dalam range 7 hari kebelakang dan mempunyai nama mengandung kata ‘a’.
SELECT * FROM operators WHERE created_at < now() && name ='%a%';

-- d. Hitung jumlah user / pelanggan dengan status gender Perempuan.
SELECT COUNT(*) FROM operators WHERE gender = 'W';

-- e. Tampilkan data pelanggan dengan urutan sesuai nama abjad
SELECT * FROM operators ORDER BY name;

-- f. Tampilkan 5 data pada data product
SELECT * FROM product LIMIT 5;

-- 3. Update

-- a.Ubah data product id 1 dengan nama ‘product dummy’
UPDATE product SET name = 'product dummy' WHERE product_id = 1;
-- b. Update qty = 3 pada transaction detail dengan product id = 1
UPDATE transaction_detail SET qty = 3 WHERE product_id = 1;

-- 4. Delete

-- a. Delete data pada tabel product dengan id = 1
DELETE FROM product WHERE id = 1;

-- b. Delete data pada tabel product dengan product type id 1
DELETE FROM product WHERE product_type_id = 1;

-- Join, Union, Sub quert, Function

-- 1. Gabungkan data transaksi dari user id 1 dan user id 2
SELECT * FROM transactions WHERE user_id = 1
UNION ALL
SELECT * FROM transactions WHERE user_id = 2;

-- 2 Tampilkan jumlah harga transaksi user id 1
SELECT user_id, SUM(total_price) AS total_harga_transaksi
FROM transactions WHERE user_id = 1;

-- 3. Tampilkan total transaksi dengan product type 2
SELECT SUM(t.total_price) AS total_harga_transaksi_type_2
FROM transactions t
INNER JOIN products p ON t.product_id = p.id
INNER JOIN product_types pt ON p.product_type_id = pt.id
WHERE pt.id = 2;

-- 4. Tampilkan semua field table product dan field name table product type yang saling berhubungan
SELECT p.*, pt.name AS product_type_name
FROM products p
INNER JOIN product_types pt ON p.product_type_id = pt.id;

-- 5. Tampilkan semua field table transaction, field name table product dan field name table user
SELECT t.*, p.name AS product_name, u.name AS user_name
FROM transactions t
INNER JOIN products p ON t.product_id = p.id
INNER JOIN users u ON t.user_id = u.id;

-- 6. Buat function setelah data transaksi dihapus maka transaction detail terhapus juga dengan transaction id yang dimaksud
DELIMITER //

CREATE FUNCTION DeleteTransactionAndDetails(transaction_id INT) RETURNS INT
BEGIN
  DECLARE rows_affected INT;

  START TRANSACTION;

  -- Hapus data dari tabel transaction_details yang sesuai dengan transaction_id
  DELETE FROM transaction_details WHERE transaction_id = transaction_id;

  -- Hapus data dari tabel transactions dengan transaction_id yang sama
  DELETE FROM transactions WHERE id = transaction_id;

  -- Mendapatkan jumlah baris yang terpengaruh oleh penghapusan
  SET rows_affected = ROW_COUNT();

  COMMIT;

  RETURN rows_affected;
END//

DELIMITER ;

-- 7. Buat function setelah data transaksi detail dihapus maka data total_qty terupdate berdasarkan qty data transaction id yang dihapus
DELIMITER //

CREATE FUNCTION UpdateTotalQty(transaction_id INT) RETURNS INT
BEGIN
  DECLARE total_qty INT;

  -- Menghitung total_qty baru berdasarkan transaction_id
  SELECT SUM(qty) INTO total_qty
  FROM transaction_details
  WHERE transaction_id = transaction_id;

  -- Mengupdate total_qty di tabel transactions
  UPDATE transactions
  SET total_qty = total_qty
  WHERE id = transaction_id;

  RETURN total_qty;
END//

DELIMITER ;

-- 8. Tampilkan data products yang tidak pernah ada di tabel transaction_details dengan sub-query
SELECT p.*
FROM products p
WHERE p.id NOT IN (
  SELECT DISTINCT td.product_id
  FROM transaction_details td
);
