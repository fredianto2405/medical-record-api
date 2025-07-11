CREATE TABLE emr_menu.menus (
    id integer NOT NULL,
    name character varying,
    icon character varying,
    sort_number integer,
    path_name character varying,
    parent_id integer,
    is_active boolean,
    CONSTRAINT menus_pkey PRIMARY KEY (id)
);