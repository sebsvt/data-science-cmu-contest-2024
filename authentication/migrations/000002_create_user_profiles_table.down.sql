CREATE TABLE user_profiles (
    user_id UUID PRIMARY KEY,
    first_name TEXT,
    last_name TEXT,
    phone_number TEXT,
    address TEXT,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
);
