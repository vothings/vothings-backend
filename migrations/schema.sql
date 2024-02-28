--
-- PostgreSQL database dump
--

-- Dumped from database version 16.2 (Ubuntu 16.2-1.pgdg22.04+1)
-- Dumped by pg_dump version 16.2 (Ubuntu 16.2-1.pgdg22.04+1)

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
-- Name: multiple_choice_candidate; Type: TABLE; Schema: public; Owner: ekokurniawan
--

CREATE TABLE public.multiple_choice_candidate (
    id uuid NOT NULL,
    choice character varying,
    question_id uuid
);


ALTER TABLE public.multiple_choice_candidate OWNER TO ekokurniawan;

--
-- Name: multiple_choice_result; Type: TABLE; Schema: public; Owner: ekokurniawan
--

CREATE TABLE public.multiple_choice_result (
    id uuid NOT NULL,
    candidate_id uuid,
    users_id uuid
);


ALTER TABLE public.multiple_choice_result OWNER TO ekokurniawan;

--
-- Name: open_text_question; Type: TABLE; Schema: public; Owner: ekokurniawan
--

CREATE TABLE public.open_text_question (
    id uuid NOT NULL,
    question character varying,
    users_id uuid
);


ALTER TABLE public.open_text_question OWNER TO ekokurniawan;

--
-- Name: open_text_result; Type: TABLE; Schema: public; Owner: ekokurniawan
--

CREATE TABLE public.open_text_result (
    id uuid NOT NULL,
    answer character varying,
    question_id uuid,
    users_id uuid
);


ALTER TABLE public.open_text_result OWNER TO ekokurniawan;

--
-- Name: question_answer_quiz; Type: TABLE; Schema: public; Owner: ekokurniawan
--

CREATE TABLE public.question_answer_quiz (
    id uuid NOT NULL,
    choice character varying,
    question_id uuid
);


ALTER TABLE public.question_answer_quiz OWNER TO ekokurniawan;

--
-- Name: question_multiple_choice; Type: TABLE; Schema: public; Owner: ekokurniawan
--

CREATE TABLE public.question_multiple_choice (
    id uuid NOT NULL,
    question character varying,
    users_id uuid
);


ALTER TABLE public.question_multiple_choice OWNER TO ekokurniawan;

--
-- Name: question_quiz; Type: TABLE; Schema: public; Owner: ekokurniawan
--

CREATE TABLE public.question_quiz (
    id uuid NOT NULL,
    question character varying,
    users_id uuid
);


ALTER TABLE public.question_quiz OWNER TO ekokurniawan;

--
-- Name: question_rating; Type: TABLE; Schema: public; Owner: ekokurniawan
--

CREATE TABLE public.question_rating (
    id uuid NOT NULL,
    question character varying,
    users_id uuid
);


ALTER TABLE public.question_rating OWNER TO ekokurniawan;

--
-- Name: question_word_cloud; Type: TABLE; Schema: public; Owner: ekokurniawan
--

CREATE TABLE public.question_word_cloud (
    id uuid NOT NULL,
    question character varying,
    users_id uuid
);


ALTER TABLE public.question_word_cloud OWNER TO ekokurniawan;

--
-- Name: quiz_result; Type: TABLE; Schema: public; Owner: ekokurniawan
--

CREATE TABLE public.quiz_result (
    id uuid NOT NULL,
    answer character varying,
    question_answer_id uuid,
    users_id uuid
);


ALTER TABLE public.quiz_result OWNER TO ekokurniawan;

--
-- Name: rating_result; Type: TABLE; Schema: public; Owner: ekokurniawan
--

CREATE TABLE public.rating_result (
    id uuid NOT NULL,
    answer character varying,
    question_rating_id uuid,
    users_id uuid
);


ALTER TABLE public.rating_result OWNER TO ekokurniawan;

--
-- Name: schema_migration; Type: TABLE; Schema: public; Owner: ekokurniawan
--

CREATE TABLE public.schema_migration (
    version character varying(14) NOT NULL
);


ALTER TABLE public.schema_migration OWNER TO ekokurniawan;

--
-- Name: users; Type: TABLE; Schema: public; Owner: ekokurniawan
--

CREATE TABLE public.users (
    id uuid NOT NULL,
    username character varying,
    email character varying,
    password character varying,
    role character varying,
    created_at timestamp without time zone
);


ALTER TABLE public.users OWNER TO ekokurniawan;

