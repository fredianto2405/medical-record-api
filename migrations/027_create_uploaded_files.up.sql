CREATE TABLE emr_clinic.uploaded_files (
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    filename TEXT NOT NULL,
    filepath TEXT NOT NULL,
    size BIGINT,
    mime_type TEXT,
    uploaded_at timestamp without time zone DEFAULT now(),
    CONSTRAINT uploaded_files_pkey PRIMARY KEY (id)
);