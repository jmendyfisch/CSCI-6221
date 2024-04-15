--
-- PostgreSQL database dump
--

-- Dumped from database version 14.11 (Ubuntu 14.11-0ubuntu0.22.04.1)
-- Dumped by pg_dump version 14.11 (Ubuntu 14.11-0ubuntu0.22.04.1)

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
    address_zip character(5),
    gpt_summary text
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

COPY public.cases (id, created_at, client_first_name, client_last_name, phone_number, email_address, type, description, lawyer_id, address_street, address_city, address_state, address_zip, gpt_summary) FROM stdin;
8	2024-04-12 12:26:47.029913	Sarah	Doe	2672054340	swarupgt2@gmail.com	Divorce/Family Law	Me and my husband are having a divorce, it is fine but we are unable to decide on the custody of our dog Max.	13	1600 S Joyce St	Arlington	VA	22202	Sarah and John are pursuing a mutual divorce after a 10-year marriage. Their primary point of contention is the custody of their dog, Max. Despite having amicably divided their assets, they both insist on full custody of Max. Sarah, currently based in Washington D.C, believes Max has acclimated better to her residence, which is larger and near a park - an advantage for the dog's exercise and overall happiness. She asserts she has a stronger bond with Max, largely due to her home-based work enabling more time spent with him. Though John initially wanted to purchase Max during their marriage, Sarah feels she has since developed a closer relationship with the dog. Her aim is to achieve peaceful resolution on this matter before the divorce is finalized.
11	2024-04-14 20:46:46.248546	Ram	Shekhar	4596879457	rshekhar@gmail.com	Wills and Estates	I would like to create a will in which I want to donate 90% of my assets towards wildlife preservation and conservation. The rest 10% would be equally be split between my wife and son.	13	132 Major Avenue	New York City	NY	11007	\N
13	2024-04-14 22:01:31.178893	Johnny	Silverhand	4538734567	silverhand11@mailtech.com	Employment	I have been recently fired from Arasaka Corporation on grounds of moonlighting, but this is not the case. I have only worked at Arasaka throughout my employment there. I would like to file a lawsuit for defamation and wrongful termination.	1	666 Kabuki Lane	Los Angeles	CA	90134	\N
9	2024-04-14 16:39:29.26351	John	Doe	1234567890	test@test.com	Landlord-Tenant	I received an eviction notice from my landlord when i was out of country serving in the military. 	13	16000	Arlington	VA	22202	Client received an eviction notice from the landlord while serving in the military overseas. The landlord alleges lack of rent payment and property damage. The client, based in Arlington, VA, highlights the roommate's responsibility for rent payment, who is unresponsive, and suspicions that the landlord is aware of their military service. The client seeks advice on handling these issues given their absence during the relevant time, emphasizing the need to establish the property's condition pre-departure, collect evidence of their upkeep, address the roommate's silence, and consider the landlord's knowledge of the client's military duties.
12	2024-04-14 21:54:06.269335	Alan	Wake	2342349876	notawake@remedy.fi	Contracts	I am an author, and I want to break my contract with my current publisher due to creative direction differences. I have been offered to meet with another famous publisher, Zane Printing Co., who seem to have a similar passion for the genre of my works. How can I switch to a new publisher without creating a potential lawsuit? 	13	854 National Road	New York City	NY	10001	An author from New York City seeks to terminate a publishing contract with current publisher due to creative differences and explore switching to Zane Printing Co. The lawyer advised reviewing the termination clause, negotiating with the current publisher, and seeking an amicable resolution to avoid potential legal disputes. The client aims for a smooth transition without inviting lawsuits.
\.


--
-- Data for Name: gpt_resp; Type: TABLE DATA; Schema: public; Owner: swarup
--

