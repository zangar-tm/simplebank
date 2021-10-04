package db 

var testQueries *Queries

func TestMain(t *testing.M) {
	conn, err := sql.Open(dbDriver,dbSource)
}