--
-- Name: word_cloud_result; Type: TABLE; Schema: public; Owner: ekokurniawan
--

CREATE TABLE public.word_cloud_result (
    id uuid NOT NULL,
    answer character varying,
    users_id uuid,
    question_id uuid
);


ALTER TABLE public.word_cloud_result OWNER TO ekokurniawan;

--
-- Name: multiple_choice_candidate multiple_choice_candidate_pkey; Type: CONSTRAINT; Schema: public; Owner: ekokurniawan
--

ALTER TABLE ONLY public.multiple_choice_candidate
    ADD CONSTRAINT multiple_choice_candidate_pkey PRIMARY KEY (id);


--
-- Name: multiple_choice_result multiple_choice_result_pkey; Type: CONSTRAINT; Schema: public; Owner: ekokurniawan
--

ALTER TABLE ONLY public.multiple_choice_result
    ADD CONSTRAINT multiple_choice_result_pkey PRIMARY KEY (id);


--
-- Name: open_text_question open_text_question_pkey; Type: CONSTRAINT; Schema: public; Owner: ekokurniawan
--

ALTER TABLE ONLY public.open_text_question
    ADD CONSTRAINT open_text_question_pkey PRIMARY KEY (id);


--
-- Name: open_text_result open_text_result_pkey; Type: CONSTRAINT; Schema: public; Owner: ekokurniawan
--

ALTER TABLE ONLY public.open_text_result
    ADD CONSTRAINT open_text_result_pkey PRIMARY KEY (id);


--
-- Name: question_answer_quiz question_answer_quiz_pkey; Type: CONSTRAINT; Schema: public; Owner: ekokurniawan
--

ALTER TABLE ONLY public.question_answer_quiz
    ADD CONSTRAINT question_answer_quiz_pkey PRIMARY KEY (id);


--
-- Name: question_multiple_choice question_multiple_choice_pkey; Type: CONSTRAINT; Schema: public; Owner: ekokurniawan
--

ALTER TABLE ONLY public.question_multiple_choice
    ADD CONSTRAINT question_multiple_choice_pkey PRIMARY KEY (id);


--
-- Name: question_quiz question_quiz_pkey; Type: CONSTRAINT; Schema: public; Owner: ekokurniawan
--

ALTER TABLE ONLY public.question_quiz
    ADD CONSTRAINT question_quiz_pkey PRIMARY KEY (id);


--
-- Name: question_rating question_rating_pkey; Type: CONSTRAINT; Schema: public; Owner: ekokurniawan
--

ALTER TABLE ONLY public.question_rating
    ADD CONSTRAINT question_rating_pkey PRIMARY KEY (id);


--
-- Name: question_word_cloud question_word_cloud_pkey; Type: CONSTRAINT; Schema: public; Owner: ekokurniawan
--

ALTER TABLE ONLY public.question_word_cloud
    ADD CONSTRAINT question_word_cloud_pkey PRIMARY KEY (id);


--
-- Name: quiz_result quiz_result_pkey; Type: CONSTRAINT; Schema: public; Owner: ekokurniawan
--

ALTER TABLE ONLY public.quiz_result
    ADD CONSTRAINT quiz_result_pkey PRIMARY KEY (id);


--
-- Name: rating_result rating_result_pkey; Type: CONSTRAINT; Schema: public; Owner: ekokurniawan
--

ALTER TABLE ONLY public.rating_result
    ADD CONSTRAINT rating_result_pkey PRIMARY KEY (id);


--
-- Name: schema_migration schema_migration_pkey; Type: CONSTRAINT; Schema: public; Owner: ekokurniawan
--

ALTER TABLE ONLY public.schema_migration
    ADD CONSTRAINT schema_migration_pkey PRIMARY KEY (version);


--
-- Name: users users_email_key; Type: CONSTRAINT; Schema: public; Owner: ekokurniawan
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: ekokurniawan
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: word_cloud_result word_cloud_result_pkey; Type: CONSTRAINT; Schema: public; Owner: ekokurniawan
--

ALTER TABLE ONLY public.word_cloud_result
    ADD CONSTRAINT word_cloud_result_pkey PRIMARY KEY (id);


--
-- Name: schema_migration_version_idx; Type: INDEX; Schema: public; Owner: ekokurniawan
--

CREATE UNIQUE INDEX schema_migration_version_idx ON public.schema_migration USING btree (version);


--
-- Name: multiple_choice_candidate multiple_choice_candidate_question_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ekokurniawan
--

