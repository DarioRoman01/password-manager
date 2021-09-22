package store

import (
	"bytes"
	"encoding/gob"

	"github.com/dgraph-io/badger/v3"
)

type Password struct {
	Key string
	Pwd string
}

func NewPassword(key, pwd string) *Password {
	return &Password{key, pwd}
}

func (p *Password) Encode() ([]byte, error) {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)

	if err := encoder.Encode(p); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func DecodePassword(data []byte) (*Password, error) {
	var pwd Password
	decoder := gob.NewDecoder(bytes.NewReader(data))
	if err := decoder.Decode(&pwd); err != nil {
		return nil, err
	}

	return &pwd, nil
}

func FindPassWord(key string, db *badger.DB) *Password {
	var data []byte
	err := db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}

		data, _ = item.ValueCopy(nil)
		return nil
	})

	if err != nil {
		return nil
	}

	pwd, err := DecodePassword(data)
	if err != nil {
		return nil
	}

	return pwd
}

func SetPassword(pwd *Password, db *badger.DB) error {
	data, err := pwd.Encode()
	if err != nil {
		return err
	}

	return db.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(pwd.Key), data)
	})
}
