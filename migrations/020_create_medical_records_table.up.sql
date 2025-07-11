CREATE TABLE emr_core.medical_records (
      id uuid NOT NULL DEFAULT uuid_generate_v4(),
      patient_id uuid,
      doctor_id uuid,
      diagnosis text,
      notes text,
      status_id integer,
      payment_method_id uuid,
      payment_status_id integer,
      insurance_id uuid,
      anamnesis text,
      created_at timestamp without time zone DEFAULT now(),
      CONSTRAINT medical_records_pkey PRIMARY KEY (id),
      CONSTRAINT medical_records_status_id_fkey FOREIGN KEY (status_id) REFERENCES emr_core.medical_record_statuses(id)
);