COPY public.gpt_resp (id, created_at, meeting_id, questions, summary, points) FROM stdin;
2	2024-04-12 12:38:23.272495	2	Can you provide more details on why both you and your husband want the custody of Max?, How has been your arrangement regarding Max’s care during your marriage, and who generally took more responsibility?, Are there any specific reasons to believe that Max would be better off with you or your husband?, Did you or your husband purchase Max or was he a gift?, Are there any previous agreements or discussions about Max you've had that could influence the decision?, Does Max show any preference in staying with you or John?, Have there been any incidents or issues in the past that may affect who gets custody of Max?, Aside from the dog, are there other contentious issues in the divorce that we should be aware of?, 	Sarah and her husband John, residents of Arlington, VA, are going through a divorce after 10 years of marriage. They don't have children but are in conflict over the custody of their dog, Max. They have already sorted out how to divide their assets and have been sharing expenses equally. Sarah owns slightly more and desires a peaceful resolution to the issue. She seeks guidance in navigating this complicated situation.	Sarah wants a peaceful resolution, important to tailor the case approach to minimize conflict, Sarah owns slightly more, she may be in a better financial position to care for Max, Consider researching state laws in Arlington, VA regarding pet custody in divorce cases., Check if pre-nuptial or post-nuptial agreements were in place regarding pet custody., Look for evidence that Max would be better off with either Sarah or John, consider factors such as work schedules, living conditions, Max's preferences, and past care arrangements., 
5	2024-04-14 16:41:51.570482	6	Can you provide more details on the condition of the house when you left for military service?\nDo you have any proof of the rent payments you made to your roommate?\nHave you tried discussing the situation with your roommate?\nIs there any way to reach out to your roommate to clarify the unpaid rent situation?\nHave you informed your landlord about the situation with your roommate and the rent payments?\n	The client, based in Arlington, VA, received an eviction notice from the landlord while serving in the military overseas. Despite paying rent to the roommate, the landlord claims rent wasn't received and the property was damaged. The client seeks guidance on addressing these issues given his absence during the period in question.	It's important to review the lease agreement to understand responsibilities and liabilities of both tenants.\nVerification of rent payments through bank statements or receipts will be crucial to support your case.\nCommunicating directly with the roommate to address the unpaid rent issue could help in resolving the situation.\nExplore options for legal protections provided to military personnel under the Servicemembers Civil Relief Act (SCRA).\n
1	2024-04-12 12:28:35.399161	1	Did you and your husband adopt Max together or was he initially owned by just one of you?\n What attempts have you made to resolve the custody issue?\n Are there any precedent for pet custody cases in Arlington, VA?\n Is there a shared custody or visitation agreement feasible in your situation?\n Do either of you have primarily taken the responsibilities for the dog?	Sarah is seeking legal advice regarding her divorce with her husband, John. They have managed to decide on the splitting of assets on their own, but the custody of their dog, Max, is proving to be a difficult area. She's looking for a fair solution that will minimize conflict.	Sarah earns significantly more than her husband which may impact the dog's care and lifestyle., Sarah seems to want to avoid unnecessary conflict, therefore a peaceful solution must be sought., The case is based in Arlington, VA which may have specific laws regarding pet custody., Both Sarah and John share equal responsibilities towards Max., 
3	2024-04-14 16:22:13.159223	3	Can you provide details about your ex's relationship with Max?\nWhat is John's daily routine with Max?\nDoes Max have any particular needs that need to be addressed?\nHave there been any incidents of negligence or harm towards Max from either of you?\nHas Max shown any behavioral changes since the separation took place?\nDid both of you jointly purchase Max or was he acquired by one of you individually?\n	During the latest meeting, Sarah expressed that her dog, Max, seems happier and has adjusted well to her new place in Washington, D.C., despite the unsettled divorce situation. Sarah and John, currently living separately, are both financially stable but are adamant about getting full custody of Max. Although they lived in a two-bedroom apartment in Arlington, VA, Sarah believes Max has more space and comfort in her current residence. The conflict of Max's custody remains unsolved.	Consider the change in Max's well-being observed by Sarah as a possible argument.\nResearch possibilities for joint custody or visitation schedules for pets.\nLook into the legal standpoint of pet custody in divorce cases in Virginia and Washington, D.C.\nInvestigate whether Max's preference, if it can be determined, could factor into the custody decision.\n
4	2024-04-14 16:26:19.415814	3	Has Max shown any preference towards either you or John?\nHave you discussed shared custody of Max?\nWhat kind of arrangements have been made so far for Max's care during this transition period?\nAre there any documents proving that John purchased Max from the shelter?\nDoes Max have any health concerns that necessitate one of you taking a primary caretaker role?\n	In the current meeting, the lawyer explored the client's and her ex's relationship with their dog, Max. The client admitted that Max was originally her ex's idea, but she believes she has formed a stronger bond with him, due in part to her ability to spend more time with Max since she works from home. They generally split duties, with each taking Max for a walk, but the client feels that her new place, near a park, is a better environment for Max. She revealed that they purchased Max after they were married, and while the purchase itself was made by her ex, she is keen on finalizing Max's custody before they officially separate.	There might be potential implications to John purchasing Max. Further investigation needed on pet ownership laws.\nEmotional attachments and pet's wellbeing should be main considerations for custody.\nPotential need to gather further evidence about life at each home.\nConsider pursuing shared custody or visitation rights if laws permit.\n
6	2024-04-14 16:46:04.70481	7	What steps did you take to ensure the apartment was left in good condition before leaving for military service?\nCan you provide any witnesses or evidence that can support your claim of taking good care of the apartment?\nDo you have a lease agreement that clearly states your responsibilities during your absence due to military service?\nHave you documented any communication with the roommate regarding the party and the damage to the kitchen sink?\nIs there any proof of attempts made to contact the roommate regarding these issues?\n	The client mentioned that despite taking good care of the apartment, there appeared to have been a party in their absence resulting in damage to the kitchen sink. The roommate, who was responsible for paying rent, is unresponsive. The client suspects the landlord is aware of their military service based on casual conversations. There's a need to establish the condition of the house before departure, gather evidence of care, address the roommate's absence, and consider the implications of the landlord's knowledge of military service.	It would be crucial to gather any evidence or witnesses supporting the client's claim of maintaining the apartment's good condition.\nReviewing the lease agreement to understand the responsibilities during the client's absence is essential.\nDocumenting communication attempts with the roommate can strengthen the client's position in the case.\nExploring legal protections under the Servicemembers Civil Relief Act (SCRA) due to military service could be beneficial in addressing landlord-tenant issues.\nConsidering the roommate's default on rent payment and potential negligence, further legal action may be necessary.\n
7	2024-04-15 00:22:51.212076	11	What specific terms in your current contract are concerning to you?\nHave you discussed with Zane Printing Company about the potential switch and their expectations?\nDo you have a timeline in mind for transitioning to a new publisher?\nAre there any key deadlines or obligations in your current contract that we should be aware of?\nHave you considered the financial implications of terminating your current contract early?\n	The client, an author from New York City, is seeking to break the current publishing contract due to creative disagreements. Interested in switching to Zane Printing Co., the client is wary of potential legal issues. The lawyer advised reviewing the termination clause, negotiating with the current publisher, and exploring an amicable resolution. The client aims to avoid conflicts and legal disputes, emphasizing a smooth transition to a new publisher.	Verify the specific termination clause in the current contract and assess its implications.\nConsider the importance of communication with both the current and potential publishers throughout the transition process.\nExplore alternative dispute resolution methods like mediation or arbitration if negotiation fails.\nEnsure any new contract offered by Zane Printing Co. is thoroughly reviewed to protect the client's interests.\n
\.


