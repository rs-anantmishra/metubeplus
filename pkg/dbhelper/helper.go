package dbhelper

/*
- initialize DB from scripts
- Read script files to get queries to execute
*/

type IService interface {
	GetQuery(queryName string) string
}

func GetQuery(qName string) string {

	return ""
}
