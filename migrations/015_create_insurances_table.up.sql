CREATE TABLE emr_payment.insurances (
        id uuid NOT NULL DEFAULT uuid_generate_v4(),
        name character varying,
        code character varying,
        contact character varying,
        email character varying,
        is_active boolean DEFAULT true,
        created_at timestamp without time zone DEFAULT now(),
        deleted_at timestamp without time zone,
        CONSTRAINT insurances_pkey PRIMARY KEY (id)
);