# Description

The Facade Design Pattern is a structural design that provides a simplified interface to a complex and large body of code. It hides the complexities of the system and provides a single interface to the client.

## Example

Consider the following example:

```
package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

func main() {
	mysqlcfg := mysql.Config{
		Net:                  "tcp",
		Addr:                 fmt.Sprintf("%s:%S", os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT")),
		User:                 os.Getenv("MYSQL_USER"),
		Passwd:               os.Getenv("MYSQL_PASSWORD"),
		DBName:               os.Getenv("MYSQL_DBNAME"),
		ParseTime:            true,
		AllowNativePasswords: true,
		CheckConnLiveness:    true,
	}

	db, err := sql.Open("mysql", mysqlcfg.FormatDSN())
	if err != nil {
		log.Fatalf("Error connecting MySQL: %s\n", err)
	}

	db.Ping()
}
```

From the above example, we can hide the complexity of connecting to MySQL database by using a Facade Design Pattern, as follows:

```
// a facade
type MySQLDB struct {
	db *sql.DB
}

func (m *MySQLDB) Connect() error {
	mysqlcfg := mysql.Config{
		Net:                  "tcp",
		Addr:                 fmt.Sprintf("%s:%S", os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT")),
		User:                 os.Getenv("MYSQL_USER"),
		Passwd:               os.Getenv("MYSQL_PASSWORD"),
		DBName:               os.Getenv("MYSQL_DBNAME"),
		ParseTime:            true,
		AllowNativePasswords: true,
		CheckConnLiveness:    true,
	}

	db, err := sql.Open("mysql", mysqlcfg.FormatDSN())
	if err != nil {
		return err
	}

	m.db = db
	return nil
}
```

Now we just create this Facade and call its method `Connect()`.

```
func main() {
	mysqlDb := MySQLDB{}
	err := mysqlDb.Connect()
	if err != nil {
		log.Fatalf("Error connecting MySQL: %s\n", err)
	}

    ...
}
```
