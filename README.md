# tiny-index-db

### A tiny key-value disk-based Golang database that is capable of querying one type of indices

[![Go Reference](https://pkg.go.dev/badge/github.com/oddmario/tiny-index-db.svg)](https://pkg.go.dev/github.com/oddmario/tiny-index-db)

-----

## Documentation
TODO

## Features
- Write locks to avoid race conditions
- Concurrent non-blocking reading
- Tables to keep your little database organised

## Installation
```
go get -u github.com/oddmario/tiny-index-db
```

## Initialisation in your code
```go
import tinyindexdb "github.com/oddmario/tiny-index-db"

...
```

## Example usage

#### Create a new table
```go
tinyindexdb.NewTable("MyTable")
```
(returns an error only on the occurrence of an I/O error)

#### Destroy (remove) a table
```go
tinyindexdb.DestroyTable("MyTable")
```
(returns an error only on the occurrence of an I/O error)

#### Check if a table exists
```go
tinyindexdb.TableExists("NotMyTable")
```
(returns a boolean depending on the situation)

#### Query a record
```go
tinyindexdb.Query("MyTable", "MyRecordIndex")
```
returns:
- a `map[string]interface{}` if the record exists in the table.
- a `the specified table does not exist` error if MyTable does not exist.
- a `the specified record does not exist` error if MyRecordIndex does not exist.
- a `corrupted record data` error if the record has corrupted JSON data.

#### Write/update a record
```go
tinyindexdb.Query("MyTable", "MyRecordIndex", map[string]interface{}{
    "name": "mario",
    "message": "i like pianos",
    "timestamp": time.Now().Unix(),
})
```
(returns a `the specified table does not exist` error if MyTable does not exist)

Note that at the moment, for updating a record, you still need to pass the whole `map[string]interface{}`. A proper `UpdateRecord()` function may be available in the future.

#### Delete a record
```go
tinyindexdb.DeleteRecord("MyTable", "MyRecordIndex")
```
returns:
- an error on the occurrence of an I/O error
- a `the specified record does not exist` error if MyRecordIndex does not exist.

#### Clear the datababse  cache
```go
tinyindexdb.ClearCache()
```