package app

import (
	"errors"
	"os/user"
	"path"
	"github.com/syndtr/goleveldb/leveldb"
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

func FetchCheckpoint(name string) (string, error) {
	db, err := GetDatabase()
	if err != nil {
		return "", err
	}

	defer db.Close()
	path, err := db.Get([]byte(name), nil)
	if err != nil {
		if err == leveldb.ErrNotFound{
			return "", errors.New("jump: Checkpoint not found")
		}
	return "", err
	}

	return string(path), nil
}

func AddCheckpoint(name string, path string) error {
	walk, err := FetchCheckpoint(name)
	if len(walk) > 0 {
		return errors.New("jump: Checkpoint already exists")
	}

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
	walk, err := FetchCheckpoint(name)
	if len(walk) == 0 {
		return errors.New("jump: Checkpoint doesn't exist")
	}

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

func ShowCheckpoints() (Checkpoints, error) {
	db, err := GetDatabase()
	checkpoints := Checkpoints{}

	if err != nil {
		return checkpoints, err
	}

	defer db.Close()
	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		key := iter.Key()
		value := iter.Value()

		checkpoints = append(checkpoints, Checkpoint{Name: string(key), Path: string(value),})
	}
	iter.Release()

	err = iter.Error()
	if err != nil {
		return checkpoints, err
	}

	return checkpoints, nil
}
