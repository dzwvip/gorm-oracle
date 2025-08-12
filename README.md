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
    //dsn:="oracle://username:password@127.0.0.1:1521/orcl"
    options := map[string]string{
		"client charset": "ZHS16GBK",
	}
    dsn:= oracle.BuildUrl("127.0.0.1", "1521", "orcl", "username", "password", options)    
    db, err := gorm.Open(oracle.Open(dsn), &gorm.Config{})
    if err != nil {
        // panic error or log error info
    } 
    
    // do somethings
}
```
### US7ASCII
ORACLE字符集为US7ASCII时，连接参数加上 client charset=ZHS16GBK 例如："oracle://system:manager@127.0.0.1:1521/orcl?client charset=ZHS16GBK"
### 创建Oracle函数生成指定Table和View对应的Struct
```sql
CREATE OR REPLACE FUNCTION getmodel(inuser varchar2,intablename varchar2) RETURN SYS_REFCURSOR
IS
  return_cursor SYS_REFCURSOR;
  dbuser varchar2(20);
  tablename varchar2(100);
BEGIN
  dbuser := upper(inuser);
  tablename  := upper(intablename);
  OPEN return_cursor FOR
  select 'package model ' from dual
  union all
  select 'type '|| replace(initcap(tablename),'_','') ||' struct {' from dual
  union all
  select cols from (
  select   replace(initcap(a.COLUMN_NAME),'_','')||' '|| decode(a.DATA_TYPE,'VARCHAR2','string','NUMBER',decode(DATA_SCALE,null,'int',0,'int','float64'),'DATE','utils.LocalTime','string')||' `json:"'||lower(a.COLUMN_NAME)||'"  gorm:"column:'||a.COLUMN_NAME||';"`'  as cols
  from All_Tab_Columns a where a.OWNER = dbuser and table_name =tablename order by a.COLUMN_ID )
  UNION all
  select '}' from DUAL
  union all
  select 'func ('||replace(initcap(tablename),'_','') ||') TableName() string {' from DUAL
  UNION all
  select 'return "'||dbuser||'.'||tablename ||'" }'from dual  ;

  RETURN return_cursor;
END;
```
