# Chirpy

Chirpy je ukázkový Go projekt simulující základní Twitter-like aplikaci.
Umožňuje uživatelům přidávat "chirpy" (krátké zprávy do 140 znaků),
přihlašovat se pomocí JWT a ukládat data do PostgreSQL.

## Funkce
- Registrace a přihlášení uživatelů (JWT autentizace)
- Tvorba, čtení a mazání chirpů (REST API)
- Krátkodobé Access tokeny, odvolatelné Refresh tokeny
- Webhook příklady

## Technologie
- [Go](https://go.dev/) 1.20+
- [PostgreSQL](https://www.postgresql.org/)
- [Bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt)
- [SQLC](https://github.com/sqlc-dev/sqlc) pro generování DB dotazů

## Instalace a spuštění
1. Naklonujte repozitář:
   ```bash
   git clone https://github.com/vas-username/chirpy.git
   cd chirpy
   ```

2. Vytvořte .env s DB_URL, např.:
   ```
   DB_URL="postgres://user:pass@localhost:5432/chirpy?sslmode=disable"
   ```
3. Spusťte migrace (např. goose nebo sqlc).
4. Otestujte:
   ```
   go test ./...
   ```
5. Spusťte server:
   ```
   go run .
   ```
   Server bude na http://localhost:8080.
