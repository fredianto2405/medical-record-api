CREATE TABLE emr_core.medical_record_histories (
       id uuid NOT NULL DEFAULT uuid_generate_v4(),
       medical_record_id uuid,
       status_id integer,
       created_at timestamp without time zone DEFAULT now(),
       CONSTRAINT medical_record_histories_pkey PRIMARY KEY (id),
       CONSTRAINT medical_record_histories_medical_record_id_fkey FOREIGN KEY (medical_record_id) REFERENCES emr_core.medical_records(id),
       CONSTRAINT medical_record_histories_status_id_fkey FOREIGN KEY (status_id) REFERENCES emr_core.medical_record_statuses(id)
);