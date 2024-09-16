#!/bin/bash
set -e

until psql -h "postgres" -U "admin" -d "godex" -c '\q'; do
  >&2 echo "Postgres is unavailable - sleeping"
  sleep 1
done

psql -h "postgres" -U "admin" -d "godex" -c "CREATE TABLE public.packages (id SERIAL PRIMARY KEY, name VARCHAR(128) NOT NULL, link VARCHAR(256) NOT NULL);"
psql -h "postgres" -U "admin" -d "godex" -c "CREATE TABLE public.files (id SERIAL PRIMARY KEY, name VARCHAR(128) NOT NULL, package_id INT REFERENCES public.packages(id) ON DELETE CASCADE);"
psql -h "postgres" -U "admin" -d "godex" -c "CREATE TABLE public.functions (id SERIAL PRIMARY KEY, name VARCHAR(128) NOT NULL, signature VARCHAR(128) NOT NULL, file_id INT REFERENCES public.files(id) ON DELETE CASCADE, comment TEXT);"