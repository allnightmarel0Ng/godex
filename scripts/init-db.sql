DROP TABLE IF EXISTS public.functions;
DROP TABLE IF EXISTS public.files;
DROP TABLE IF EXISTS public.packages;

CREATE TABLE public.packages (
    id SERIAL PRIMARY KEY,
    name VARCHAR(128) NOT NULL,
    link VARCHAR(256) NOT NULL
);

CREATE TABLE public.files (
    id SERIAL PRIMARY KEY,
    name VARCHAR(128) NOT NULL,
    package_id INT REFERENCES public.packages(id) ON DELETE CASCADE
);

CREATE TABLE public.functions (
    id SERIAL PRIMARY KEY,
    name VARCHAR(128) NOT NULL,
    signature VARCHAR(128) NOT NULL,
    file_id INT REFERENCES public.files(id) ON DELETE CASCADE,
    comment TEXT
);