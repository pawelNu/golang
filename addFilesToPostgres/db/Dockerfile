# Używamy oficjalnego obrazu PostgreSQL
FROM postgres:latest

ENV POSTGRES_PASSWORD=test

# Skopiowanie pliku SQL do folderu, w którym PostgreSQL automatycznie go załaduje
COPY schema.sql /docker-entrypoint-initdb.d/

# PostgreSQL używa domyślnych portów i zmiennych środowiskowych, nie musisz ich zmieniać
