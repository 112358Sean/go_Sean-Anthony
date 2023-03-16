-- 1
create database alta_online_shop;

-- 2a
create table users(
	user_id smallint primary key auto_increment,
	name varchar(255),
	address varchar(255),
	dob date,
	status varchar(255),
	gender varchar(255),
	created_at timestamp,
	updated_at timestamp
);

-- 2b
create table product_types(
	product_type_id smallint primary key auto_increment,
	name varchar(255),
	created_at timestamp,
	updated_at timestamp
);

create table operators(
	operator_id smallint primary key auto_increment,
	name varchar(255),
	created_at timestamp,
	updated_at timestamp
);

create table product_descriptions(
	product_description_id smallint primary key auto_increment,
	description text,
	created_at timestamp,
	updated_at timestamp
);

create table products(
	product_id smallint primary key auto_increment,
	product_type_id smallint references product_types,
	operator_id smallint references operator,
	product_description_id smallint references product_description,
	code varchar(255),
	name varchar(255),
	quantity int,
	price int,
	created_at timestamp,
	updated_at timestamp
);

create table payment_methods(
	payment_method_id smallint primary key auto_increment,
	method varchar(255),
	created_at timestamp,
	updated_at timestamp
);

-- 2c
create table transactions(
	transaction_id smallint primary key auto_increment,
	method_id smallint references payment_methods,
	user_id smallint references users,
	code varchar(255),
	quantity int,
	total_price int,
	created_at timestamp,
	updated_at timestamp
);

create table transaction_details(
	transaction_detail_id smallint primary key auto_increment,
	transaction_id smallint references transactions,
	product_id smallint references users,
	created_at timestamp,
	updated_at timestamp
);

-- 3
create table kurir(
	kurir_id smallint primary key auto_increment,
	name varchar(255),
	created_at timestamp,
	updated_at timestamp
);

-- 4
alter table kurir add ongkos_dasar INT;

-- 5
rename table kurir to shipping;

-- 6
drop table shipping;

-- 7
-- Telah dilakukan pada soal nomor 2