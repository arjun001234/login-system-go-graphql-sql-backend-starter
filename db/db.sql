
DROP DATABASE IF EXISTS login_system;
DROP USER IF EXISTS arjun;

CREATE USER arjun WITH PASSWORD 'arjun';
CREATE DATABASE login_system;
GRANT ALL PRIVILEGES ON DATABASE login_system to arjun;

\c login_system arjun;
CREATE SCHEMA dev;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA dev TO arjun;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";



