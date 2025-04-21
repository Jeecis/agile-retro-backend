package service

import (
	"crypto/rand"
	"math/big"
	"strconv"

	"github.com/google/uuid"
)

func boardIDGenerate() (string, error) {
	id := ""
	for i := 0; i < 6; i++ {

		numBig, err := rand.Int(rand.Reader, big.NewInt(36)) // 0-9 + A-Z = 36
		if err != nil {
			return "", err
		}

		num := numBig.Int64()
		if num < 10 {
			id += strconv.Itoa(int(num))
		} else {
			id += string('A' + num - 10) // Minus 10 since 0-9 is already taken by numbers
		}

	}
	return id, nil
}

func deletionIDGenerate() (string, error) {
	id := ""
	for i := 0; i < 10; i++ {

		numBig, err := rand.Int(rand.Reader, big.NewInt(36)) // 0-9 + A-Z = 36
		if err != nil {
			return "", err
		}

		num := numBig.Int64()
		if num < 10 {
			id += strconv.Itoa(int(num))
		} else {
			id += string('A' + num - 10) // Minus 10 since 0-9 is already taken by numbers
		}

	}
	return id, nil
}

func generateUUID() (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return id.String(), nil
}
