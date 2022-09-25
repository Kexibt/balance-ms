#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "postgres" --dbname "postgres" <<-EOSQL
	CREATE DATABASE micro_balance;
	\c micro_balance
	CREATE TABLE public.balances (userID VARCHAR(100) PRIMARY KEY NOT NULL, balance FLOAT NOT NULL);
EOSQL