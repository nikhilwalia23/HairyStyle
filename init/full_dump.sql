--
-- PostgreSQL database cluster dump
--

SET default_transaction_read_only = off;

SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;

--
-- Roles
--

DO $$ BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_roles WHERE rolname = 'authencation_app') THEN
        CREATE ROLE authencation_app WITH SUPERUSER INHERIT CREATEROLE CREATEDB LOGIN REPLICATION BYPASSRLS PASSWORD 'SCRAM-SHA-256$4096:805kX8Ic2kHefMp/nOOj2Q==$86ESODVyDgPf1hR3BsYDxQWW4wiC0Pw3/NVeErsDvXs=:41zziIijTiH+i5b+7pwt3nt5f8vEfJ0G6QPyH8crm84=';
    END IF;
END $$;
DO $$ BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_roles WHERE rolname = 'login_app_user') THEN
        CREATE ROLE login_app_user WITH NOSUPERUSER INHERIT NOCREATEROLE NOCREATEDB LOGIN NOREPLICATION NOBYPASSRLS PASSWORD 'hairStyle@123';
    END IF;
END $$;

--
-- User Configurations
--








--
-- Databases
--

--
-- Database "template1" dump
--

\connect template1

--
-- PostgreSQL database dump
--

-- Dumped from database version 17.4 (Debian 17.4-1.pgdg120+2)
-- Dumped by pg_dump version 17.4 (Debian 17.4-1.pgdg120+2)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- PostgreSQL database dump complete
--

--
-- Database "authencation_app" dump
--

--
-- PostgreSQL database dump
--

-- Dumped from database version 17.4 (Debian 17.4-1.pgdg120+2)
-- Dumped by pg_dump version 17.4 (Debian 17.4-1.pgdg120+2)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: authencation_app; Type: DATABASE; Schema: -; Owner: authencation_app
--

CREATE DATABASE authencation_app WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'en_US.utf8';
ALTER DATABASE authencation_app OWNER TO authencation_app;

\connect authencation_app

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: auth_application_schema; Type: SCHEMA; Schema: -; Owner: authencation_app
--

CREATE SCHEMA auth_application_schema;


ALTER SCHEMA auth_application_schema OWNER TO authencation_app;

--
-- Name: pgcrypto; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS pgcrypto WITH SCHEMA public;


--
-- Name: EXTENSION pgcrypto; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION pgcrypto IS 'cryptographic functions';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: users; Type: TABLE; Schema: auth_application_schema; Owner: authencation_app
--

CREATE TABLE auth_application_schema.users (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    name text NOT NULL,
    role text NOT NULL,
    email text NOT NULL,
    phone text NOT NULL,
    is_verified boolean DEFAULT false,
    aadhar_card character(12) NOT NULL,
    password text NOT NULL,
    created_at timestamp with time zone DEFAULT now(),
    CONSTRAINT users_aadhar_card_check CHECK ((aadhar_card ~ '^\d{12}$'::text)),
    CONSTRAINT users_password_check CHECK (((length(password) >= 8) AND (length(password) <= 20))),
    CONSTRAINT users_role_check CHECK ((role = ANY (ARRAY['user'::text, 'stylist'::text])))
);


ALTER TABLE auth_application_schema.users OWNER TO authencation_app;

--
-- Data for Name: users; Type: TABLE DATA; Schema: auth_application_schema; Owner: authencation_app
--

COPY auth_application_schema.users (id, name, role, email, phone, is_verified, aadhar_card, password, created_at) FROM stdin;
\.


--
-- Name: users users_email_key; Type: CONSTRAINT; Schema: auth_application_schema; Owner: authencation_app
--

ALTER TABLE ONLY auth_application_schema.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- Name: users users_phone_key; Type: CONSTRAINT; Schema: auth_application_schema; Owner: authencation_app
--

ALTER TABLE ONLY auth_application_schema.users
    ADD CONSTRAINT users_phone_key UNIQUE (phone);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: auth_application_schema; Owner: authencation_app
--

ALTER TABLE ONLY auth_application_schema.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: SCHEMA auth_application_schema; Type: ACL; Schema: -; Owner: authencation_app
--

GRANT ALL ON SCHEMA auth_application_schema TO login_app_user;


--
-- Name: TABLE users; Type: ACL; Schema: auth_application_schema; Owner: authencation_app
--

GRANT ALL ON TABLE auth_application_schema.users TO login_app_user;


--
-- Name: DEFAULT PRIVILEGES FOR TABLES; Type: DEFAULT ACL; Schema: auth_application_schema; Owner: authencation_app
--

ALTER DEFAULT PRIVILEGES FOR ROLE authencation_app IN SCHEMA auth_application_schema GRANT ALL ON TABLES TO login_app_user;
GRANT ALL PRIVILEGES ON SCHEMA auth_application_schema TO login_app_user;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA auth_application_schema TO login_app_user;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA auth_application_schema TO login_app_user;
GRANT ALL PRIVILEGES ON ALL FUNCTIONS IN SCHEMA auth_application_schema TO login_app_user;


--
-- PostgreSQL database dump complete
--

--
-- Database "logindb" dump
--

--
-- PostgreSQL database dump
--

-- Dumped from database version 17.4 (Debian 17.4-1.pgdg120+2)
-- Dumped by pg_dump version 17.4 (Debian 17.4-1.pgdg120+2)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: logindb; Type: DATABASE; Schema: -; Owner: authencation_app
--

CREATE DATABASE logindb WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'en_US.utf8';
ALTER DATABASE logindb OWNER TO authencation_app;

\connect logindb

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: auth_application_schema; Type: SCHEMA; Schema: -; Owner: authencation_app
--

CREATE SCHEMA auth_application_schema;


ALTER SCHEMA auth_application_schema OWNER TO authencation_app;

--
-- Name: SCHEMA auth_application_schema; Type: ACL; Schema: -; Owner: authencation_app
--

GRANT ALL ON SCHEMA auth_application_schema TO login_app_user;


--
-- Name: DEFAULT PRIVILEGES FOR TABLES; Type: DEFAULT ACL; Schema: auth_application_schema; Owner: authencation_app
--

ALTER DEFAULT PRIVILEGES FOR ROLE authencation_app IN SCHEMA auth_application_schema GRANT ALL ON TABLES TO login_app_user;


--
-- PostgreSQL database dump complete
--

--
-- Database "postgres" dump
--

\connect postgres

--
-- PostgreSQL database dump
--

-- Dumped from database version 17.4 (Debian 17.4-1.pgdg120+2)
-- Dumped by pg_dump version 17.4 (Debian 17.4-1.pgdg120+2)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- PostgreSQL database dump complete
--

--
-- PostgreSQL database cluster dump complete
--

