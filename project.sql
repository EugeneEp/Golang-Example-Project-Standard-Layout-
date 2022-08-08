--
-- PostgreSQL database dump
--

-- Dumped from database version 13.7
-- Dumped by pg_dump version 14.4

-- Started on 2022-08-08 13:32:55

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

--
-- TOC entry 6 (class 2615 OID 16395)
-- Name: user; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA "user";


ALTER SCHEMA "user" OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 201 (class 1259 OID 16396)
-- Name: users; Type: TABLE; Schema: user; Owner: postgres
--

CREATE TABLE "user".users (
    id character varying NOT NULL,
    display_name character varying NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL
);


ALTER TABLE "user".users OWNER TO postgres;

--
-- TOC entry 2983 (class 0 OID 16396)
-- Dependencies: 201
-- Data for Name: users; Type: TABLE DATA; Schema: user; Owner: postgres
--

COPY "user".users (id, display_name, created_at) FROM stdin;
\.


--
-- TOC entry 2852 (class 2606 OID 16403)
-- Name: users users_pkey; Type: CONSTRAINT; Schema: user; Owner: postgres
--

ALTER TABLE ONLY "user".users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


-- Completed on 2022-08-08 13:32:55

--
-- PostgreSQL database dump complete
--

