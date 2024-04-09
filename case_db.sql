--
-- PostgreSQL database dump
--

-- Dumped from database version 14.11 (Homebrew)
-- Dumped by pg_dump version 14.11 (Homebrew)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: cases; Type: TABLE; Schema: public; Owner: swarup
--

CREATE TABLE public.cases (
    id integer NOT NULL,
    created_at timestamp without time zone NOT NULL,
    client_first_name character varying(255),
    client_last_name character varying(255),
    phone_number character varying(20),
    email_address character varying(255),
    type character varying(255),
    description text,
    lawyer_id integer,
    address_street character varying(255),
    address_city character varying(255),
    address_state character(2),
    address_zip character(5)
);


ALTER TABLE public.cases OWNER TO swarup;

--
-- Name: cases_id_seq; Type: SEQUENCE; Schema: public; Owner: swarup
--

CREATE SEQUENCE public.cases_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.cases_id_seq OWNER TO swarup;

--
-- Name: cases_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: swarup
--

ALTER SEQUENCE public.cases_id_seq OWNED BY public.cases.id;


--
-- Name: gpt_resp; Type: TABLE; Schema: public; Owner: swarup
--

CREATE TABLE public.gpt_resp (
    id integer NOT NULL,
    created_at timestamp without time zone NOT NULL,
    meeting_id integer,
    questions text,
    summary text,
    points text
);


ALTER TABLE public.gpt_resp OWNER TO swarup;

--
-- Name: gpt_resp_id_seq; Type: SEQUENCE; Schema: public; Owner: swarup
--

CREATE SEQUENCE public.gpt_resp_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.gpt_resp_id_seq OWNER TO swarup;

--
-- Name: gpt_resp_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: swarup
--

ALTER SEQUENCE public.gpt_resp_id_seq OWNED BY public.gpt_resp.id;


--
-- Name: lawyers; Type: TABLE; Schema: public; Owner: swarup
--

CREATE TABLE public.lawyers (
    id integer NOT NULL,
    lawyer_first_name character varying(255),
    lawyer_last_name character varying(255),
    email_address character varying(255),
    password character varying(255)
);


ALTER TABLE public.lawyers OWNER TO swarup;

--
-- Name: lawyers_id_seq; Type: SEQUENCE; Schema: public; Owner: swarup
--

CREATE SEQUENCE public.lawyers_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.lawyers_id_seq OWNER TO swarup;

--
-- Name: lawyers_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: swarup
--

ALTER SEQUENCE public.lawyers_id_seq OWNED BY public.lawyers.id;


--
-- Name: meetings; Type: TABLE; Schema: public; Owner: swarup
--

CREATE TABLE public.meetings (
    id integer NOT NULL,
    created_at timestamp without time zone NOT NULL,
    case_id integer,
    lawyer_id integer,
    lawyer_notes text
);


ALTER TABLE public.meetings OWNER TO swarup;

--
-- Name: meetings_id_seq; Type: SEQUENCE; Schema: public; Owner: swarup
--

CREATE SEQUENCE public.meetings_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.meetings_id_seq OWNER TO swarup;

--
-- Name: meetings_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: swarup
--

ALTER SEQUENCE public.meetings_id_seq OWNED BY public.meetings.id;


--
-- Name: cases id; Type: DEFAULT; Schema: public; Owner: swarup
--

ALTER TABLE ONLY public.cases ALTER COLUMN id SET DEFAULT nextval('public.cases_id_seq'::regclass);


--
-- Name: gpt_resp id; Type: DEFAULT; Schema: public; Owner: swarup
--

ALTER TABLE ONLY public.gpt_resp ALTER COLUMN id SET DEFAULT nextval('public.gpt_resp_id_seq'::regclass);


--
-- Name: lawyers id; Type: DEFAULT; Schema: public; Owner: swarup
--

ALTER TABLE ONLY public.lawyers ALTER COLUMN id SET DEFAULT nextval('public.lawyers_id_seq'::regclass);


--
-- Name: meetings id; Type: DEFAULT; Schema: public; Owner: swarup
--

ALTER TABLE ONLY public.meetings ALTER COLUMN id SET DEFAULT nextval('public.meetings_id_seq'::regclass);


--
-- Data for Name: cases; Type: TABLE DATA; Schema: public; Owner: swarup
--

COPY public.cases (id, created_at, client_first_name, client_last_name, phone_number, email_address, type, description, lawyer_id, address_street, address_city, address_state, address_zip) FROM stdin;
3	2024-04-07 17:21:08.298046	Test	Client	1234567890	a@b.c	Contracts	I have a problem with contracts	1	\N	\N	\N	\N
4	2024-04-07 17:21:08.298046	Test2	Client2	1234567890	d@e.f	Divorce/Family Law	I have a problem with divorce	1	\N	\N	\N	\N
5	2024-04-07 17:38:50.489463	Test4	Test4	1234567899	g@h.i	Consumer Law	I have a problem with consumer law	1	\N	\N	\N	\N
6	2024-04-09 03:31:16.949363	TestFirst5	TestLast5	6097312540	mendyman@gmail.com	Employment	I was fired from my job	1	504 West 110th St Apt 8A	New York	NY	10024
7	2024-04-09 03:31:16.949363	TestFirst6	TestLast6	6097312540	A@b.c	Wills and Estates	I would like to create a will	1	2400 Virginia Ave NW	Washington	DC	20037
\.


