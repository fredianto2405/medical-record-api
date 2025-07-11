CREATE TABLE emr_nurse.nurses (
      id uuid NOT NULL DEFAULT uuid_generate_v4(),
      name character varying,
      gender character varying,
      address character varying,
      phone character varying,
      registration_number character varying,
      sharing_fee numeric,
      created_at timestamp without time zone DEFAULT now(),
      deleted_at timestamp without time zone,
      CONSTRAINT nurses_pkey PRIMARY KEY (id)
);