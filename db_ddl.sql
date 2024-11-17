-- DROP SCHEMA public;

CREATE SCHEMA public AUTHORIZATION "admin";

-- DROP SEQUENCE public."Event_id_seq";

CREATE SEQUENCE public."Event_id_seq"
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 2147483647
	START 1
	CACHE 1
	NO CYCLE;-- public."Event" definition

-- Drop table

-- DROP TABLE public."Event";

CREATE TABLE public."Event" (
	id int4 NOT NULL GENERATED ALWAYS AS IDENTITY( INCREMENT BY 1 MINVALUE 1 MAXVALUE 2147483647 START 1 CACHE 1 NO CYCLE),
	"name" varchar NULL,
	begindate timestamp NULL,
	enddate timestamp NULL,
	details varchar NULL,
	architect int4 NOT NULL
);


-- public.architects definition

-- Drop table

-- DROP TABLE public.architects;

CREATE TABLE public.architects (
	id int4 NULL,
	"name" varchar NULL
);


-- public.eventmanager definition

-- Drop table

-- DROP TABLE public.eventmanager;

CREATE TABLE public.eventmanager (
	id int4 NULL,
	managerid int4 NULL,
	eventid int4 NULL
);


-- public.manager definition

-- Drop table

-- DROP TABLE public.manager;

CREATE TABLE public.manager (
	id int4 NULL,
	"name" varchar NULL,
	birthday timestamp NULL,
	qualifications varchar NULL,
	photo varchar NULL,
	bio varchar NULL
);


-- public.schedule definition

-- Drop table

-- DROP TABLE public.schedule;

CREATE TABLE public.schedule (
	id int4 NULL,
	eventid int4 NULL,
	managerid int4 NULL,
	volunteerid int4 NULL,
	startdate timestamp NULL,
	enddate timestamp NULL,
	weakly bool NULL
);


-- public."user" definition

-- Drop table

-- DROP TABLE public."user";

CREATE TABLE public."user" (
	id int4 NULL,
	"password" varchar NULL,
	"user" varchar NULL,
	"role" int4 NULL,
	profileid int4 NULL,
	"token" varchar NULL
);


-- public.volunteer definition

-- Drop table

-- DROP TABLE public.volunteer;

CREATE TABLE public.volunteer (
	id int4 NULL,
	"name" varchar NULL,
	birthday timestamp NULL,
	photo varchar NULL,
	bio varchar NULL
);
