CREATE TABLE emr_patient.insurance_patient (
       id uuid NOT NULL DEFAULT uuid_generate_v4(),
       insurance_id uuid,
       insurance_number character varying,
       patient_id uuid,
       created_at timestamp without time zone DEFAULT now(),
       deleted_at timestamp without time zone,
       CONSTRAINT insurance_patient_pkey PRIMARY KEY (id),
       CONSTRAINT insurance_patient_patient_id_fkey FOREIGN KEY (patient_id) REFERENCES emr_patient.patients(id)
);