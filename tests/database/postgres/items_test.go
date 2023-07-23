package postgres

import (
	"errors"
	"github.com/stretchr/testify/mock"
	"key-shop/internal/database/sql/postgres"
	"key-shop/tests/mocks"
	"testing"
)

const (
	testedLogin = "testedLogin3000"
)

var (
	testedItem = postgres.Item{
		Name:        "nameTestItem",
		Description: "descriptionTestItem",
		Count:       "100",
		Cost:        "100000",
	}

	testedError = errors.New("error")
)

func TestAddItem(t *testing.T) {
	storage := getTestedStorageExec(t, nil)
	err := storage.AddItem(testedLogin, testedItem)
	if err != nil {
		t.Error(err)
	}
}

func getTestedStorageExec(t *testing.T, returnedError error) postgres.Storage {
	mockDB := mocks.NewDB(t)
	mockDB.On("Exec",
		mock.Anything,
		testedItem.Name,
		testedItem.Description,
		testedItem.Count,
		testedItem.Cost,
		testedLogin,
	).
		Return(nil, returnedError).
		Once()

	storage := postgres.Postgres{
		Database: mockDB,
	}
	return storage
}

func TestDeleteItem(t *testing.T) {
	storage := getTestedStorageExec(t, nil)
	err := storage.DeleteItem(testedLogin, testedItem)
	if err != nil {
		t.Error(err)
	}
}

func TestErrorAddItem(t *testing.T) {
	storage := getTestedStorageExec(t, testedError)
	err := storage.AddItem(testedLogin, testedItem)
	if err == nil {
		t.Error("Expected error, got nil")
	}
}

func TestErrorDeleteItem(t *testing.T) {
	storage := getTestedStorageExec(t, testedError)
	err := storage.DeleteItem(testedLogin, testedItem)
	if err == nil {
		t.Error("Expected error, got nil")
	}
}
