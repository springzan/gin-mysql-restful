package main

import (
    ."gin-mysql-restful/database/policydb"
    ."gin-mysql-restful/database/userdb"
    . "gin-mysql-restful/router"
)

func main() {
    defer PolicyDB.Close()
    defer UserDB.Close()
    router := InitRouter()
    router.Run(":2020")
}