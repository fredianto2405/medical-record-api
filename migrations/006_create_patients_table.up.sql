CREATE TABLE emr_patient.patients (
      id uuid NOT NULL DEFAULT uuid_generate_v4(),
      medical_record_number character varying,
      name character varying,
      gender character varying,
      birth_date date,
      blood_type character varying,
      address character varying,
      phone character varying,
      email character varying,
      history_of_illness text,
      allergies text,
      identity_type character varying,
      identity_number character varying,
      created_at timestamp without time zone DEFAULT now(),
      deleted_at timestamp without time zone,
      CONSTRAINT patients_pkey PRIMARY KEY (id)
);