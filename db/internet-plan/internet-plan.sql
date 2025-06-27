CREATE TABLE internet_plan (
	`id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  	`velocity` decimal(6,1) DEFAULT NULL,
  	`price` decimal(6,1) DEFAULT NULL,
  	`discount` decimal(6,1) DEFAULT NULL,
  	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

CREATE TABLE client (
	`id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  	`first_name` varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  	`last_name` varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  	`date_of_birth` timestamp NULL DEFAULT NULL,
  	`city` varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  	`internet_plan_id` int(10) unsigned DEFAULT NULL,
  	PRIMARY KEY (`id`),
  	KEY `internet_plan_id_foreign` (`internet_plan_id`),
  	CONSTRAINT `internet_plan_id_foreign` FOREIGN KEY (`internet_plan_id`) REFERENCES `internet_plan` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

INSERT INTO internet_plan 
(velocity, price, discount) VALUES
(10.0, 50.0, 0.0),
(20.0, 75.5, 5.5),
(50.0, 110.0, 10.0),
(100.0, 150.0, 20.0),
(200.0, 200.0, 30.0);

INSERT INTO client 
(first_name, last_name, date_of_birth, city, internet_plan_id) VALUES
('Alice', 'Smith', '1990-01-15 10:00:00', 'New York', 1),
('Bob', 'Johnson', '1985-03-22 12:30:00', 'Los Angeles', 2),
('Charlie', 'Williams', '2000-05-10 08:15:00', 'Chicago', 3),
('Diana', 'Brown', '1992-07-01 16:45:00', 'Houston', 4),
('Edward', 'Jones', '1988-09-14 11:00:00', 'Phoenix', 5),
('Fiona', 'Garcia', '1995-11-23 05:25:00', 'Philadelphia', 1),
('George', 'Martinez', '1997-04-30 21:10:00', 'San Antonio', 2),
('Hannah', 'Rodriguez', '2001-02-12 19:00:00', 'San Diego', 3),
('Isaac', 'Lee', '1983-06-18 14:40:00', 'Dallas', 4),
('Joan', 'Walker', '1998-08-09 20:55:00', 'San Jose', 5);

select * from client c where c.internet_plan_id = 2;

select * from client c where c.city = 'Los Angeles';

select * from client c where c.city in ('Los Angeles', 'San Diego');

select * from client c where c.date_of_birth between '1998-01-01' and '2010-12-31';

select * from internet_plan i where i.price > 50;

select * from internet_plan i where i.discount > 10;

select * from internet_plan i where i.price < 200;

select * from internet_plan i where i.velocity = 100;

select * from client c join internet_plan i on i.id = c.internet_plan_id where i.price > 50;

select * from client c join internet_plan i on i.id = c.internet_plan_id where i.velocity = 50;