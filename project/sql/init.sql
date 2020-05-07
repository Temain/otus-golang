CREATE TABLE public.banners (
	id int NOT NULL GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
	title varchar NOT NULL
);

INSERT INTO public.banners (title) VALUES
('Баннер 1')
,('Баннер 2')
,('Баннер 3')
,('Баннер 4')
;


CREATE TABLE public.slots (
	id int NOT NULL GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
	title varchar NOT NULL
);

INSERT INTO public.slots (title) VALUES
('Слот 1')
,('Слот 2')
,('Слот 3')
;


CREATE TABLE public.groups (
	id int NOT NULL GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
	title varchar NOT NULL
);

INSERT INTO public.groups (title) VALUES
('Старики')
,('Дети')
;


CREATE TABLE public.rotation (
	banner_id int NOT NULL,
	slot_id int NOT NULL,
	started_at timestamp NOT NULL DEFAULT Now(),
	CONSTRAINT rotation_fk FOREIGN KEY (banner_id) REFERENCES public.banners(id),
	CONSTRAINT rotation_fk_2 FOREIGN KEY (slot_id) REFERENCES public.slots(id)
);
CREATE INDEX rotation_banner_id_idx ON public.rotation (banner_id,slot_id);

INSERT INTO public.rotation (banner_id,slot_id,started_at) VALUES
(1,1,'2020-05-01 00:00:00.000'),
(2,1,'2020-05-01 00:00:00.000'),
(3,1,'2020-05-01 00:00:00.000'),
(2,2,'2020-05-01 00:00:00.000'),
(3,2,'2020-05-01 00:00:00.000'),
(3,3,'2020-05-01 00:00:00.000')
;


CREATE TABLE public.statistics_type (
	id int NOT NULL GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
	title varchar NOT NULL
);

INSERT INTO public.statistics_type (title) VALUES
('Клик')
,('Показ')
;

CREATE TABLE public."statistics" (
	type_id int NOT NULL,
	banner_id int NOT NULL,
	slot_id int NOT NULL,
	group_id int NOT NULL,
	date_time timestamp NOT NULL,
	CONSTRAINT statistics_fk_type FOREIGN KEY (type_id) REFERENCES public.statistics_type(id),
	CONSTRAINT statistics_fk_slot FOREIGN KEY (slot_id) REFERENCES public.slots(id),
	CONSTRAINT statistics_fk_group FOREIGN KEY (group_id) REFERENCES public."groups"(id)
);
CREATE INDEX statistics_group_id_idx ON public."statistics" (type_id,banner_id,slot_id,group_id);

INSERT INTO public."statistics" (type_id,banner_id,slot_id,group_id,date_time) VALUES
(1,1,1,1,'2020-05-01 00:00:01.000'),
(1,2,1,1,'2020-05-01 00:00:02.000'),
(1,2,1,1,'2020-05-01 00:00:03.000'),
(1,2,1,1,'2020-05-01 00:00:04.000'),
(1,3,1,1,'2020-05-01 00:00:05.000'),
(1,3,1,1,'2020-05-01 00:00:06.000'),
(2,2,1,1,'2020-05-01 00:00:07.000'),
(1,1,1,2,'2020-05-01 00:00:01.000'),
(1,1,1,2,'2020-05-01 00:00:02.000'),
(1,1,1,2,'2020-05-01 00:00:03.000'),
(1,2,1,2,'2020-05-01 00:00:04.000'),
(1,3,1,2,'2020-05-01 00:00:05.000'),
(1,1,1,2,'2020-05-01 00:00:06.000'),
(2,1,1,2,'2020-05-01 00:00:07.000')
;

