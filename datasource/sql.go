package datasource

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/chipx/go-translator/internal"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"strings"
	"time"
)

func NewSqlDataSource(db *sqlx.DB, tableName string, updateLastModifiedPeriod time.Duration) DataSource {
	source := &sqlDatasource{
		db:        db,
		tableName: tableName,
	}

	if updateLastModifiedPeriod > 0 {
		go func() {
			t := time.NewTicker(updateLastModifiedPeriod)
			for {
				source.updateLastModified()
				<-t.C
			}
		}()
	}

	return source
}

const (
	defaultQueryLimit = 100
)

var ErrorNotFound = errors.New("Not found ")

type tableRow struct {
	Id         int
	Lang       string
	Key        string
	Message    string
	ModifiedAt time.Time `db:"modified_at"`
	Translated int
}

type sqlDatasource struct {
	db           *sqlx.DB
	tableName    string
	lastModified time.Time
}

func (s *sqlDatasource) LoadAll(criteria Criteria) (map[string]*internal.Vocabulary, error) {
	tx, err := s.db.Beginx()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	var rowData tableRow
	ctl := make(map[string]*internal.Vocabulary)

	neePagination := false
	if criteria.Limit < 1 {
		criteria.Limit = defaultQueryLimit
		criteria.Page = 0
		neePagination = true
	}

	queryString, queryParams := s.buildSearchQuery(criteria)

	for {
		rows, err := tx.NamedQuery(queryString, queryParams)
		if err != nil {
			return nil, err
		}

		if rows.Next() {
			for {
				err = rows.StructScan(&rowData)
				if err != nil {
					fmt.Println("Fail scan row: ", err)
					return nil, err
				}

				if _, ok := ctl[rowData.Lang]; !ok {
					ctl[rowData.Lang] = internal.NewVocabulary(rowData.Lang)
				}

				ctl[rowData.Lang].Set(rowData.Key, rowData.Message)

				if s.lastModified.Before(rowData.ModifiedAt) {
					s.lastModified = rowData.ModifiedAt
				}

				if !rows.Next() {
					break
				}
			}
		} else {
			break
		}

		if neePagination {
			criteria.Page += 1
			queryParams["offset"] = criteria.GetOffset()
		} else {
			break
		}
	}

	return ctl, nil
}

func (s *sqlDatasource) buildSearchQuery(criteria Criteria) (string, map[string]interface{}) {
	queryParams := make(map[string]interface{})
	var where []string

	if !criteria.Since.IsZero() {
		where = append(where, "modified_at > :since")
		queryParams["since"] = criteria.Since
	}

	if criteria.Translated > NOT_SET {
		where = append(where, "translated=:translated")
		queryParams["translated"] = criteria.Translated
	}

	if len(criteria.KeyPrefix) > 0 {
		where = append(where, "key LIKE :key_like")
		queryParams["key_like"] = criteria.KeyPrefix + "%"
	}

	if len(criteria.Langs) > 0 {
		var langIn []string
		for i := 0; i < len(criteria.Langs); i++ {
			key := fmt.Sprintf("lang_%d", i)
			langIn = append(langIn, ":"+key)
			queryParams[key] = criteria.Langs[i]
		}
		where = append(where, fmt.Sprintf("lang IN (%s)", strings.Join(langIn, ",")))
	}

	queryStr := fmt.Sprintf("SELECT id, lang, key, message, modified_at, translated FROM \"%s\"", s.tableName)
	if len(where) > 0 {
		queryStr += "\nWHERE " + strings.Join(where, " AND ")
	}

	if criteria.Limit > 0 {
		queryStr += fmt.Sprintf("\nLIMIT %d OFFSET :offset", criteria.Limit)
		queryParams["offset"] = criteria.GetOffset()
	}

	return queryStr, queryParams
}

func (s *sqlDatasource) updateLastModified() {
	row := s.db.QueryRow(fmt.Sprintf("SELECT MAX(modified_at) FROM \"%s\"", s.tableName))
	var timeResult sql.NullTime
	if err := row.Scan(&timeResult); err != nil {
		log.WithError(err).Error("Update last modified failed")
	}
	if timeResult.Valid {
		s.lastModified = timeResult.Time
	}
}

func (s *sqlDatasource) GetLastModified() time.Time {
	return s.lastModified
}
func (s *sqlDatasource) Get(lang string, key string) (msg string, err error) {
	rows, err := s.db.NamedQuery(fmt.Sprintf(
		"SELECT message FROM \"%s\" WHERE lang=:lang AND key=:key LIMIT 1",
		s.tableName,
	), map[string]interface{}{
		"lang": lang,
		"key":  key,
	},
	)

	if err != nil {
		return "", err
	}

	if !rows.Next() {
		return "", ErrorNotFound
	}

	rows.Scan(&msg)
	return
}

func (s *sqlDatasource) Set(lang string, key string, msg string) (err error) {
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

	if err := s.create(lang, key, msg, YES); err == nil {
		return nil
	}

	return s.update(lang, key, msg)
}

func (s *sqlDatasource) create(lang string, key string, msg string, translated Translated) error {
	modifyTime := time.Now()
	_, err := s.db.NamedExec(
		fmt.Sprintf("INSERT INTO \"%s\" (lang, key, message, modified_at, translated) VALUES (:lang, :key, :message, :modified_at, :translated)", s.tableName),
		map[string]interface{}{
			"lang":        lang,
			"key":         key,
			"message":     msg,
			"modified_at": modifyTime,
			"translated":  translated,
		},
	)

	if err == nil {
		s.lastModified = time.Now()
	}

	return err
}
func (s *sqlDatasource) update(lang string, key string, msg string) error {
	modifyTime := time.Now()
	_, err := s.db.NamedExec(
		fmt.Sprintf("UPDATE \"%s\" SET message=:message, modified_at=:modified_at, translated=:translated WHERE lang=:lang AND key=:key", s.tableName),
		map[string]interface{}{
			"lang":        lang,
			"key":         key,
			"message":     msg,
			"translated":  YES,
			"modified_at": modifyTime,
		},
	)

	if err == nil {
		s.lastModified = time.Now()
	}

	return err
}

func (s *sqlDatasource) Delete(lang string, key string) error {
	_, err := s.db.NamedExec(
		fmt.Sprintf("DELETE FROM \"%s\" WHERE lang=:lang AND key=:key", s.tableName),
		map[string]interface{}{
			"lang": lang,
			"key":  key,
		},
	)

	return err
}

func (s *sqlDatasource) MarkAsUntranslated(lang string, key string) (err error) {
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

	_, err = s.Get(lang, key)
	if err == nil {
		return
	}

	return s.create(lang, key, key, NO)
}
