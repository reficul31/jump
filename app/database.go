package app

import (
    "errors"
    "os"
    "os/user"
    "path"
    "github.com/syndtr/goleveldb/leveldb"
)

// GetDatabasePath return the path of the database
func GetDatabasePath() (string, error) {
    usr, err := user.Current()
    if err != nil {
        return "", err
    }
    return path.Join(usr.HomeDir, ".jmpdata"), nil
}

// GetDatabase returns a levelDB object that can be used to access the database
func GetDatabase() (*leveldb.DB, error) {
    databasePath, err := GetDatabasePath()
    if err != nil {
        return nil, err
    }

    db, err := leveldb.OpenFile(databasePath, nil)
    if err != nil {
        return nil, err
    }

    return db, err
}

// DestroyDatabase removes the directory where the leveldb stores the checkpoints
func DestroyDatabase() error {
    databasePath, err := GetDatabasePath()
    if err != nil {
        return err
    }

    os.RemoveAll(databasePath)
    return nil
}

// FetchCheckpoint return the string path of a checkpoint in the database
func FetchCheckpoint(name string) (string, error) {
    db, err := GetDatabase()
    if err != nil {
        return "", err
    }

    defer db.Close()
    path, err := db.Get([]byte(name), nil)
    if err != nil {
        if err == leveldb.ErrNotFound {
            return "", errors.New("jump: Checkpoint not found")
        }
    return "", err
    }

    return string(path), nil
}

// AddCheckpoint adds a single key value pair in the form of name and path of a checkpoint to the database
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

// RemoveCheckpoint removes a single checkpoint from the database
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

// ShowCheckpoints return a list of all the checkpoints
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
