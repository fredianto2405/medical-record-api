CREATE TABLE emr_auth.users (
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    email character varying,
    password character varying,
    is_active boolean DEFAULT true,
    role_id integer,
    created_at timestamp without time zone DEFAULT now(),
    deleted_at timestamp without time zone,
    CONSTRAINT users_pkey PRIMARY KEY (id),
    CONSTRAINT users_role_id_fkey FOREIGN KEY (role_id) REFERENCES emr_auth.roles(id)
);