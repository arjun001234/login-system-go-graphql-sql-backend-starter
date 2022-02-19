DROP TABLE IF EXISTS dev.admins;
DROP TABLE IF EXISTS dev.users;
DROP TYPE IF EXISTS dev.user_role;
DROP TYPE IF EXISTS dev.provider_options

CREATE TYPE dev.user_role AS ENUM ('USER','ADMIN');
CREATE TYPE dev.provider_options AS ENUM ('GOOGLE','FACEBOOK');

CREATE TABLE dev.users(
  id UUID NOT NULL DEFAULT uuid_generate_v4 (),
  first_name VARCHAR(25) NOT NULL,
  last_name VARCHAR(25) NOT NULL,
  email VARCHAR(50) UNIQUE NOT NULL,
  password VARCHAR(60),
  user_role dev.user_role NOT NULL DEFAULT 'USER',
  created_at timestamptz DEFAULT NOW(),
  updated_at timestamptz DEFAULT NOW(),
  provider dev.provider_options,
  provider_id VARCHAR,
  picture TEXT,
  CONSTRAINT user_pk PRIMARY KEY (id)
);

CREATE TABLE dev.admins(
    user_id uuid NOT NULL UNIQUE ,
    CONSTRAINT admin_pk PRIMARY KEY (user_id)
);

ALTER TABLE dev.admins ADD CONSTRAINT admin_fk0 FOREIGN KEY (user_id) REFERENCES  dev.users(id) ON DELETE CASCADE ;
