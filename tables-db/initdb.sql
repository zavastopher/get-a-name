CREATE USER oracle;
CREATE DATABASE randomtables;
GRANT ALL PRIVILEGES ON DATABASE randomtables to oracle;

CREATE TABLE names (
	category	varchar(10),
	culture 	varchar(10),
	name		varchar(10)
);
