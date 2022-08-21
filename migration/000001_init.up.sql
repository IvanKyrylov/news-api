CREATE TABLE IF NOT EXISTS news_api.news (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR (255) NOT NULL,
    content TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT(now() AT TIME ZONE 'utc') ,
    updated_at TIMESTAMP DEFAULT NULL
);

-- CREATE INDEX idx_news_create_at_pagination ON news_api.news (created_at, id);

