CREATE USER pord WITH ENCRYPTED PASSWORD 'root';
CREATE DATABASE news_api;
GRANT ALL PRIVILEGES ON DATABASE news_api TO pord;