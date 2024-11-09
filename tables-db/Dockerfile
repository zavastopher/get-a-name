FROM postgres:latest
ENV POSTGRES_USER oracle
ENV POSTGRES_PASSWORD supersecurepassword
COPY 1randomtables.sql /docker-entrypoint-initdb.d/
COPY names/*.sql /docker-entrypoint-initdb.d/
