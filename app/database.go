package app

import (
	"fmt"
	"os"
	"os/user"
	"path"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/olekukonko/tablewriter"
)

func GetDatabase() (*leveldb.DB, error) {
	usr, err := user.Current()
	if err != nil {
		return nil, err
	}
	databasePath := path.Join(usr.HomeDir, ".jmpdata")
	db, err := leveldb.OpenFile(databasePath, nil)
	if err != nil {
		return nil, err
	}

	return db, err
}

func AddCheckpoint(name string, path string) error {
	db, err := GetDatabase()
	if err != nil {
		return err
	}

	defer db.Close()
	err = db.Put([]byte(name), []byte(path), nil)
	if err != nil {
		return err
	}

	return nil
}

func RemoveCheckpoint(name string) error {
	db, err := GetDatabase()
	if err != nil {
		return err
	}

	defer db.Close()
	err = db.Delete([]byte(name), nil)
	if err != nil {
		return err
	}

	return nil
}

func ShowCheckpoints() error {
	db, err := GetDatabase()
	if err != nil {
		return err
	}

	count := 0

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Path"})

	defer db.Close()
	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		key := iter.Key()
		value := iter.Value()

		table.Append([]string{string(key), string(value)})
		count+=1
	}
	iter.Release()

	err = iter.Error()
	if err != nil {
		return err
	}

	fmt.Printf("%v Checkpoints Found\n", count)
	table.Render()
	return nil
}

func FetchCheckpoint(name string) (string, error) {
	db, err := GetDatabase()
	if err != nil {
		return "", err
	}

	defer db.Close()
	path, err := db.Get([]byte(name), nil)
	if err != nil {
		return "", err
	}

	return string(path), nil
}