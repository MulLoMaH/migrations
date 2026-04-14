package main

import (
	"context"
	"log"
	"migrations/postgres_connect"
	"migrations/table_sql"
	//"migrations/postgres_connect"
	//"migrations/table_sql"
)

//Создание файлов миграции
//migrate create -ext sql -dir migrations -seq init

//migrate_библиотека
//  create_создать -ext_расширение
//  sql_название расширения базы
// -dir_директория
//  migrations_название папки в которой будут храниться миграции
//  -seq_указывает на создание названия
// init_название -

//Перемещение между версиями
//migrate -path migrations -database "postgres://postgres:2906@localhost:5432/postgres?sslmode=disable" up
//migrate -path migrations -database "postgres://postgres:2906@localhost:5432/postgres?sslmode=disable" down    (?sslmode=desable)
//migrate -path migrations -database "postgres://postgres:2906@localhost:5432/postgres?sslmode=disable" up 1
//migrate -path migrations -database "postgres://postgres:2906@localhost:5432/postgres?sslmode=disable" down

//migrate_библиотека
//_path_путь
//migrations_наша папка
//-database_название модуля
// "postgres://postgres:2906@localhost:5432/postgres?sslmode=desable" строка подключение + ?sslmode=desable - отключение шифрования
// up, down, up 1, down 1 - на сколько переместиться по версии (на все или 1)

//сброс флага грязной миграции
//migrate -path migrations -database "postgres://postgres:2906@localhost:5432/postgres?sslmode=disable" force 1

func main() {
	ctx := context.Background()
	conn, err := postgres_connect.Sql_conn(ctx)
	if err != nil {
		log.Fatal(err)
	}

	if err := table_sql.Create_table(ctx, conn); err != nil {
		log.Fatal(err)
	}

}
