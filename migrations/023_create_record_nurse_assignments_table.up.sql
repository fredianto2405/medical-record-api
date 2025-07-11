CREATE TABLE emr_core.record_nurse_assignments (
       id uuid NOT NULL DEFAULT uuid_generate_v4(),
       medical_record_id uuid,
       nurse_id uuid,
       created_at timestamp without time zone DEFAULT now(),
       CONSTRAINT record_nurse_assignments_pkey PRIMARY KEY (id),
       CONSTRAINT record_nurse_assignments_medical_record_id_fkey FOREIGN KEY (medical_record_id) REFERENCES emr_core.medical_records(id)
);