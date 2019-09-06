package main

import (
  "database/sql"
  "fmt"
  "log"
  _ "gopkg.in/goracle.v2"
)
func main() {
  fmt.Println(sql.Drivers())
  db, err := sql.Open("Oracle", "User Id=ESB_USER;Password=6tfc1qaz;Data Source=(DESCRIPTION=(ADDRESS_LIST=(LOAD_BALANCE=on)(LOAD_BALANCE=on)(FAILOVER=off)(ADDRESS =(PROTOCOL = TCP)(HOST = 172.16.62.55)(PORT = 1526))(ADDRESS = (PROTOCOL = TCP)(HOST = 172.16.62.56)(PORT = 1526)))(ADDRESS_LIST=(LOAD_BALANCE=off)(FAILOVER=on)(ADDRESS = (PROTOCOL = TCP)(HOST = 172.16.62.55)(PORT = 1526))(ADDRESS = (PROTOCOL = TCP)(HOST = 172.16.62.56)(PORT = 1526)))(CONNECT_DATA=(SERVER = DEDICATED)(SERVICE_NAME= EASF.kzc)(FAILOVER_MODE =(TYPE = SELECT)(METHOD = BASIC)(DELAY = 5)(RETRIES = 100))))")
  if err != nil {
    panic(err)
  }
  value, err := db.Query("select table_name from tabs where ... ", 1)

  var table_name string

  if err != nil {
    log.Fatal(err)
  }
defer value.Close()
  for value.Next() {
    err := value.Scan(&table_name)
    if err != nil {
      log.Fatal(err)
    }
    log.Println(table_name)
  }
  err = value.Err()
  if err != nil {
    log.Fatal(err)
  }
}