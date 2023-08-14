CREATE TABLE public."user" (
	id_user uuid NOT NULL DEFAULT uuid_generate_v4(),
	email_user varchar NOT NULL,
	"password" varchar(255) NOT NULL,
	phone_number varchar(20) NOT NULL,
	created_at timestamp NOT NULL DEFAULT now(),
	updated_at timestamp NULL,
	image_user varchar(255) NULL,
	"role" varchar NULL,
	username varchar(255) NULL,
	CONSTRAINT user_email_user_key UNIQUE (email_user),
	CONSTRAINT user_pkey PRIMARY KEY (id_user),
	CONSTRAINT user_username_key UNIQUE (username)
);