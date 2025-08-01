// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package db

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.createMessageStmt, err = db.PrepareContext(ctx, createMessage); err != nil {
		return nil, fmt.Errorf("error preparing query CreateMessage: %w", err)
	}
	if q.createSessionStmt, err = db.PrepareContext(ctx, createSession); err != nil {
		return nil, fmt.Errorf("error preparing query CreateSession: %w", err)
	}
	if q.deleteMessageStmt, err = db.PrepareContext(ctx, deleteMessage); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteMessage: %w", err)
	}
	if q.deleteSessionStmt, err = db.PrepareContext(ctx, deleteSession); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteSession: %w", err)
	}
	if q.deleteSessionMessagesStmt, err = db.PrepareContext(ctx, deleteSessionMessages); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteSessionMessages: %w", err)
	}
	if q.getMessageStmt, err = db.PrepareContext(ctx, getMessage); err != nil {
		return nil, fmt.Errorf("error preparing query GetMessage: %w", err)
	}
	if q.getSessionByIDStmt, err = db.PrepareContext(ctx, getSessionByID); err != nil {
		return nil, fmt.Errorf("error preparing query GetSessionByID: %w", err)
	}
	if q.listMessagesBySessionStmt, err = db.PrepareContext(ctx, listMessagesBySession); err != nil {
		return nil, fmt.Errorf("error preparing query ListMessagesBySession: %w", err)
	}
	if q.listSessionsStmt, err = db.PrepareContext(ctx, listSessions); err != nil {
		return nil, fmt.Errorf("error preparing query ListSessions: %w", err)
	}
	if q.updateMessageStmt, err = db.PrepareContext(ctx, updateMessage); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateMessage: %w", err)
	}
	if q.updateSessionStmt, err = db.PrepareContext(ctx, updateSession); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateSession: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.createMessageStmt != nil {
		if cerr := q.createMessageStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createMessageStmt: %w", cerr)
		}
	}
	if q.createSessionStmt != nil {
		if cerr := q.createSessionStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createSessionStmt: %w", cerr)
		}
	}
	if q.deleteMessageStmt != nil {
		if cerr := q.deleteMessageStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteMessageStmt: %w", cerr)
		}
	}
	if q.deleteSessionStmt != nil {
		if cerr := q.deleteSessionStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteSessionStmt: %w", cerr)
		}
	}
	if q.deleteSessionMessagesStmt != nil {
		if cerr := q.deleteSessionMessagesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteSessionMessagesStmt: %w", cerr)
		}
	}
	if q.getMessageStmt != nil {
		if cerr := q.getMessageStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getMessageStmt: %w", cerr)
		}
	}
	if q.getSessionByIDStmt != nil {
		if cerr := q.getSessionByIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getSessionByIDStmt: %w", cerr)
		}
	}
	if q.listMessagesBySessionStmt != nil {
		if cerr := q.listMessagesBySessionStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listMessagesBySessionStmt: %w", cerr)
		}
	}
	if q.listSessionsStmt != nil {
		if cerr := q.listSessionsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listSessionsStmt: %w", cerr)
		}
	}
	if q.updateMessageStmt != nil {
		if cerr := q.updateMessageStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateMessageStmt: %w", cerr)
		}
	}
	if q.updateSessionStmt != nil {
		if cerr := q.updateSessionStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateSessionStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                        DBTX
	tx                        *sql.Tx
	createMessageStmt         *sql.Stmt
	createSessionStmt         *sql.Stmt
	deleteMessageStmt         *sql.Stmt
	deleteSessionStmt         *sql.Stmt
	deleteSessionMessagesStmt *sql.Stmt
	getMessageStmt            *sql.Stmt
	getSessionByIDStmt        *sql.Stmt
	listMessagesBySessionStmt *sql.Stmt
	listSessionsStmt          *sql.Stmt
	updateMessageStmt         *sql.Stmt
	updateSessionStmt         *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                        tx,
		tx:                        tx,
		createMessageStmt:         q.createMessageStmt,
		createSessionStmt:         q.createSessionStmt,
		deleteMessageStmt:         q.deleteMessageStmt,
		deleteSessionStmt:         q.deleteSessionStmt,
		deleteSessionMessagesStmt: q.deleteSessionMessagesStmt,
		getMessageStmt:            q.getMessageStmt,
		getSessionByIDStmt:        q.getSessionByIDStmt,
		listMessagesBySessionStmt: q.listMessagesBySessionStmt,
		listSessionsStmt:          q.listSessionsStmt,
		updateMessageStmt:         q.updateMessageStmt,
		updateSessionStmt:         q.updateSessionStmt,
	}
}
