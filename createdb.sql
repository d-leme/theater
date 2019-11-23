CREATE DATABASE theater_db
    WITH 
    OWNER = postgres
    ENCODING = 'UTF8'
    LC_COLLATE = 'en_US.utf8'
    LC_CTYPE = 'en_US.utf8'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1;


CREATE TABLE public.movies
(
    id uuid NOT NULL,
    name varchar(150) NOT NULL,
    category varchar(50) NOT NULL,
    PRIMARY KEY (id)
);

ALTER TABLE public.movies
    OWNER to postgres;