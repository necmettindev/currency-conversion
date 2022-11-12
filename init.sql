CREATE USER arf WITH PASSWORD '123arf';
ALTER USER arf WITH SUPERUSER;
ALTER ROLE arf CREATEROLE CREATEDB;
-- CREATE DATABASE conversion;
-- GRANT ALL PRIVILEGES ON DATABASE conversion to arf;
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
-- TOC entry 4 (class 2615 OID 2200)
-- Name: public; Type: SCHEMA; Schema: -; Owner: arf
--

-- *not* creating schema, since initdb creates it
ALTER SCHEMA public OWNER TO arf;
SET default_tablespace = '';
--
-- TOC entry 211 (class 1259 OID 17829)
-- Name: accounts; Type: TABLE; Schema: public; Owner: arf
--

CREATE TABLE public.accounts (
    id integer NOT NULL,
    user_id integer NOT NULL,
    currency character varying(3),
    balance numeric(15, 6),
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP
);
ALTER TABLE public.accounts OWNER TO arf;
--
-- TOC entry 213 (class 1259 OID 17871)
-- Name: accounts_id_seq; Type: SEQUENCE; Schema: public; Owner: arf
--

ALTER TABLE public.accounts
ALTER COLUMN id
ADD GENERATED ALWAYS AS IDENTITY (
        SEQUENCE NAME public.accounts_id_seq START WITH 1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1
    );
--
-- TOC entry 212 (class 1259 OID 17842)
-- Name: transactions; Type: TABLE; Schema: public; Owner: arf
--

CREATE TABLE public.transactions (
    id integer NOT NULL,
    user_id integer NOT NULL,
    first_currency character varying(5) [] NOT NULL,
    second_currency character varying(5) NOT NULL,
    amount numeric(15, 6) NOT NULL,
    rate numeric(15, 6) NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);
ALTER TABLE public.transactions OWNER TO arf;
--
-- TOC entry 209 (class 1259 OID 17797)
-- Name: users; Type: TABLE; Schema: public; Owner: arf
--

CREATE TABLE public.users (
    id integer NOT NULL,
    username character varying(16) NOT NULL,
    password character varying(155) NOT NULL,
    first_name character varying(30),
    last_name character varying(30),
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);
ALTER TABLE public.users OWNER TO arf;
--
-- TOC entry 210 (class 1259 OID 17802)
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: arf
--

ALTER TABLE public.users
ALTER COLUMN id
ADD GENERATED ALWAYS AS IDENTITY (
        SEQUENCE NAME public.users_id_seq START WITH 1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1
    );
--
-- TOC entry 3608 (class 0 OID 17829)
-- Dependencies: 211
-- Data for Name: accounts; Type: TABLE DATA; Schema: public; Owner: arf
--
INSERT INTO public.accounts (
id,
user_id,
currency,
balance,
created_at,
updated_at,
deleted_at
) OVERRIDING SYSTEM VALUE
VALUES (
        1,
        1,
        'USD',
        500,
        '2022-11-11 15:05:37.054269+03',
        '2022-11-11 15:05:37.054269+03',
        '2022-11-11 15:05:37.054269+03'
    );
INSERT INTO public.accounts (
        id,
        user_id,
        currency,
        balance,
        created_at,
        updated_at,
        deleted_at
    ) OVERRIDING SYSTEM VALUE
VALUES (
        2,
        1,
        'TRY',
        100,
        '2022-11-11 15:05:37.054269+03',
        '2022-11-11 15:05:37.054269+03',
        '2022-11-11 15:05:37.054269+03'
    );
--
-- TOC entry 3609 (class 0 OID 17842)
-- Dependencies: 212
-- Data for Name: transactions; Type: TABLE DATA; Schema: public; Owner: arf
--

--
-- TOC entry 3606 (class 0 OID 17797)
-- Dependencies: 209
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: arf
--

INSERT INTO public.users (
        id,
        username,
        password,
        first_name,
        last_name,
        created_at,
        updated_at,
        deleted_at
    ) OVERRIDING SYSTEM VALUE
VALUES (
        1,
        'sedatcan',
        '$2a$10$mfScCnzdR6QK2EmJhNQrmem0LZeGumob1Zp/KvhxbmGXu/08Wyd1K',
        'sedatcan',
        'sonat',
        '2022-11-11 13:43:23.612827+03',
        '2022-11-11 13:43:23.612827+03',
        NULL
    );
--
-- TOC entry 3617 (class 0 OID 0)
-- Dependencies: 213
-- Name: accounts_id_seq; Type: SEQUENCE SET; Schema: public; Owner: arf
--

SELECT pg_catalog.setval('public.accounts_id_seq', 21, true);
--
-- TOC entry 3618 (class 0 OID 0)
-- Dependencies: 210
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: arf
--

SELECT pg_catalog.setval('public.users_id_seq', 13, true);
--
-- TOC entry 3461 (class 2606 OID 17836)
-- Name: accounts accounts_pkey; Type: CONSTRAINT; Schema: public; Owner: arf
--

ALTER TABLE ONLY public.accounts
ADD CONSTRAINT accounts_pkey PRIMARY KEY (id);
--
-- TOC entry 3465 (class 2606 OID 17854)
-- Name: transactions transactions_pkey; Type: CONSTRAINT; Schema: public; Owner: arf
--

ALTER TABLE ONLY public.transactions
ADD CONSTRAINT transactions_pkey PRIMARY KEY (amount);
--
-- TOC entry 3463 (class 2606 OID 17873)
-- Name: accounts unique_currency_user_id; Type: CONSTRAINT; Schema: public; Owner: arf
--

ALTER TABLE ONLY public.accounts
ADD CONSTRAINT unique_currency_user_id UNIQUE (currency, user_id);
--
-- TOC entry 3457 (class 2606 OID 17801)
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: arf
--

ALTER TABLE ONLY public.users
ADD CONSTRAINT users_pkey PRIMARY KEY (id);
--
-- TOC entry 3459 (class 2606 OID 17808)
-- Name: users users_ukey; Type: CONSTRAINT; Schema: public; Owner: arf
--

ALTER TABLE ONLY public.users
ADD CONSTRAINT users_ukey UNIQUE (username);
--
-- TOC entry 3466 (class 2606 OID 17837)
-- Name: accounts fk_user_id; Type: FK CONSTRAINT; Schema: public; Owner: arf
--

ALTER TABLE ONLY public.accounts
ADD CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES public.users(id) NOT VALID;
--
-- TOC entry 3616 (class 0 OID 0)
-- Dependencies: 4
-- Name: SCHEMA public; Type: ACL; Schema: -; Owner: arf
--

REVOKE USAGE ON SCHEMA public
FROM PUBLIC;
GRANT ALL ON SCHEMA public TO PUBLIC;