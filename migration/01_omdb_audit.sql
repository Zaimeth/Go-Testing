-- Drop table

-- DROP TABLE public.omdb_audit;

CREATE TABLE public.omdb_audit (
	id bigserial NOT NULL,
    client_path varchar(255) NULL,
    client_response jsonb NULL,
	created_at timestamp NULL,
    updated_at timestamp NULL,
    deleted_at timestamp NULL
);
