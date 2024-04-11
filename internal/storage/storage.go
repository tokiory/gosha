package storage

import (
	"os"
	"path"
)

const storageFilename = ".gosha"

// storage is a function that returns path to the storage file
func storage() (string, error) {
  home, err := os.UserHomeDir()
  if err != nil {
    return "", err
  }
  
  return path.Join(home, storageFilename), nil
}

// Create is a function, that creates new storage file
func Create() (*os.File, error) {
  storagePath, err := storage()
  if err != nil {
    return nil, err
  }

  file, err := os.Create(storagePath)
  if err != nil {
    return nil, err
  }
  
  return file, nil
}

// Check if storage file exists
func Exist() (bool, error) {
  storagePath, err := storage()
  if err != nil {
    return false, err
  }

  if _, err = os.Stat(storagePath); err != nil && os.IsNotExist(err) {
    return false, nil
  }

  return true, nil
}

// Get is a function, that returns the storage file
func Get() (*os.File, error) {
  storagePath, err := storage()
  if err != nil {
    return nil, err
  }
  
  file, err := os.OpenFile(storagePath, os.O_APPEND | os.O_RDWR, 0666)
  if err != nil {
    return nil, err
  }

  return file, nil
}

// Write is a function, that write input string to the storage file
func Write(input string) error {
  storagePath, err := storage()
  if err != nil {
    return err
  }

  file, err := os.Open(storagePath)
  if err != nil {
    return err
  }

  _, err = file.WriteString(input)
  if err != nil {
    return err
  }

  return nil
}
