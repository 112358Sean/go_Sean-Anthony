-- 1
select * from transactions where transaction_id = 1
union 
select * from transactions where transaction_id = 2

-- 2
select sum(total_price) from transactions where user_id = 1

-- 3
select count(*) from transactions t 
join transaction_details td on t.transaction_id = td.transaction_id 
where td.product_id = 2;

-- 4
select p.*, pt.name from products p join product_types pt on p.product_id = pt.product_type_id;

-- 5
select t.*, p.name, u.name from transactions t 
join transaction_details td on t.transaction_id = td.transaction_detail_id 
join products p on p.product_id = td.transaction_detail_id 
join users u on u.user_id = t.user_id;

-- 6
DELIMITER $$
create function deleteTransaction (p_transaction_id int)
returns int deterministic
begin
delete from transactions where transaction_id = p_transaction_id;
return 1;
end$$
DELIMITER

DELIMITER $$
create trigger deleteTrigger
before delete on transactions for each row 
begin
declare trans_id int;
set trans_id = old.transaction_id;
delete from transaction_details where transaction_detail_id = trans_id;
end$$
DELIMITER

-- 7
DELIMITER $$
create function deleteTransDetail(p_transaction_id int)
returns int deterministic
begin
delete from transaction_details where transaction_id = p_transaction_id;
return 1;
end$$
DELIMITER

DELIMITER $$
create trigger updateQty
after delete on transaction_details for each row
begin
declare total_qty int;
declare total_qty2 int;
set total_qty = (select total_qty from transactions where transaction_id = transaction_detail_id);
update transaction SET total_qty = total_qty - quantity where transaction_id = transaction_detail_id;
end$$
DELIMITER

-- 8
select * from products where product_id not in (select product_id from transaction_details);