package main

import (
    db "gin-mysql-restful/database"
    userdb "gin-mysql-restful/database/userdb"
    . "gin-mysql-restful/router"
)

func main() {
    defer db.SqlDB.Close()
    defer userdb.UserDB.Close()
    router := InitRouter()
    router.Run(":2020")
}