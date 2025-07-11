CREATE TABLE emr_treatment.treatments (
      id uuid NOT NULL DEFAULT uuid_generate_v4(),
      name character varying,
      price integer,
      description text,
      created_at timestamp without time zone DEFAULT now(),
      deleted_at timestamp without time zone,
      CONSTRAINT treatments_pkey PRIMARY KEY (id)
);