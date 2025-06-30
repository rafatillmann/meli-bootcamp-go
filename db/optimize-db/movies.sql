drop table twd;

create temporary table twd as
select e.* from episodes e
join seasons s on e.season_id = s.id
join series se on s.serie_id = se.id
where se.title = 'The Walking Dead';

select t.* from twd t;

select t.* from twd t
join seasons s on t.season_id = s.id
where s.number = 1;  

select * from seasons s;

create index in_number
on seasons (number);

show index from seasons;

-- --------

select * from movies m;

insert into movies (title, rating, awards, release_date, length, genre_id) values ('Brother Bear', 10, 1, '2000-01-01', 120, 7);

select * from genres g;

insert into genres (name, ranking, active) values ('Test', 13, 1)

update movies m set genre_id = 13 where m.title = 'Brother Bear'

select * from actors a;

update actors a set favorite_movie_id = 22 where a.id = 3;

create temporary table t_movies as
select m.* from movies m;

select * from t_movies t;

delete from t_movies t where awards < 5;

select distinct g.* from genres g
join movies m on g.id = m.genre_id;

select a.* from actors a
join movies m on a.favorite_movie_id = m.id
where m.awards > 3;

create index in_title_movie
on movies (title);

show index from movies;

create index in_release_date_serie
on series (release_date);

show index from series;