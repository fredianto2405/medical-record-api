CREATE TABLE emr_auth.password_resets (
  id bigint GENERATED ALWAYS AS IDENTITY NOT NULL,
  created_at timestamp without time zone NOT NULL,
  user_id uuid NOT NULL,
  token character varying NOT NULL,
  expired_at timestamp without time zone NOT NULL,
  used boolean NOT NULL DEFAULT false,
  CONSTRAINT password_resets_pkey PRIMARY KEY (id),
  CONSTRAINT password_resets_user_id_fkey FOREIGN KEY (user_id) REFERENCES emr_auth.users(id)
);