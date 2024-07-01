# Clean Architecture Project

## Established by

- Kokiat Tankoem

### Introduction

- This project is to build a REST API with Golang and GO Fiber using the clean architecture of Uncle Bob.
- thank for ref Ruangyot Nanchiang

> !!! เอกสารเพิ่มเติม อยู่ใน Folder Document

### Requirements

- GO 1.21.3+
- PostgreSQL

### nstall the postgreSQL on Docker

**Pull Image** -> <a href="https://hub.docker.com/_/postgres" target="_blank">PostgreSQL Docker Image</a>

```cmd
docker pull postgres:alpine
```

Run the container

```cmd
docker run --name clean-arch-db -e POSTGRES_PASSWORD=123456 -p 1122:5432 -d postgres:alpine
```

Config the postgres

```cmd
docker exec -it clean-arch-db bash
```

```cmd
psql -U postgres
```

```cmd
create database clean_arch_db;
```

```cmd
\c clean_arch_db;
```

```cmd
CREATE TABLE "users"(
    "id" SERIAL PRIMARY KEY,
    "username" VARCHAR(255) UNIQUE NOT NULL,
    "password" VARCHAR(255) NOT NULL
);
```

Check the database that created or not

```cmd
\l
```

```cmd
      Name      |  Owner   | Encoding |  Collate   |   Ctype    |   Access privileges
----------------+----------+----------+------------+------------+-----------------------
 clean_arch_db | postgres | UTF8     | en_US.utf8 | en_US.utf8 |
```

### 🚀 Start a Project

```zsh
cd app/
go build main.go&&./main.exe
```

---
