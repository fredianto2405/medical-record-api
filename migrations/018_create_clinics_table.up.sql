CREATE TABLE emr_clinic.clinics (
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    name character varying,
    address character varying,
    logo character varying,
    sharing_fee_type character varying,
    patient_medication_cost boolean,
    nurse_sharing_fee boolean,
    CONSTRAINT clinics_pkey PRIMARY KEY (id)
);