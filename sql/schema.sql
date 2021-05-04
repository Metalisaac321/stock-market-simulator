CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE account (
  id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
  cash real NOT NULL CHECK (cash >= 0::double precision)
);

CREATE TABLE issuer (
  id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
  issuer_name VARCHAR(10) NOT NULL,
  total_shares INTEGER NOT NULL CHECK (total_shares >= 0),
  share_price real NOT NULL CHECK (share_price > 0::double precision),
  account_id uuid NOT NULL REFERENCES account(id)
);

CREATE TYPE order_type_operation AS ENUM (
  'BUY',
  'SELL'
);

CREATE TABLE "order" (
  id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
  issuer_name VARCHAR(10) NOT NULL,
  total_shares INTEGER NOT NULL CHECK (total_shares >= 0),
  share_price real NOT NULL CHECK (share_price > 0::double precision),
  order_type_operation order_type_operation NOT NULL,
  account_id uuid NOT NULL REFERENCES account(id),
  created_at timestamp with time zone NOT NULL DEFAULT now()
);
