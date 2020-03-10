package datasource

import (
	"database/sql"
	"fmt"
	"github.com/chipx/go-transaltor/internal"
	"time"
)

func NewSqlDataSource(db *sql.DB, tableName string) *Sql {
	return &Sql{
		db:        db,
		tableName: tableName,
	}
}

const (
	queryLimit = 100
)

type tableRow struct {
	id         int
	lang       string
	key        string
	value      string
	modifiedAt time.Time
}

type Sql struct {
	db           *sql.DB
	tableName    string
	lastModified time.Time
}

func (s *Sql) LoadAll() (map[string]*internal.Vocabulary, error) {
	return s.LoadLast(time.Unix(0, 0))
}

func (s *Sql) LoadLast(since time.Time) (map[string]*internal.Vocabulary, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()
	var rowData tableRow
	ctl := make(map[string]*internal.Vocabulary)
	offset := 0

	for {
		rows, err := tx.Query(
			fmt.Sprintf(
				"SELECT id, lang, key, message, modified_at FROM \"%s\" WHERE modified_at>$1 LIMIT %d OFFSET %d",
				s.tableName,
				queryLimit,
				offset,
			),
			since,
		)
		if err != nil {
			return nil, err
		}

		if rows.Next() {
			for {
				err = rows.Scan(&rowData.id, &rowData.lang, &rowData.key, &rowData.value, &rowData.modifiedAt)
				if err != nil {
					fmt.Println("Fail scan row: ", err)
					continue
				}

				if _, ok := ctl[rowData.lang]; !ok {
					ctl[rowData.lang] = internal.NewVocabulary(rowData.lang)
				}

				ctl[rowData.lang].Set(rowData.key, rowData.value)

				if s.lastModified.Before(rowData.modifiedAt) {
					s.lastModified = rowData.modifiedAt
				}

				if !rows.Next() {
					break
				}
			}
		} else {
			break
		}

		offset += queryLimit
	}

	return ctl, nil
}
func (s *Sql) GetLastModified() time.Time {
	return s.lastModified
}
func (s *Sql) Get(lang string, key string) (msg string, err error) {
	err = s.db.QueryRow(fmt.Sprintf(
		"SELECT message FROM \"%s\" WHERE lang=$1 AND key=$2 LIMIT 1",
		s.tableName,
	),
		lang,
		key,
	).Scan(&msg)

	return
}

func (s *Sql) Set(lang string, key string, msg string) (err error) {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	if err := s.create(lang, key, msg); err == nil {
		return nil
	}

	return s.update(lang, key, msg)
}

func (s *Sql) create(lang string, key string, msg string) error {
	modifyTime := time.Now()
	_, err := s.db.Exec(
		fmt.Sprintf("INSERT INTO \"%s\" (lang, key, message, modified_at) VALUES ($1, $2, $3, $4)", s.tableName),
		lang,
		key,
		msg,
		modifyTime,
	)

	if err == nil {
		s.lastModified = time.Now()
	}

	return err
}
func (s *Sql) update(lang string, key string, msg string) error {
	modifyTime := time.Now()
	_, err := s.db.Exec(
		fmt.Sprintf("UPDATE \"%s\" SET message=$1, modified_at=$2 WHERE lang=$3 AND key=$4 LIMIT 1", s.tableName),
		msg,
		modifyTime,
		lang,
		key,
	)

	if err == nil {
		s.lastModified = time.Now()
	}

	return err
}

func (s *Sql) Delete(lang string, key string) error {
	_, err := s.db.Exec(
		fmt.Sprintf("DELETE FROM \"%s\" WHERE lang=$3 AND key=$4 LIMIT 1", s.tableName),
		lang,
		key,
	)

	return err
}

func (s *Sql) MarkAsUntranslated(lang string, key string) (err error) {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	modifyTime := time.Now()
	_, err = s.Get(lang, key)
	if err != nil {
		_, err = s.db.Exec(
			fmt.Sprintf("INSERT INTO \"%s\" (lang, key, message, modified_at, untranslated) VALUES ($1, $2, $3, $4, 1)", s.tableName),
			lang,
			key,
			key,
			modifyTime,
		)
		return
	}

	_, err = s.db.Exec(
		fmt.Sprintf("UPDATE \"%s\" SET modified_at=$1, untranslated=1 WHERE lang=$2 AND key=$3 LIMIT 1", s.tableName),
		modifyTime,
		lang,
		key,
	)
	return
}
