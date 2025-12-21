package db_test

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/JingolBong/task-6/internal/db"
	"github.com/stretchr/testify/require"
)

const (
	nameQuery = "SELECT name FROM users"
)

var (
	errorQuery = errors.New("error query")
)

func TestDBGetNamesSuccess(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	row := sqlmock.NewRows([]string{"name"}).AddRow("Jingol")

	mock.ExpectQuery(nameQuery).WillReturnRows(row)

	service := db.New(mockDB)
	name, err := service.GetNames()
	require.Equal(t, []string{"Jingol"}, name)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestDBGetNamesErrors(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	mock.ExpectQuery(nameQuery).WillReturnError(errorQuery)

	service := db.New(mockDB)
	name, err := service.GetNames()
	require.Nil(t, name)
	require.ErrorContains(t, err, errorQuery.Error())
	require.NoError(t, mock.ExpectationsWereMet())
}
