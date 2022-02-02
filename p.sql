--
-- PostgreSQL database dump
--

-- Dumped from database version 9.5.9
-- Dumped by pg_dump version 9.5.9

-- Started on 2022-02-01 14:45:33

SET statement_timeout = 0;
SET lock_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

DROP DATABASE praktika1;
--
-- TOC entry 2130 (class 1262 OID 16393)
-- Name: praktika1; Type: DATABASE; Schema: -; Owner: postgres
--

CREATE DATABASE praktika1 WITH TEMPLATE = template0 ENCODING = 'UTF8' LC_COLLATE = 'Russian_Russia.1251' LC_CTYPE = 'Russian_Russia.1251';


ALTER DATABASE praktika1 OWNER TO postgres;

\connect praktika1

SET statement_timeout = 0;
SET lock_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- TOC entry 1 (class 3079 OID 12355)
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- TOC entry 2133 (class 0 OID 0)
-- Dependencies: 1
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET search_path = public, pg_catalog;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- TOC entry 183 (class 1259 OID 16409)
-- Name: disciplines; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE disciplines (
    id_disciplines integer NOT NULL,
    hours_of_practice integer,
    hours_of_lecture integer,
    disciplines character varying
);


ALTER TABLE disciplines OWNER TO postgres;

--
-- TOC entry 182 (class 1259 OID 16406)
-- Name: grupi; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE grupi (
    id_group integer NOT NULL,
    "group" character varying
);


ALTER TABLE grupi OWNER TO postgres;

--
-- TOC entry 181 (class 1259 OID 16400)
-- Name: load; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE load (
    id_load integer NOT NULL,
    id_discipline integer NOT NULL,
    id_group integer NOT NULL,
    id_teacher integer NOT NULL,
    year integer
);


ALTER TABLE load OWNER TO postgres;

--
-- TOC entry 185 (class 1259 OID 16504)
-- Name: teachers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE teachers (
    id_teacher integer NOT NULL,
    snp character varying,
    post character varying,
    date_of_hiring date
);


ALTER TABLE teachers OWNER TO postgres;

--
-- TOC entry 184 (class 1259 OID 16502)
-- Name: teachers_id_teacher_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE teachers_id_teacher_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE teachers_id_teacher_seq OWNER TO postgres;

--
-- TOC entry 2134 (class 0 OID 0)
-- Dependencies: 184
-- Name: teachers_id_teacher_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE teachers_id_teacher_seq OWNED BY teachers.id_teacher;


--
-- TOC entry 2123 (class 0 OID 16409)
-- Dependencies: 183
-- Data for Name: disciplines; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO disciplines (id_disciplines, hours_of_practice, hours_of_lecture, disciplines) VALUES (1, 100, 100, 'Program');
INSERT INTO disciplines (id_disciplines, hours_of_practice, hours_of_lecture, disciplines) VALUES (2, 90, 110, 'Liter');
INSERT INTO disciplines (id_disciplines, hours_of_practice, hours_of_lecture, disciplines) VALUES (3, 110, 90, 'Russian');
INSERT INTO disciplines (id_disciplines, hours_of_practice, hours_of_lecture, disciplines) VALUES (4, 120, 80, 'English');
INSERT INTO disciplines (id_disciplines, hours_of_practice, hours_of_lecture, disciplines) VALUES (5, 80, 120, 'Math');


--
-- TOC entry 2122 (class 0 OID 16406)
-- Dependencies: 182
-- Data for Name: grupi; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO grupi (id_group, "group") VALUES (1, 'PKS-101');
INSERT INTO grupi (id_group, "group") VALUES (2, 'PKS-202');
INSERT INTO grupi (id_group, "group") VALUES (3, 'PKS-303');
INSERT INTO grupi (id_group, "group") VALUES (4, 'PKS-404');
INSERT INTO grupi (id_group, "group") VALUES (5, 'PKS-505');


