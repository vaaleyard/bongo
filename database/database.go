package database

type Service struct {
	Client Database
}

type Database interface {
	ListDatabaseNames() ([]string, error)
	ListCollections(db string) ([]string, error)
	ListViews(db string) ([]string, error)
	ListUsers(db string) ([]string, error)
	RunCommand(db string, command string) string
}

func New(database Database) *Service {
	return &Service{
		Client: database,
	}
}
