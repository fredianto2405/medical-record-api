CREATE TABLE emr_medicine.units (
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    name character varying,
    created_at timestamp without time zone DEFAULT now(),
    deleted_at timestamp without time zone,
    CONSTRAINT units_pkey PRIMARY KEY (id)
);