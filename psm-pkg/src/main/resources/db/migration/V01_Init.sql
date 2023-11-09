CREATE DATABASE psm_db;

CREATE SCHEMA public;

CREATE TABLE public.account (
    id integer NOT NULL,
    document_number character varying(20) NOT NULL
);

ALTER TABLE public.account OWNER TO postgres;

CREATE TABLE public.transaction (
    id uuid NOT NULL,
    account_id integer NOT NULL,
    operation_type integer NOT NULL,
    amount numeric NOT NULL,
    event_date timestamp without time zone NOT NULL
);

ALTER TABLE public.transaction OWNER TO postgres;

ALTER TABLE ONLY public.account
    ADD CONSTRAINT account_pk PRIMARY KEY (id);

ALTER TABLE ONLY public.account
    ADD CONSTRAINT account_pk2 UNIQUE (document_number);

ALTER TABLE ONLY public.transaction
    ADD CONSTRAINT transaction_pk PRIMARY KEY (id);

ALTER TABLE ONLY public.transaction
    ADD CONSTRAINT transaction_account__fk FOREIGN KEY (account_id) REFERENCES public.account(id);

CREATE INDEX idx_transaction_per_user_operation ON public.transaction (account_id, operation_type);

CREATE INDEX idx_transaction_per_user_period ON public.transaction (account_id, event_date);

CREATE INDEX idx_transaction_per_user_operation_period ON public.transaction (account_id, operation_type, event_date);