CREATE TABLE emr_core.record_recipes (
     id uuid NOT NULL DEFAULT uuid_generate_v4(),
     medical_record_id uuid,
     medicine_id uuid,
     price integer,
     quantity integer,
     dosage character varying,
     instruction character varying,
     created_at timestamp without time zone DEFAULT now(),
     CONSTRAINT record_recipes_pkey PRIMARY KEY (id),
     CONSTRAINT record_recipes_medical_record_id_fkey FOREIGN KEY (medical_record_id) REFERENCES emr_core.medical_records(id)
);