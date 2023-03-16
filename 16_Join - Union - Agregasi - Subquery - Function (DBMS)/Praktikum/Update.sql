-- 3a
update products set name = 'product dummy' where product_id = 1

-- 3b
-- Tidak ada qty pada tabel transaction_details yang saya buat jadi saya alihkan ke transaction
update transactions set quantity = 3 where transaction_id = 1