--
-- Data for Name: lawyers; Type: TABLE DATA; Schema: public; Owner: swarup
--

COPY public.lawyers (id, lawyer_first_name, lawyer_last_name, email_address, password) FROM stdin;
2	Mendy	Fisch	mendyman@gmail.com	$2a$10$9M0HCdrknfgj8gbA1ZecZ.z1qqER.gCSNCOVfs7xcnDDWKyJdvs1S
13	Swarup	Totloor	swarupgt@gmail.com	$2a$10$7T2DCIyehOoEjlHXS4953.W35DfbyRssuSbd1fp/ZDDVpwQk8cScO
1	Unassigned	\N	\N	\N
\.


--
-- Data for Name: meetings; Type: TABLE DATA; Schema: public; Owner: swarup
--

COPY public.meetings (id, created_at, case_id, lawyer_id, lawyer_notes) FROM stdin;
1	2024-04-12 12:28:35.390237	8	13	\N
2	2024-04-12 12:38:23.066379	8	13	\N
3	2024-04-14 16:18:24.974287	8	13	Consider shared custody between the client and her ex. 
4	2024-04-14 16:32:04.715385	8	13	\N
5	2024-04-14 16:34:23.978838	8	13	\N
8	2024-04-14 22:18:00.343751	11	13	\N
9	2024-04-14 23:30:02.16289	8	13	\N
10	2024-04-14 23:32:29.663326	8	13	\N
6	2024-04-14 16:39:57.819039	9	13	\N
7	2024-04-14 16:43:09.203561	9	13	\N
12	2024-04-15 00:26:00.599924	12	13	\N
11	2024-04-15 00:20:39.071928	12	13	Client appears to be adamant about switching publishers.
\.


--
-- Name: cases_id_seq; Type: SEQUENCE SET; Schema: public; Owner: swarup
--

SELECT pg_catalog.setval('public.cases_id_seq', 13, true);


--
-- Name: gpt_resp_id_seq; Type: SEQUENCE SET; Schema: public; Owner: swarup
--

SELECT pg_catalog.setval('public.gpt_resp_id_seq', 7, true);


--
-- Name: lawyers_id_seq; Type: SEQUENCE SET; Schema: public; Owner: swarup
--

SELECT pg_catalog.setval('public.lawyers_id_seq', 13, true);


--
-- Name: meetings_id_seq; Type: SEQUENCE SET; Schema: public; Owner: swarup
--

SELECT pg_catalog.setval('public.meetings_id_seq', 12, true);


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

