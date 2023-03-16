-- 1a
insert into operators 
(name, created_at, updated_at) values
('Telkomsel', now(), now()), 
('Tri', now(), now()), 
('Axis', now(), now()), 
('Smartfren', now(), now()),
('XL', now(), now());

-- 1b
insert into product_types 
(name, created_at, updated_at) values
('Paket Malam', now(), now()), 
('Paket Unlimited', now(), now()),
('Paket Sosial Media', now(), now());

-- 1c
insert into products
(product_type_id, operator_id, product_description_id, code, name, quantity, price, created_at, updated_at) values
(1, 3, 1, 'SI21SSXSQWE23', 'Paket OMG', 1, 40000, now(), now()), 
(1, 3, 1, 'KSOUW86HSYW545', 'Paket Berkah', 1, 20000, now(), now());

-- 1d
insert into products
(product_type_id, operator_id, product_description_id, code, name, quantity, price, created_at, updated_at) values
(2, 1, 1, 'JDSIAJ7665ASDUH', 'Paket Berlimpah', 1, 30000, now(), now()), 
(2, 1, 1, 'IIJ65ASDUH5677H', 'Paket Seru', 1, 20000, now(), now()), 
(2, 1, 1, 'OJDSA878JADS767', 'Paket Sempurna', 1, 10000, now(), now());

-- 1e
insert into products
(product_type_id, operator_id, product_description_id, code, name, quantity, price, created_at, updated_at) values
(3, 4, 1, 'OSAIJ876HUDSA57', 'Paket Maknyus', 1, 30000, now(), now()), 
(3, 4, 1, 'IJDSAUH767HSAU5', 'Paket Berharga', 1, 20000, now(), now()), 
(3, 4, 1, 'DUSAHU675UHSADU5', 'Paket Diam', 1, 40000, now(), now());

-- 1f
insert into product_descriptions 
(description, created_at, updated_at) values
('Paket OMG gratis telpon dan internetan', now(), now()), 
('Dengan paket berkah hidup akan semakin berkah', now(), now()), 
('Paket berlimpah kuota internet', now(), now()), 
('Seru seruan dengan kuoata dan paket telepon', now(), now()), 
('Sempurnakan hidupmu dengan paket sempurna', now(), now()), 
('Maknyus banget dengan paket telpon dan internet', now(), now()), 
('Paket berharga dapat menemai hari - hari mu yang indah', now(), now()), 
('Ssstttt... diam-diam kalo internetan', now(), now());

-- 1g
insert into payment_methods 
(method, created_at, updated_at) values 
('OVO', now(), now()), 
('Dana', now(), now()), 
('Gopay', now(), now());

-- 1h
insert into users
(name, address, dob, status, gender, created_at, updated_at) values
('Adi', 'Jalan Amburadul 1', '2000-01-01', 'Active', 'M', now(), now()),
('Budi', 'Jalan Amburadul 2', '2000-02-02', 'Active', 'M', now(), now()),
('Citra', 'Jalan Amburadul 3', '2000-03-03', 'Active', 'F', now(), now()),
('Dina', 'Jalan Amburadul 4', '2000-04-04', 'Active', 'F', now(), now()),
('Eka', 'Jalan Amburadul 5', '2000-05-05', 'Active', 'M', now(), now());

-- 1i
insert into transactions 
(method_id, user_id, code, quantity, total_price, created_at, updated_at) values
(1, 1, 'AETSRUUDFYIGO', 1, 30000, now(), now()),
(2, 1, 'WERUYOUPJLKHI', 2, 70000, now(), now()),
(3, 1, 'AEYTSRYTUYIUL', 3, 70000, now(), now()),
(1, 2, 'PIYOUILYUYHTG', 4, 60000, now(), now()),
(2, 2, 'POIUTYHDRPOII', 5, 130000, now(), now()),
(3, 2, 'LKNJHVGGFSTFU', 1, 20000, now(), now()),
(1, 3, 'RTYUIFDHGFHHJ', 2, 80000, now(), now()),
(2, 3, 'AYESTRHDJHGYU', 3, 50000, now(), now()),
(3, 3, 'BVHJHKUJTHRTE', 4, 80000, now(), now()),
(1, 4, 'MBHRYETRUFYUG', 5, 200000, now(), now()),
(2, 4, 'LJKJEYRTUFYIU', 1, 110000, now(), now()),
(3, 4, 'CFYGUIHOIJOFI', 2, 80000, now(), now()),
(1, 5, 'MJTYETRWUTFYG', 3, 70000, now(), now()),
(2, 5, 'RJTUYUUOIOUHO', 4, 120000, now(), now()),
(3, 5, 'SI21SSXSQWE23', 5, 80000, now(), now());

-- 1j
insert into transaction_details 
(transaction_id, product_id, created_at, updated_at) values
(1, 1, now(), now()),
(1, 2, now(), now()),
(1, 3, now(), now()),
(2, 4, now(), now()),
(2, 5, now(), now()),
(2, 6, now(), now()),
(3, 7, now(), now()),
(3, 8, now(), now()),
(3, 1, now(), now()),
(4, 2, now(), now()),
(4, 3, now(), now()),
(4, 4, now(), now()),
(5, 5, now(), now()),
(5, 6, now(), now()),
(5, 7, now(), now()),
(6, 8, now(), now()),
(6, 1, now(), now()),
(6, 2, now(), now()),
(7, 3, now(), now()),
(7, 4, now(), now()),
(7, 5, now(), now()),
(8, 6, now(), now()),
(8, 7, now(), now()),
(8, 8, now(), now()),
(9, 1, now(), now()),
(9, 2, now(), now()),
(9, 3, now(), now()),
(10, 4, now(), now()),
(10, 5, now(), now()),
(10, 6, now(), now()),
(11, 7, now(), now()),
(11, 8, now(), now()),
(11, 1, now(), now()),
(12, 2, now(), now()),
(12, 3, now(), now()),
(12, 4, now(), now()),
(13, 5, now(), now()),
(13, 6, now(), now()),
(13, 7, now(), now()),
(14, 8, now(), now()),
(14, 1, now(), now()),
(14, 2, now(), now()),
(15, 3, now(), now()),
(15, 4, now(), now()),
(15, 5, now(), now());