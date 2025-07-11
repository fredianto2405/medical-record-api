CREATE TABLE emr_menu.menu_permissions (
       id integer GENERATED ALWAYS AS IDENTITY NOT NULL,
       menu_id integer,
       role_id integer,
       is_active boolean,
       CONSTRAINT menu_permissions_pkey PRIMARY KEY (id),
       CONSTRAINT menu_permissions_menu_id_fkey FOREIGN KEY (menu_id) REFERENCES emr_menu.menus(id)
);