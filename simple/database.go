package simple

type DatabasePostgreSQL Database
type DatabaseMongoDB Database
type Database struct {
	Name string
}

type DatabaseRepository struct {
	DatabaseMongoDB    *DatabaseMongoDB
	DatabasePostgreSQL *DatabasePostgreSQL
}

func NewDatabaseMongoDB() *DatabaseMongoDB {
	return (*DatabaseMongoDB)(&Database{Name: "MongoDB"})
}
func NewDatabasePostgreSQL() *DatabasePostgreSQL {
	return (*DatabasePostgreSQL)(&Database{Name: "PostgreSQL"})
}

func NewDatabaseRepository(postgreSQL *DatabasePostgreSQL, mongoDB *DatabaseMongoDB) *DatabaseRepository {
	return &DatabaseRepository{DatabaseMongoDB: mongoDB, DatabasePostgreSQL: postgreSQL}
}