--
-- Data for Name: gpt_resp; Type: TABLE DATA; Schema: public; Owner: swarup
--

COPY public.gpt_resp (id, created_at, meeting_id, questions, summary, points) FROM stdin;
\.


--
-- Data for Name: lawyers; Type: TABLE DATA; Schema: public; Owner: swarup
--

COPY public.lawyers (id, lawyer_first_name, lawyer_last_name, email_address, password) FROM stdin;
1	Unassigned			$2a$10$9M0HCdrknfgj8gbA1ZecZ.z1qqER.gCSNCOVfs7xcnDDWKyJdvs1S
2	Mendy	Fisch	mendyman@gmail.com	$2a$10$9M0HCdrknfgj8gbA1ZecZ.z1qqER.gCSNCOVfs7xcnDDWKyJdvs1S
3	a	b	c	d
4			a@b.c	$2a$10$atB5kk8om3ixfSCHdPXbcOuJUiRLewZrbyV6CBFmJccqaXRuZkQxm
7			b@c.d	$2a$10$8K0F8QVICNGYUWU4icttP.O/I1NVFjvEFYbVtExPoO66BEaquCIay
8			d@e.f	$2a$10$rg1zUlITVxgzqM/3El1SeemQnVeBGabZhYyJD2zi4slMeXL5GYIFi
9	Test	Lawyer2	h@i.j	$2a$10$I9eZAiZK9rTT3HwjbOVpWOeg8PUbT/CzBP4D72VsEEBU.uI5dRcny
11	Test	Lawyer3	k@l.m	$2a$10$O8HCLITvQ4jb4qF2.QOxTuwjbE9nZ0/zgd.rdOU.1I2lACGQ0NwWe
12	Mrimportant	Lawyer	greatemail@address.com	$2a$10$0pARFQ7gaSBYFgTLM7yJdOUALa7eW8iB2abbh82qmCaBS514THuMy
\.


--
-- Data for Name: meetings; Type: TABLE DATA; Schema: public; Owner: swarup
--

COPY public.meetings (id, created_at, case_id, lawyer_id, lawyer_notes) FROM stdin;
\.


--
-- Name: cases_id_seq; Type: SEQUENCE SET; Schema: public; Owner: swarup
--

SELECT pg_catalog.setval('public.cases_id_seq', 7, true);


--
-- Name: gpt_resp_id_seq; Type: SEQUENCE SET; Schema: public; Owner: swarup
--

SELECT pg_catalog.setval('public.gpt_resp_id_seq', 1, false);


--
-- Name: lawyers_id_seq; Type: SEQUENCE SET; Schema: public; Owner: swarup
--

SELECT pg_catalog.setval('public.lawyers_id_seq', 12, true);


--
-- Name: meetings_id_seq; Type: SEQUENCE SET; Schema: public; Owner: swarup
--

SELECT pg_catalog.setval('public.meetings_id_seq', 1, false);


--
-- Name: cases cases_pkey; Type: CONSTRAINT; Schema: public; Owner: swarup
--

ALTER TABLE ONLY public.cases
    ADD CONSTRAINT cases_pkey PRIMARY KEY (id);


--
-- Name: gpt_resp gpt_resp_pkey; Type: CONSTRAINT; Schema: public; Owner: swarup
--

ALTER TABLE ONLY public.gpt_resp
    ADD CONSTRAINT gpt_resp_pkey PRIMARY KEY (id);


--
-- Name: lawyers lawyers_email_address_key; Type: CONSTRAINT; Schema: public; Owner: swarup
--

ALTER TABLE ONLY public.lawyers
    ADD CONSTRAINT lawyers_email_address_key UNIQUE (email_address);


--
-- Name: lawyers lawyers_pkey; Type: CONSTRAINT; Schema: public; Owner: swarup
--

ALTER TABLE ONLY public.lawyers
    ADD CONSTRAINT lawyers_pkey PRIMARY KEY (id);


--
-- Name: meetings meetings_pkey; Type: CONSTRAINT; Schema: public; Owner: swarup
--

ALTER TABLE ONLY public.meetings
    ADD CONSTRAINT meetings_pkey PRIMARY KEY (id);


--
-- Name: cases cases_lawyer_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: swarup
--

ALTER TABLE ONLY public.cases
    ADD CONSTRAINT cases_lawyer_id_fkey FOREIGN KEY (lawyer_id) REFERENCES public.lawyers(id);


--
-- Name: gpt_resp gpt_resp_meeting_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: swarup
--

ALTER TABLE ONLY public.gpt_resp
    ADD CONSTRAINT gpt_resp_meeting_id_fkey FOREIGN KEY (meeting_id) REFERENCES public.meetings(id);


--
-- Name: meetings meetings_case_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: swarup
--

ALTER TABLE ONLY public.meetings
    ADD CONSTRAINT meetings_case_id_fkey FOREIGN KEY (case_id) REFERENCES public.cases(id);


--
-- Name: meetings meetings_lawyer_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: swarup
--

ALTER TABLE ONLY public.meetings
    ADD CONSTRAINT meetings_lawyer_id_fkey FOREIGN KEY (lawyer_id) REFERENCES public.lawyers(id);


--
-- PostgreSQL database dump complete
--

