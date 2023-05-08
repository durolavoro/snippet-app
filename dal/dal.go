package dal

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/durolavoro/snippet-app/model"
)

const (
	INSERT_INTO_SNIPPETS = "INSERT INTO snippets (title, content, created, expires) " +
		"VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))"
	GET_SNIPPET  = "SELECT * FROM snippets WHERE id=? AND expires > UTC_TIMESTAMP()"
	GET_SNIPPETS = "SELECT id, title, content, created, expires FROM snippets " +
		"WHERE expires > UTC_TIMESTAMP() ORDER BY created DESC LIMIT ?"
	DELETE_FROM_SNIPPETS = "DELETE FROM snippets WHERE ID=?"
)

// MB Database/DAL related stuff

type DAL struct {
	DB *sql.DB
}

func CreateDAL(db *sql.DB) *DAL {
	return &DAL{
		DB: db,
	}
}

func (d *DAL) InsertError(title, _, _ string) (int, error) {
	return 0, fmt.Errorf("error inserting into DB %s", title)
}

func (d *DAL) Insert(title, content, expires string) (int, error) {
	_, err := d.DB.Exec(INSERT_INTO_SNIPPETS, title, content, expires)
	if err != nil {
		return 1, fmt.Errorf("error inserting into DB %w", err)
	}
	return 0, nil
}

func (d *DAL) Get(id int) (*model.Snippet, error) {
	row := d.DB.QueryRow(GET_SNIPPET, id)
	if row == nil {
		return nil, fmt.Errorf("error querying db for id %d", id)
	}
	snip := &model.Snippet{}
	err := row.Scan(&snip.ID, &snip.Title, &snip.Content, &snip.Created, &snip.Expires)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, model.ErrNoRecord
	} else if err != nil {
		return nil, err
	}
	return snip, nil
}

func (d *DAL) Delete(id int) (int, error) {
	_, err := d.DB.Exec(DELETE_FROM_SNIPPETS, id)
	if err != nil {
		return 1, fmt.Errorf("error deleting %d from DB: %w", id, err)
	}
	return 0, nil
}

func (d *DAL) Latest(limit int) ([]*model.Snippet, error) {
	rows, err := d.DB.Query(GET_SNIPPETS, limit)
	if err != nil {
		return nil, fmt.Errorf("error querying latest %d snippets: %w", limit, err)
	}
	var snips []*model.Snippet
	for rows.Next() {
		snip := &model.Snippet{}
		err := rows.Scan(&snip.ID, &snip.Title, &snip.Content, &snip.Created, &snip.Expires)
		if err != nil {
			return nil, fmt.Errorf("error scanning snip: %w", err)
		}
		snips = append(snips, snip)
	}
	return snips, nil
}
