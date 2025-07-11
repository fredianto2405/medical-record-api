CREATE TABLE emr_core.medical_record_statuses (
      id integer NOT NULL,
      name character varying,
      description character varying,
      sort_number integer,
      created_at timestamp without time zone DEFAULT now(),
      deleted_at timestamp without time zone,
      CONSTRAINT medical_record_statuses_pkey PRIMARY KEY (id)
);