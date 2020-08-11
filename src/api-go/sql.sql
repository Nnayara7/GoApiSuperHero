create database superhero;

\c superhero;

create table if not exists super (
	id SERIAL PRIMARY KEY,
	uuid VARCHAR(40) NOT NULL UNIQUE,
	name VARCHAR(100) NOT NULL,
	intelligence INT,
	strength INT,
	speed INT,
	durability INT,
	power INT,
	combat INT,
	created_at timestamp default current_timestamp,
	update_at timestamp default current_timestamp
);


create table if not exists biography (
	id SERIAL PRIMARY KEY,
	uuid VARCHAR(40) NOT NULL UNIQUE,
	full_name VARCHAR(500),
	alter_ego VARCHAR(500),
	aliases VARCHAR(1000),
	place_birth VARCHAR(500),
	first_appearance text,
	publisher VARCHAR(500),
	alignment VARCHAR(50)
);

create table if not exists appearance (
	id SERIAL PRIMARY KEY,
	uuid VARCHAR(40) NOT NULL UNIQUE,
	gender VARCHAR(50),
	race VARCHAR(255),
	height VARCHAR(255),
	weight VARCHAR(255),
	eye_color VARCHAR(40),
	hair_color VARCHAR(40),
	alignment VARCHAR(50)
);

create table if not exists work (
	id SERIAL PRIMARY KEY,
	uuid VARCHAR(40) NOT NULL UNIQUE,
	occupation VARCHAR(500),
	base VARCHAR(255)
);

create table if not exists connections (
	id SERIAL PRIMARY KEY,
	uuid VARCHAR(40) NOT NULL UNIQUE,
	group_affiliation text,
	relatives text
);

create table if not exists image (
	id SERIAL PRIMARY KEY,
	uuid VARCHAR(40) NOT NULL UNIQUE,
	url text
);
