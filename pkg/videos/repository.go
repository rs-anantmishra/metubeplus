package videos

type IRepository interface{}

type repository struct {
	//here we have the db connection object (or the connection string?) to execute queries
	Connection string
}

func InstantiateRepo(conn string) IRepository {
	return &repository{
		Connection: conn,
	}
}
