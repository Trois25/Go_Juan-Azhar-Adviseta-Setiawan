--Data Definition Language (DDL)

-- 1. create database alta_online_shop
CREATE DATABASE alta_online_shop;

-- 2. membuat tabel dari schema Olshop
CREATE TABLE `user`(
    user_id int,
    username varchar(25),
    password varchar(25),
    user_data_id int,
    PRIMARY KEY(user_id),
    FOREIGN KEY (user_data_id) REFERENCES user_data(user_data_id)
)

CREATE TABLE user_data(
    user_data_id int,
    nama varchar(25),
    jenis_kelamin varchar (15),
    alamat varchar (100),
    PRIMARY KEY (user_data_id)
)

CREATE TABLE product(
    product_id int,
    nama_product varchar(50),
    harga int,
    product_type_id int,
    PRIMARY KEY(product_id),
    FOREIGN KEY (product_type_id) REFERENCES product_type(product_type_id)
)

CREATE TABLE product_type(
    product_type_id int,
    type varchar(25),
    PRIMARY KEY(product_type_id)
)

CREATE TABLE transaction(
    transaction_id int,
    product_id int,
    PRIMARY KEY(transaction_id),
    FOREIGN KEY (product_id) REFERENCES product(product_id)
)

CREATE TABLE operators(
    operators_id int,

    PRIMARY KEY(operators_id)
)

CREATE TABLE payment_method(
    payment_method_id int,
    pembayaran varchar(50),
    product_description_id int,
    PRIMARY KEY(payment_method_id),
    FOREIGN KEY (product_description_id) REFERENCES product_description(product_description_id)
)

CREATE TABLE product_description(
    product_description_id int,
    description varchar(100),
    product_id int,
    PRIMARY KEY(product_description_id),
    FOREIGN KEY (product_id) REFERENCES product(product_id)
)

CREATE TABLE transaction_detail(
    transaction_detail_id int,
    user_id int,
    product_id int,
    total_harga int,
    PRIMARY KEY(transaction_detail_id),
    FOREIGN KEY (product_id) REFERENCES product(product_id),
    FOREIGN KEY (user_id) REFERENCES user(user_id)
)


-- 3. Create table kurir
CREATE TABLE kurir(
    kurir_id int,
    name varchar(50),
    created_at datetime,
    updated_at datetime,
    PRIMARY KEY(kurir_id)
)

-- 4. tambahkan ongkos_dasar ketabel kurir
ALTER TABLE kurir ADD ongkos_dasar int;

-- 5. rename tabel kurir menjadi shipping
ALTER TABLE kurir RENAME TO shipping;

-- 6. drop tabel shipping
DROP TABLE shipping;

-- 7. menambahkan entity relation 
--a. 1 - 1 payment method description
ALTER TABLE payment_method ADD FOREIGN KEY (product_description_id) REFERENCES product_description(product_description_id);

--b. 1 - m user dan alamat
ALTER TABLE `user` ADD FOREIGN KEY (user_data_id) REFERENCES user_data(user_data_id);

--c. m - m user dengan payment method
CREATE TABLE user_payment_method_detail(
user_payment_method_detail_id int,
user_id int,
payment_method_id int,
PRIMARY KEY(user_payment_method_detail_id),
FOREIGN KEY (user_id) REFERENCES `user`(user_id),
FOREIGN KEY (payment_method_id) REFERENCES payment_method(payment_method_id)
)