ALTER TABLE ONLY public.multiple_choice_candidate
    ADD CONSTRAINT multiple_choice_candidate_question_id_fkey FOREIGN KEY (question_id) REFERENCES public.question_multiple_choice(id);


--
-- Name: multiple_choice_result multiple_choice_result_candidate_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ekokurniawan
--

ALTER TABLE ONLY public.multiple_choice_result
    ADD CONSTRAINT multiple_choice_result_candidate_id_fkey FOREIGN KEY (candidate_id) REFERENCES public.multiple_choice_candidate(id);


--
-- Name: multiple_choice_result multiple_choice_result_users_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ekokurniawan
--

ALTER TABLE ONLY public.multiple_choice_result
    ADD CONSTRAINT multiple_choice_result_users_id_fkey FOREIGN KEY (users_id) REFERENCES public.users(id);


--
-- Name: open_text_question open_text_question_users_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ekokurniawan
--

ALTER TABLE ONLY public.open_text_question
    ADD CONSTRAINT open_text_question_users_id_fkey FOREIGN KEY (users_id) REFERENCES public.users(id);


--
-- Name: open_text_result open_text_result_question_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ekokurniawan
--

ALTER TABLE ONLY public.open_text_result
    ADD CONSTRAINT open_text_result_question_id_fkey FOREIGN KEY (question_id) REFERENCES public.open_text_question(id);


--
-- Name: open_text_result open_text_result_users_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ekokurniawan
--

ALTER TABLE ONLY public.open_text_result
    ADD CONSTRAINT open_text_result_users_id_fkey FOREIGN KEY (users_id) REFERENCES public.users(id);


--
-- Name: question_answer_quiz question_answer_quiz_question_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ekokurniawan
--

ALTER TABLE ONLY public.question_answer_quiz
    ADD CONSTRAINT question_answer_quiz_question_id_fkey FOREIGN KEY (question_id) REFERENCES public.question_quiz(id);


--
-- Name: question_multiple_choice question_multiple_choice_users_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ekokurniawan
--

ALTER TABLE ONLY public.question_multiple_choice
    ADD CONSTRAINT question_multiple_choice_users_id_fkey FOREIGN KEY (users_id) REFERENCES public.users(id);


--
-- Name: question_quiz question_quiz_users_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ekokurniawan
--

ALTER TABLE ONLY public.question_quiz
    ADD CONSTRAINT question_quiz_users_id_fkey FOREIGN KEY (users_id) REFERENCES public.users(id);


--
-- Name: question_rating question_rating_users_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ekokurniawan
--

ALTER TABLE ONLY public.question_rating
    ADD CONSTRAINT question_rating_users_id_fkey FOREIGN KEY (users_id) REFERENCES public.users(id);


--
-- Name: question_word_cloud question_word_cloud_users_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ekokurniawan
--

ALTER TABLE ONLY public.question_word_cloud
    ADD CONSTRAINT question_word_cloud_users_id_fkey FOREIGN KEY (users_id) REFERENCES public.users(id);


--
-- Name: quiz_result quiz_result_question_answer_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ekokurniawan
--

ALTER TABLE ONLY public.quiz_result
    ADD CONSTRAINT quiz_result_question_answer_id_fkey FOREIGN KEY (question_answer_id) REFERENCES public.question_answer_quiz(id);


--
-- Name: rating_result rating_result_question_rating_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ekokurniawan
--

ALTER TABLE ONLY public.rating_result
    ADD CONSTRAINT rating_result_question_rating_id_fkey FOREIGN KEY (question_rating_id) REFERENCES public.question_rating(id);


--
-- Name: rating_result rating_result_users_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ekokurniawan
--

ALTER TABLE ONLY public.rating_result
    ADD CONSTRAINT rating_result_users_id_fkey FOREIGN KEY (users_id) REFERENCES public.users(id);


--
-- Name: word_cloud_result word_cloud_result_question_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ekokurniawan
--

ALTER TABLE ONLY public.word_cloud_result
    ADD CONSTRAINT word_cloud_result_question_id_fkey FOREIGN KEY (question_id) REFERENCES public.question_word_cloud(id);


--
-- Name: word_cloud_result word_cloud_result_users_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ekokurniawan
--

ALTER TABLE ONLY public.word_cloud_result
    ADD CONSTRAINT word_cloud_result_users_id_fkey FOREIGN KEY (users_id) REFERENCES public.users(id);


--
-- PostgreSQL database dump complete
--

