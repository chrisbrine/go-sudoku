package sql

func Start(path string) (*DBData, error) {
	db, err := Connect(path)
	if err != nil {
		return nil, err
	}
	return db, nil
}