--
-- TOC entry 2121 (class 0 OID 16400)
-- Dependencies: 181
-- Data for Name: load; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO load (id_load, id_discipline, id_group, id_teacher, year) VALUES (1, 1, 1, 1, 2021);
INSERT INTO load (id_load, id_discipline, id_group, id_teacher, year) VALUES (2, 2, 2, 2, 2021);
INSERT INTO load (id_load, id_discipline, id_group, id_teacher, year) VALUES (3, 3, 3, 3, 2021);
INSERT INTO load (id_load, id_discipline, id_group, id_teacher, year) VALUES (4, 4, 4, 4, 2021);
INSERT INTO load (id_load, id_discipline, id_group, id_teacher, year) VALUES (5, 5, 5, 5, 2021);


--
-- TOC entry 2125 (class 0 OID 16504)
-- Dependencies: 185
-- Data for Name: teachers; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO teachers (id_teacher, snp, post, date_of_hiring) VALUES (2, 'Kamenev Vladimir Veniaminovich', 'Director', '2012-02-02');
INSERT INTO teachers (id_teacher, snp, post, date_of_hiring) VALUES (3, 'Brevno Vitaliy Vladimirovich', 'Prepod', '2013-03-03');
INSERT INTO teachers (id_teacher, snp, post, date_of_hiring) VALUES (4, 'Rechniy Nicolay Stepanovich', 'Zam', '2014-04-04');
INSERT INTO teachers (id_teacher, snp, post, date_of_hiring) VALUES (5, '1', '2', '2012-12-12');
INSERT INTO teachers (id_teacher, snp, post, date_of_hiring) VALUES (1, 'Bulkin', 'Prepod', '2012-11-11');
INSERT INTO teachers (id_teacher, snp, post, date_of_hiring) VALUES (0, '', '', '0001-01-01');
INSERT INTO teachers (id_teacher, snp, post, date_of_hiring) VALUES (8, '9', '9', '2012-12-12');


--
-- TOC entry 2135 (class 0 OID 0)
-- Dependencies: 184
-- Name: teachers_id_teacher_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('teachers_id_teacher_seq', 25, true);


--
-- TOC entry 2001 (class 2606 OID 16523)
-- Name: disciplines_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY disciplines
    ADD CONSTRAINT disciplines_pkey PRIMARY KEY (id_disciplines);


--
-- TOC entry 1999 (class 2606 OID 16521)
-- Name: grupi_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY grupi
    ADD CONSTRAINT grupi_pkey PRIMARY KEY (id_group);


--
-- TOC entry 1997 (class 2606 OID 16525)
-- Name: load_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY load
    ADD CONSTRAINT load_pkey PRIMARY KEY (id_load);


--
-- TOC entry 2003 (class 2606 OID 16519)
-- Name: teachers_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY teachers
    ADD CONSTRAINT teachers_pkey PRIMARY KEY (id_teacher);


--
-- TOC entry 2004 (class 2606 OID 16526)
-- Name: load_id_discipline_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY load
    ADD CONSTRAINT load_id_discipline_fkey FOREIGN KEY (id_discipline) REFERENCES disciplines(id_disciplines) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- TOC entry 2005 (class 2606 OID 16531)
-- Name: load_id_group_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY load
    ADD CONSTRAINT load_id_group_fkey FOREIGN KEY (id_group) REFERENCES grupi(id_group) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- TOC entry 2006 (class 2606 OID 16536)
-- Name: load_id_teacher_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY load
    ADD CONSTRAINT load_id_teacher_fkey FOREIGN KEY (id_teacher) REFERENCES teachers(id_teacher) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- TOC entry 2132 (class 0 OID 0)
-- Dependencies: 6
-- Name: public; Type: ACL; Schema: -; Owner: postgres
--

REVOKE ALL ON SCHEMA public FROM PUBLIC;
REVOKE ALL ON SCHEMA public FROM postgres;
GRANT ALL ON SCHEMA public TO postgres;
GRANT ALL ON SCHEMA public TO PUBLIC;


-- Completed on 2022-02-01 14:45:33

--
-- PostgreSQL database dump complete
--

