CREATE TABLE emr_payment.payment_methods (
     id uuid NOT NULL DEFAULT uuid_generate_v4(),
     name character varying,
     description character varying,
     created_at timestamp without time zone DEFAULT now(),
     deleted_at timestamp without time zone,
     CONSTRAINT payment_methods_pkey PRIMARY KEY (id)
);