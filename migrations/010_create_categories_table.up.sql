CREATE TABLE emr_medicine.categories (
     id uuid NOT NULL DEFAULT uuid_generate_v4(),
     name character varying,
     created_at timestamp without time zone DEFAULT now(),
     deleted_at timestamp without time zone,
     CONSTRAINT categories_pkey PRIMARY KEY (id)
);