// Package model contains the types for schema 'gwa'.
package model

// Code generated by xo. DO NOT EDIT.

import (
	"errors"
	"time"
)

// Task represents a row from 'gwa.task'.
type Task struct {
	ID        uint      `json:"id"`         // id
	CreatedAt time.Time `json:"created_at"` // created_at
	UpdatedAt time.Time `json:"updated_at"` // updated_at
	Title     string    `json:"title"`      // title

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Task exists in the database.
func (t *Task) Exists() bool {
	return t._exists
}

// Deleted provides information if the Task has been deleted from the database.
func (t *Task) Deleted() bool {
	return t._deleted
}

// Insert inserts the Task to the database.
func (t *Task) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if t._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key must be provided
	const sqlstr = `INSERT INTO gwa.task (` +
		`id, created_at, updated_at, title` +
		`) VALUES (` +
		`?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, t.ID, t.CreatedAt, t.UpdatedAt, t.Title)
	_, err = db.Exec(sqlstr, t.ID, t.CreatedAt, t.UpdatedAt, t.Title)
	if err != nil {
		return err
	}

	// set existence
	t._exists = true

	return nil
}

// Update updates the Task in the database.
func (t *Task) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !t._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if t._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE gwa.task SET ` +
		`created_at = ?, updated_at = ?, title = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, t.CreatedAt, t.UpdatedAt, t.Title, t.ID)
	_, err = db.Exec(sqlstr, t.CreatedAt, t.UpdatedAt, t.Title, t.ID)
	return err
}

// Save saves the Task to the database.
func (t *Task) Save(db XODB) error {
	if t.Exists() {
		return t.Update(db)
	}

	return t.Insert(db)
}

// Delete deletes the Task from the database.
func (t *Task) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !t._exists {
		return nil
	}

	// if deleted, bail
	if t._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM gwa.task WHERE id = ?`

	// run query
	XOLog(sqlstr, t.ID)
	_, err = db.Exec(sqlstr, t.ID)
	if err != nil {
		return err
	}

	// set deleted
	t._deleted = true

	return nil
}

// TaskByID retrieves a row from 'gwa.task' as a Task.
//
// Generated from index 'task_id_pkey'.
func TaskByID(db XODB, id uint) (*Task, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, created_at, updated_at, title ` +
		`FROM gwa.task ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	t := Task{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&t.ID, &t.CreatedAt, &t.UpdatedAt, &t.Title)
	if err != nil {
		return nil, err
	}

	return &t, nil
}
