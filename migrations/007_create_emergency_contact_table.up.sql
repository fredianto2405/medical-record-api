CREATE TABLE emr_patient.emergency_contact (
       id uuid NOT NULL DEFAULT uuid_generate_v4(),
       name character varying,
       phone character varying,
       relation character varying,
       patient_id uuid,
       created_at timestamp without time zone DEFAULT now(),
       deleted_at timestamp without time zone,
       CONSTRAINT emergency_contact_pkey PRIMARY KEY (id),
       CONSTRAINT emergency_contact_patient_id_fkey FOREIGN KEY (patient_id) REFERENCES emr_patient.patients(id)
);