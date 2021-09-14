package database

var connection string

func init() { // <-- Package initialization
	connection = "PostgreSQL"
}

func GetDatabase() string {
	return connection
}
