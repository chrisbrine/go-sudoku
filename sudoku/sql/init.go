package sql

// func (db *DBData) CreateDatabase() error {
// 	fmt.Println("Creating database")
// 	_, err := db.db.Exec("CREATE DATABASE IF NOT EXISTS sudoku")
// 	if err != nil {
// 		fmt.Println("Error creating database:", err)
// 		return err
// 	}
// 	fmt.Println("Database created")

// 	return nil
// }

func Start(path string) (*DBData, error) {
	db, err := Connect(path)
	if err != nil {
		return nil, err
	}

	// create database if does not exist
	// err = db.CreateDatabase()
	// if err != nil {
	// 	return nil, err
	// }

	// create the game table
	err = db.CreateGameTable()
	if err != nil {
		return nil, err
	}

	// create the player table
	err = db.CreatePlayerTable()
	if err != nil {
		return nil, err
	}

	return db, nil
}
