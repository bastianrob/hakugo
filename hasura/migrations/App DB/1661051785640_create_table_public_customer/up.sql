CREATE TABLE "public"."customer" ("id" bigserial NOT NULL, "credential_id" bigint NOT NULL, "created_at" timestamptz NOT NULL DEFAULT now(), "updated_at" timestamptz NOT NULL DEFAULT now(), "activated_at" timestamptz, "name" text NOT NULL, "email" text NOT NULL, "phone" text NOT NULL, PRIMARY KEY ("id") , FOREIGN KEY ("credential_id") REFERENCES "public"."credential"("id") ON UPDATE restrict ON DELETE restrict, UNIQUE ("email"), UNIQUE ("phone"));
CREATE OR REPLACE FUNCTION "public"."set_current_timestamp_updated_at"()
RETURNS TRIGGER AS $$
DECLARE
  _new record;
BEGIN
  _new := NEW;
  _new."updated_at" = NOW();
  RETURN _new;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER "set_public_customer_updated_at"
BEFORE UPDATE ON "public"."customer"
FOR EACH ROW
EXECUTE PROCEDURE "public"."set_current_timestamp_updated_at"();
COMMENT ON TRIGGER "set_public_customer_updated_at" ON "public"."customer" 
IS 'trigger to set value of column "updated_at" to current timestamp on row update';
