# Go-Crud
Criando uma CRUD API em GO

## Pacotes externos utilizados:
- CompileDaemon (Copilador Automatico)
- Godotenv (Variaveis de Ambiente)
- Gorm (ORM)
- Gin (Web Framework)

### Comando para rodar o container do Postgres:
- docker run --name postgres-db -p 5432:5432 -e POSTGRES_PASSWORD=mysecretpassword -d postgres