CREATE TABLE emr_doctor.doctors (
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    email character varying,
    name character varying,
    gender character varying,
    specialization_id uuid,
    phone character varying,
    address character varying,
    registration_number character varying,
    sharing_fee numeric,
    created_at timestamp without time zone DEFAULT now(),
    deleted_at timestamp without time zone,
    CONSTRAINT doctors_pkey PRIMARY KEY (id),
    CONSTRAINT doctors_specialization_id_fkey FOREIGN KEY (specialization_id) REFERENCES emr_doctor.specializations(id)
);