// package main

// import "fmt"

// type IDatabaseRepo interface {
// 	create() string
// 	update() string
// 	delete() string
// }

// type SQLDatabaseRepoImpl struct{}

// // delete implements IDatabaseRepo.
// func (SQLDatabaseRepoImpl) delete() string {
// 	return "SQL: deleting data..."
// }

// // update implements IDatabaseRepo.
// func (SQLDatabaseRepoImpl) update() string {
// 	return "SQL: updating data..."
// }

// func (SQLDatabaseRepoImpl) create() string {
// 	return "SQL: created a record in database..."
// }

// type NoSQLDatabaseRepoImpl struct{}

// // create implements IDatabaseRepo.
// func (NoSQLDatabaseRepoImpl) create() string {
// 	return "NO-SQL: created a record in database..."
// }

// // delete implements IDatabaseRepo.
// func (NoSQLDatabaseRepoImpl) delete() string {
// 	return "NO-SQL: deleting data..."
// }

// // update implements IDatabaseRepo.
// func (NoSQLDatabaseRepoImpl) update() string {
// 	return "NO-SQL: updating data..."
// }

// func main() {
// 	// down below is an example of Dependency injection using interfaces

// 	// var sqlDatabase IDatabaseRepo = SQLDatabaseRepoImpl{}
// 	// fmt.Println(sqlDatabase.create())
// 	// fmt.Println(sqlDatabase.update())
// 	// fmt.Println(sqlDatabase.delete())

// 	// var noSqlDatabase IDatabaseRepo = NoSQLDatabaseRepoImpl{}
// 	// fmt.Println(noSqlDatabase.create())
// 	// fmt.Println(noSqlDatabase.update())
// 	// fmt.Println(noSqlDatabase.delete())

// 	databases := []IDatabaseRepo{SQLDatabaseRepoImpl{}, NoSQLDatabaseRepoImpl{}}
// 	for _, database := range databases {
// 		fmt.Println(database.create())
// 		fmt.Println(database.update())
// 		fmt.Println(database.delete())
// 	}
// }
