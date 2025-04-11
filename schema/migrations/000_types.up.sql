DO $$
BEGIN

IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'role') THEN
    CREATE TYPE role AS ENUM ('employee', 'moderator');
END IF;

IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'city') THEN
    CREATE TYPE city AS ENUM ('Москва', 'Санкт-Петербург', 'Казань');
END IF;

IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'reception_status') THEN
    CREATE TYPE reception_status AS ENUM ('in_progress', 'close');
END IF;

IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'product_type') THEN
    CREATE TYPE product_type AS ENUM ('электроника', 'одежда', 'обувь');
END IF;

END $$;


-- применить миграции
-- migrate -path schema/migrations -database "postgres://postgres:03795@localhost:5432/avito_intern?sslmode=disable" up
