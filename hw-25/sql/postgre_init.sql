CREATE DATABASE calendar;

CREATE TABLE public.events (
	id int4 NOT NULL GENERATED ALWAYS AS IDENTITY,
	title varchar(100) NOT NULL,
	description varchar(250) NOT NULL,
	created timestamp NOT NULL
);

CREATE INDEX events_created_idx ON public.events (created);