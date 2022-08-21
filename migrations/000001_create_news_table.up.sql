CREATE TABLE IF NOT EXISTS news (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR (255) NOT NULL,
    content TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT(now() AT TIME ZONE 'utc') ,
    updated_at TIMESTAMP DEFAULT NULL
);

-- CREATE INDEX idx_news_create_at_pagination ON news_api.news (created_at, id);


INSERT INTO news (title, content) VALUES ('1','1'),('2','2'),('3','3'),('4','4'),('5','5'),('6','6'),('7','7'),('8','8'),('9','9'),('10','10');
