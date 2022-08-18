SET check_function_bodies = false;
CREATE FUNCTION public.set_current_timestamp_updated_at() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
DECLARE
  _new record;
BEGIN
  _new := NEW;
  _new."updated_at" = NOW();
  RETURN _new;
END;
$$;
CREATE TABLE public.attachment (
    id bigint NOT NULL,
    reference_id bigint NOT NULL,
    reference_table text NOT NULL,
    name text NOT NULL,
    mime_type text NOT NULL,
    size integer NOT NULL,
    src text NOT NULL
);
CREATE SEQUENCE public.attachment_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
ALTER SEQUENCE public.attachment_id_seq OWNED BY public.attachment.id;
CREATE TABLE public.booking (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    destination_id bigint NOT NULL,
    vehicle_id bigint NOT NULL,
    "from" timestamp with time zone NOT NULL,
    until timestamp with time zone NOT NULL,
    email text NOT NULL,
    phone text NOT NULL,
    name text NOT NULL,
    code text NOT NULL
);
CREATE SEQUENCE public.booking_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
ALTER SEQUENCE public.booking_id_seq OWNED BY public.booking.id;
CREATE TABLE public.credential (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    identity text NOT NULL,
    provider text NOT NULL,
    password text NOT NULL,
    banned boolean NOT NULL
);
CREATE SEQUENCE public.credential_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
ALTER SEQUENCE public.credential_id_seq OWNED BY public.credential.id;
CREATE TABLE public.destination (
    id bigint NOT NULL,
    country_code character(2) NOT NULL,
    province text NOT NULL,
    district text NOT NULL,
    city text NOT NULL,
    zip_code text NOT NULL
);
CREATE SEQUENCE public.destination_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
ALTER SEQUENCE public.destination_id_seq OWNED BY public.destination.id;
CREATE TABLE public.headquarter (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    partner_id bigint NOT NULL,
    destination_id bigint NOT NULL,
    name text NOT NULL,
    point point
);
CREATE SEQUENCE public.headquarter_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
ALTER SEQUENCE public.headquarter_id_seq OWNED BY public.headquarter.id;
CREATE TABLE public.partner (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    activated_at timestamp with time zone DEFAULT now() NOT NULL,
    credential_id bigint NOT NULL,
    name text NOT NULL,
    type text NOT NULL
);
CREATE SEQUENCE public.partner_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
ALTER SEQUENCE public.partner_id_seq OWNED BY public.partner.id;
CREATE TABLE public.partner_type (
    value text NOT NULL,
    comment text NOT NULL
);
COMMENT ON TABLE public.partner_type IS 'enum table for partner type';
CREATE TABLE public.pricing (
    id bigint NOT NULL,
    term integer NOT NULL,
    price integer NOT NULL,
    tags text[] NOT NULL
);
CREATE SEQUENCE public.pricing_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
ALTER SEQUENCE public.pricing_id_seq OWNED BY public.pricing.id;
CREATE TABLE public.vehicle (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    partner_id bigint NOT NULL,
    headquarter_id bigint NOT NULL,
    brand text NOT NULL,
    model text NOT NULL,
    year integer NOT NULL
);
CREATE SEQUENCE public.vehicle_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
ALTER SEQUENCE public.vehicle_id_seq OWNED BY public.vehicle.id;
CREATE TABLE public.vehicle_pricing (
    id bigint NOT NULL,
    pricing_id bigint NOT NULL,
    vehicle_id bigint NOT NULL
);
CREATE SEQUENCE public.vehicle_pricing_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
ALTER SEQUENCE public.vehicle_pricing_id_seq OWNED BY public.vehicle_pricing.id;
ALTER TABLE ONLY public.attachment ALTER COLUMN id SET DEFAULT nextval('public.attachment_id_seq'::regclass);
ALTER TABLE ONLY public.booking ALTER COLUMN id SET DEFAULT nextval('public.booking_id_seq'::regclass);
ALTER TABLE ONLY public.credential ALTER COLUMN id SET DEFAULT nextval('public.credential_id_seq'::regclass);
ALTER TABLE ONLY public.destination ALTER COLUMN id SET DEFAULT nextval('public.destination_id_seq'::regclass);
ALTER TABLE ONLY public.headquarter ALTER COLUMN id SET DEFAULT nextval('public.headquarter_id_seq'::regclass);
ALTER TABLE ONLY public.partner ALTER COLUMN id SET DEFAULT nextval('public.partner_id_seq'::regclass);
ALTER TABLE ONLY public.pricing ALTER COLUMN id SET DEFAULT nextval('public.pricing_id_seq'::regclass);
ALTER TABLE ONLY public.vehicle ALTER COLUMN id SET DEFAULT nextval('public.vehicle_id_seq'::regclass);
ALTER TABLE ONLY public.vehicle_pricing ALTER COLUMN id SET DEFAULT nextval('public.vehicle_pricing_id_seq'::regclass);
ALTER TABLE ONLY public.attachment
    ADD CONSTRAINT attachment_pkey PRIMARY KEY (id);
