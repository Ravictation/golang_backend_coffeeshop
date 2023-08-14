CREATE TABLE public.products (
	id_product uuid NOT NULL DEFAULT uuid_generate_v4(),
	product_name varchar NOT NULL,
	price int8 NOT NULL,
	categories varchar NOT NULL,
	product_image varchar NULL,
	created_at timestamp NOT NULL DEFAULT now(),
	updated_at timestamp NULL,
	CONSTRAINT products_pkey PRIMARY KEY (id_product)
);