CREATE TABLE "user" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar(60) NOT NULL,
  "surname" varchar(60) NOT NULL,
  "email" varchar(60) UNIQUE NOT NULL,
  "password" text NOT NULL,
  "address_id" integer,
  "user_type" varchar(20) NOT NULL
);

CREATE TABLE "category" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar(80)
);

CREATE TABLE "product" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar(60) NOT NULL,
  "price" integer NOT NULL,
  "user_id" integer,
  "category_id" integer,
  "address_id" integer
);

CREATE TABLE "country" (
  "id" SERIAL PRIMARY KEY,
  "country" varchar(60) NOT NULL
);

CREATE TABLE "city" (
  "id" SERIAL PRIMARY KEY,
  "country_id" integer,
  "city" varchar(60)
);

CREATE TABLE "address" (
  "id" SERIAL PRIMARY KEY,
  "post_code" integer NOT NULL,
  "city_id" integer
);

ALTER TABLE "product" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "product" ADD FOREIGN KEY ("category_id") REFERENCES "category" ("id");

ALTER TABLE "city" ADD FOREIGN KEY ("country_id") REFERENCES "country" ("id");

ALTER TABLE "address" ADD FOREIGN KEY ("city_id") REFERENCES "city" ("id");

ALTER TABLE "user" ADD FOREIGN KEY ("address_id") REFERENCES "address" ("id");

ALTER TABLE "product" ADD FOREIGN KEY ("address_id") REFERENCES "address" ("id");

ALTER TABLE "user" ALTER COLUMN "user_type" SET DEFAULT 'USER';