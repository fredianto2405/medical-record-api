CREATE TABLE emr_core.record_treatment_details (
       id uuid NOT NULL DEFAULT uuid_generate_v4(),
       medical_record_id uuid,
       treatment_id uuid,
       price integer,
       created_at timestamp without time zone DEFAULT now(),
       CONSTRAINT record_treatment_details_pkey PRIMARY KEY (id),
       CONSTRAINT record_treatment_details_medical_record_id_fkey FOREIGN KEY (medical_record_id) REFERENCES emr_core.medical_records(id)
);