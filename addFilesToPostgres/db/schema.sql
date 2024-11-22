-- Tworzenie bazy danych (jeśli nie istnieje)
CREATE DATABASE my_database;

-- Używanie nowo utworzonej bazy danych
\c my_database;

-- Tworzenie schematu (jeśli nie istnieje)
CREATE SCHEMA IF NOT EXISTS my_schema;

-- Tworzenie tabeli w schemacie
CREATE TABLE my_schema.files (
    id SERIAL PRIMARY KEY,            
    file_name TEXT NOT NULL,          
    content_hash TEXT NOT NULL,       
    content TEXT NOT NULL,            
    CONSTRAINT unique_file UNIQUE(content_hash)
);
