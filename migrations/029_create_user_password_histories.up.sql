CREATE TABLE emr_auth.user_password_histories (
  id bigint GENERATED ALWAYS AS IDENTITY NOT NULL,
  created_at timestamp without time zone NOT NULL,
  user_id uuid NOT NULL,
  password_hash text NOT NULL,
  source character varying NOT NULL,
  CONSTRAINT user_password_histories_pkey PRIMARY KEY (id),
  CONSTRAINT user_password_histories_user_id_fkey FOREIGN KEY (user_id) REFERENCES emr_auth.users(id)
);