ALTER TABLE ONLY public.booking
    ADD CONSTRAINT booking_pkey PRIMARY KEY (id);
ALTER TABLE ONLY public.credential
    ADD CONSTRAINT credential_identity_key UNIQUE (identity);
ALTER TABLE ONLY public.credential
    ADD CONSTRAINT credential_pkey PRIMARY KEY (id);
ALTER TABLE ONLY public.destination
    ADD CONSTRAINT destination_pkey PRIMARY KEY (id);
ALTER TABLE ONLY public.headquarter
    ADD CONSTRAINT headquarter_pkey PRIMARY KEY (id);
ALTER TABLE ONLY public.partner
    ADD CONSTRAINT partner_pkey PRIMARY KEY (id);
ALTER TABLE ONLY public.partner_type
    ADD CONSTRAINT partner_type_pkey PRIMARY KEY (value);
ALTER TABLE ONLY public.pricing
    ADD CONSTRAINT pricing_pkey PRIMARY KEY (id);
ALTER TABLE ONLY public.vehicle
    ADD CONSTRAINT vehicle_pkey PRIMARY KEY (id);
ALTER TABLE ONLY public.vehicle_pricing
    ADD CONSTRAINT vehicle_pricing_pkey PRIMARY KEY (id);
CREATE UNIQUE INDEX idx_identity ON public.credential USING btree (identity);
CREATE INDEX idx_reference_id ON public.attachment USING btree (reference_id);
CREATE TRIGGER set_public_booking_updated_at BEFORE UPDATE ON public.booking FOR EACH ROW EXECUTE FUNCTION public.set_current_timestamp_updated_at();
COMMENT ON TRIGGER set_public_booking_updated_at ON public.booking IS 'trigger to set value of column "updated_at" to current timestamp on row update';
CREATE TRIGGER set_public_headquarter_updated_at BEFORE UPDATE ON public.headquarter FOR EACH ROW EXECUTE FUNCTION public.set_current_timestamp_updated_at();
COMMENT ON TRIGGER set_public_headquarter_updated_at ON public.headquarter IS 'trigger to set value of column "updated_at" to current timestamp on row update';
CREATE TRIGGER set_public_vehicle_updated_at BEFORE UPDATE ON public.vehicle FOR EACH ROW EXECUTE FUNCTION public.set_current_timestamp_updated_at();
COMMENT ON TRIGGER set_public_vehicle_updated_at ON public.vehicle IS 'trigger to set value of column "updated_at" to current timestamp on row update';
ALTER TABLE ONLY public.booking
    ADD CONSTRAINT booking_destination_id_fkey FOREIGN KEY (destination_id) REFERENCES public.destination(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
ALTER TABLE ONLY public.booking
    ADD CONSTRAINT booking_vehicle_id_fkey FOREIGN KEY (vehicle_id) REFERENCES public.vehicle(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
ALTER TABLE ONLY public.headquarter
    ADD CONSTRAINT headquarter_destination_id_fkey FOREIGN KEY (destination_id) REFERENCES public.destination(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
ALTER TABLE ONLY public.headquarter
    ADD CONSTRAINT headquarter_partner_id_fkey FOREIGN KEY (partner_id) REFERENCES public.partner(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
ALTER TABLE ONLY public.partner
    ADD CONSTRAINT partner_credential_id_fkey FOREIGN KEY (credential_id) REFERENCES public.credential(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
ALTER TABLE ONLY public.partner
    ADD CONSTRAINT partner_type_fkey FOREIGN KEY (type) REFERENCES public.partner_type(value) ON UPDATE RESTRICT ON DELETE RESTRICT;
ALTER TABLE ONLY public.vehicle
    ADD CONSTRAINT vehicle_headquarter_id_fkey FOREIGN KEY (headquarter_id) REFERENCES public.headquarter(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
ALTER TABLE ONLY public.vehicle
    ADD CONSTRAINT vehicle_partner_id_fkey FOREIGN KEY (partner_id) REFERENCES public.partner(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
ALTER TABLE ONLY public.vehicle_pricing
    ADD CONSTRAINT vehicle_pricing_pricing_id_fkey FOREIGN KEY (pricing_id) REFERENCES public.pricing(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
ALTER TABLE ONLY public.vehicle_pricing
    ADD CONSTRAINT vehicle_pricing_vehicle_id_fkey FOREIGN KEY (vehicle_id) REFERENCES public.vehicle(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
