# GORM Oracle Driver with go-ora


## Description

GORM Oracle driver for connect Oracle DB and Manage Oracle DB,Based on [dzwvip/oracle](https://github.com/dzwvip/oracle)
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
### US7ASCII
ORACLE字符集为US7ASCII时，连接参数加上 client charset=ZHS16GBK 例如："oracle://system:manager@127.0.0.1:1521/orcl?client charset=ZHS16GBK"