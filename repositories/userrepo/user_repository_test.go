package userrepo

import (
	"database/sql/driver"
	"errors"
	"log"
	"regexp"
	"testing"
	"time"

	"github.com/necmettindev/currency-conversion/models/user"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func setupDB() (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("can't create sqlmock: %s", err)
	}

	gormDB, gerr := gorm.Open("postgres", db)
	if gerr != nil {
		log.Fatalf("can't open gorm connection: %s", err)
	}
	gormDB.LogMode(true)
	return gormDB, mock
}

type AnyTime struct{}

func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

func TestGetByUsername(t *testing.T) {
	gormDB, mock := setupDB()
	defer gormDB.Close()

	t.Run("GetByUsername", func(t *testing.T) {
		expected := &user.User{
			Username: "necmettin",
		}

		u := NewUserRepo(gormDB)
		sqlStr := `SELECT * FROM "users" WHERE "users"."deleted_at" IS NULL AND ((username = $1)) ORDER BY "users"."id" ASC LIMIT 1`

		mock.
			ExpectQuery(regexp.QuoteMeta(sqlStr)).
			WithArgs("necmettin").
			WillReturnRows(
				sqlmock.NewRows([]string{"username"}).
					AddRow("necmettin"))

		result, err := u.GetByUsername("necmettin")

		assert.EqualValues(t, expected, result)
		assert.Nil(t, err)
	})

	t.Run("GetByUsernameError", func(t *testing.T) {
		expected := errors.New("Nop")

		u := NewUserRepo(gormDB)
		sqlStr := `SELECT * FROM "users" WHERE "users"."deleted_at" IS NULL AND ((username = $1)) ORDER BY "users"."id" ASC LIMIT 1`

		mock.
			ExpectQuery(regexp.QuoteMeta(sqlStr)).
			WithArgs("necmettin").
			WillReturnError(expected)

		result, err := u.GetByUsername("necmettin")

		assert.EqualValues(t, expected, err)
		assert.Nil(t, result)
	})

	t.Run("GetByUsernameNotFound", func(t *testing.T) {
		expected := errors.New("record not found")

		u := NewUserRepo(gormDB)
		sqlStr := `SELECT * FROM "users" WHERE "users"."deleted_at" IS NULL AND ((username = $1)) ORDER BY "users"."id" ASC LIMIT 1`

		mock.
			ExpectQuery(regexp.QuoteMeta(sqlStr)).
			WithArgs("necmettin").
			WillReturnRows(
				sqlmock.NewRows([]string{}))

		result, err := u.GetByUsername("necmettin")

		assert.EqualValues(t, expected, err)
		assert.Nil(t, result)
	})
}

func TestCreate(t *testing.T) {
	gormDB, mock := setupDB()
	defer gormDB.Close()

	t.Run("Create a user", func(t *testing.T) {
		user := &user.User{
			Username: "necmettin",
			Password: "abc",
		}

		u := NewUserRepo(gormDB)

		mock.ExpectBegin()

		mock.
			ExpectQuery(
				regexp.QuoteMeta(
					`INSERT INTO "users" ("created_at","updated_at","deleted_at","first_name","last_name","username","password") VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING "users"."id"`)).
			WithArgs(AnyTime{}, AnyTime{}, nil, "", "", "necmettin", "abc").
			WillReturnRows(
				sqlmock.NewRows([]string{"id"}).
					AddRow(1))

		mock.ExpectCommit()

		err := u.Create(user)
		assert.Nil(t, err)
	})

	t.Run("Create a user fails", func(t *testing.T) {
		exp := errors.New("oops")
		user := &user.User{
			Username: "necmettin",
			Password: "abc",
		}

		u := NewUserRepo(gormDB)

		mock.ExpectBegin()

		mock.
			ExpectQuery(
				regexp.QuoteMeta(
					`INSERT INTO "users" ("created_at","updated_at","deleted_at","first_name","last_name","username","password") VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING "users"."id"`)).
			WithArgs(AnyTime{}, AnyTime{}, nil, "", "", "necmettin", "abc").
			WillReturnError(exp)

		mock.ExpectCommit()

		err := u.Create(user)
		assert.NotNil(t, err)
		assert.EqualValues(t, exp, err)
	})
}
