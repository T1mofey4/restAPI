![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/evt/callback)

# restAPI

:book: 

Создание пользователя из psql:
                                CREATE USER postgres_user WITH PASSWORD '12345';

Создание базы из командной строки:      
                                        createdb -O postgres_user restapi_dev
                                        createdb -O postgres_user restapi_test

Миграции из командной строки:
                                migrate -path migrations -database "postgres://postgres_user:12345@localhost/restapi_dev?sslmode=disable" up
                                migrate -path migrations -database "postgres://postgres_user:12345@localhost/restapi_test?sslmode=disable" up