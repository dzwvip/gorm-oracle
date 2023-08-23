# GORM Oracle Driver with go-ora


## Description

GORM Oracle driver for connect Oracle DB and Manage Oracle DB, not recommended for use in a production environment
## DB Driver
[go-ora](https://github.com/sijms/go-ora) A pure golang development of Oracle driver, do not need to install Oracle client.
## Required dependency Install

- Golang 1.17+
- gorm 1.24.0+
- go-ora v2.7.16
## Quick Start
### how to install 
```bash
go get github.com/dzwvip/gorm-oracle
```
###  usage

```go
import (
	"fmt"
      oracle "github.com/dzwvip/gorm-oracle"
	"gorm.io/gorm"
	"log"
)

func main() {
    db, err := gorm.Open(oracle.Open("oracle://username:password@127.0.0.1:1521/db"), &gorm.Config{})
    if err != nil {
        // panic error or log error info
    } 
    
    // do somethings
}
```
