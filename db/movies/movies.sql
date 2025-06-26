select * from movies m;

select a.first_name, a.last_name, a.rating from actors a;

select s.title as `título` from series s;

select a.first_name, a.last_name from actors a where a.rating > 7.5;

select m.title, m.rating, m.awards from movies m where m.rating > 7.5 and m.awards > 2;

select m.title, m.rating from movies m order by m.rating;

select m.title from movies m limit 3;

select m.title, m.rating from movies m order by m.rating desc limit 5;

select * from actors a limit 10;

select m.title, m.rating from movies m where m.title like 'Toy Story%';

select * from actors a where a.first_name like 'Sam%'

select m.title, m.release_date from movies m where m.release_date between '2004-01-01' and '2008-12-31';

select m.title from movies m where m.rating > 3 and m.awards > 1 and m.release_date between '1998-01-01' and '2009-12-31';

select m.title, m.rating from movies m where m.rating > 3 and m.awards > 1 and m.release_date between '1998-01-01' and '2009-12-31' order by m.rating;