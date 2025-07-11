CREATE TABLE emr_patient.medical_record_sequence (
     sequnce_date date NOT NULL,
     last_number integer,
     CONSTRAINT medical_record_sequence_pkey PRIMARY KEY (sequnce_date)
);