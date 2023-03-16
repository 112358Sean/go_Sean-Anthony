-- 2a
select * from users where gender = 'M'

-- 2b
select * from products where product_id = 3

-- 2c
select * from users where name like '%a%' and created_at > date(now() - interval 7 day)

-- 2d
select count(*) from users where gender = 'F'

-- 2e
select * from users order by name

-- 2f
select * from products limit 5