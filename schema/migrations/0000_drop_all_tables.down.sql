DROP TABLE IF EXISTS products;
DROP TABLE IF EXISTS receptions;
DROP TABLE IF EXISTS pvzs;
DROP TABLE IF EXISTS users;


DO $$
BEGIN

    IF EXISTS (SELECT 1 FROM pg_type WHERE typname = 'product_type') THEN
        DROP TYPE product_type;
    END IF;

    IF EXISTS (SELECT 1 FROM pg_type WHERE typname = 'reception_status') THEN
        DROP TYPE reception_status;
    END IF;

    IF EXISTS (SELECT 1 FROM pg_type WHERE typname = 'city') THEN
        DROP TYPE city;
    END IF;

    IF EXISTS (SELECT 1 FROM pg_type WHERE typname = 'role') THEN
        DROP TYPE role;
    END IF;

END $$;