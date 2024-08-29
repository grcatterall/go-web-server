-- THIS DATA HAS BEEN AUTO GENERATED

CREATE TABLE "products" (
	id uuid PRIMARY KEY,
	name text UNIQUE,
	description text,
	price float8
);
