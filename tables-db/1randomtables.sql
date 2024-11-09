CREATE DATABASE randomtables;

\c randomtables oracle
BEGIN;

CREATE TABLE names (
	category	varchar(20),
	culture 	varchar(20),
	name		varchar(20)
);

COMMIT;
