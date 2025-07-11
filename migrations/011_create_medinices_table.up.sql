CREATE TABLE emr_medicine.medicines (
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    code character varying,
    name character varying,
    category_id uuid,
    unit_id uuid,
    price integer,
    stock integer,
    expiry_date timestamp without time zone,
    dosage character varying,
    created_at timestamp without time zone DEFAULT now(),
    deleted_at timestamp without time zone,
    CONSTRAINT medicines_pkey PRIMARY KEY (id),
    CONSTRAINT medicines_unit_id_fkey FOREIGN KEY (unit_id) REFERENCES emr_medicine.units(id),
    CONSTRAINT medicines_category_id_fkey FOREIGN KEY (category_id) REFERENCES emr_medicine.categories(id)
);