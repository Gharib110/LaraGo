package main

func doMigrate(arg2, arg3 string) error {
	//dsn := getDSN()
	checkForDB()

	tx, err := la.PopConnect()
	if err != nil {
		exitGracefully(err)
	}
	defer tx.Close()

	// run the migration command
	switch arg2 {
	case "up":
		err := la.RunPopMigrations(tx)
		if err != nil {
			return err
		}

	case "down":
		if arg3 == "all" {
			err := la.PopMigrateDown(tx, -1)
			if err != nil {
				return err
			}
		} else {
			err := la.PopMigrateDown(tx, 1)
			if err != nil {
				return err
			}
		}

	case "reset":
		err := la.PopMigrateReset(tx)
		if err != nil {
			return err
		}
	default:
		showHelp()
	}

	return nil
}
