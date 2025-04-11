CREATE TABLE IF NOT EXISTS receptions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    date_time TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    pvz_id UUID NOT NULL REFERENCES pvzs(id) ON DELETE CASCADE,
    status reception_status NOT NULL DEFAULT 'in_progress',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);
