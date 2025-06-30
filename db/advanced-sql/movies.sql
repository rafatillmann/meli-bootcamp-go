select s.title, g.name from series s join genres g on s.genre_id = g.id;

select e.title, a.first_name, a.last_name from episodes e 
join actor_episode ae on ae.episode_id = e.id 
join actors a on a.id = ae.actor_id;

select s.title, count(se.id) from series s join seasons se on se.serie_id = s.id group by s.title;

select g.name, count(m.id) as total from genres g join movies m on m.genre_id = g.id group by g.id having total > 3;

select distinct a.first_name, a.last_name from actors a 
join actor_movie am on am.actor_id = a.id
join movies m on m.id = am.movie_id 
where m.title like 'La Guerra de las galaxias%';