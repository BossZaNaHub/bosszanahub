package bzmongo

type MongoClient interface {
	OpenMongoConnection() error
	CloseConnection()

	ListAllCollections() (int, []string)
	SelectCollection(collection string, filter interface{}) (interface{}, error)
}
