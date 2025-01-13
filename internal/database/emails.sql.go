// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: emails.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createMail = `-- name: CreateMail :one
insert into email(id, created_at, updated_at, subject, body, sender, reciever)
values ($1, $2, $3, $4, $5, $6, $7)
returning id, created_at, updated_at, subject, body, sender, reciever
`

type CreateMailParams struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Subject   string
	Body      string
	Sender    uuid.UUID
	Reciever  uuid.UUID
}

func (q *Queries) CreateMail(ctx context.Context, arg CreateMailParams) (Email, error) {
	row := q.db.QueryRowContext(ctx, createMail,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Subject,
		arg.Body,
		arg.Sender,
		arg.Reciever,
	)
	var i Email
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Subject,
		&i.Body,
		&i.Sender,
		&i.Reciever,
	)
	return i, err
}

const getMailById = `-- name: GetMailById :one
select id, created_at, updated_at, subject, body, sender, reciever from email where id = $1 and (sender = $2 or reciever = $3)
`

type GetMailByIdParams struct {
	ID       uuid.UUID
	Sender   uuid.UUID
	Reciever uuid.UUID
}

func (q *Queries) GetMailById(ctx context.Context, arg GetMailByIdParams) (Email, error) {
	row := q.db.QueryRowContext(ctx, getMailById, arg.ID, arg.Sender, arg.Reciever)
	var i Email
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Subject,
		&i.Body,
		&i.Sender,
		&i.Reciever,
	)
	return i, err
}

const inboxMailWithFilter = `-- name: InboxMailWithFilter :many
select id, created_at, updated_at, subject, body, sender, reciever from email
where reciever = $1 and body like $2 and created_at between $3 and $4
order by updated_at desc
limit $5
offset $6
`

type InboxMailWithFilterParams struct {
	Reciever    uuid.UUID
	Body        string
	CreatedAt   time.Time
	CreatedAt_2 time.Time
	Limit       int32
	Offset      int32
}

func (q *Queries) InboxMailWithFilter(ctx context.Context, arg InboxMailWithFilterParams) ([]Email, error) {
	rows, err := q.db.QueryContext(ctx, inboxMailWithFilter,
		arg.Reciever,
		arg.Body,
		arg.CreatedAt,
		arg.CreatedAt_2,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Email
	for rows.Next() {
		var i Email
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Subject,
			&i.Body,
			&i.Sender,
			&i.Reciever,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const numOfInboxMailWithFilter = `-- name: NumOfInboxMailWithFilter :one
select count(*) from email
where reciever = $1 and body like $2 and created_at between $3 and $4
`

type NumOfInboxMailWithFilterParams struct {
	Reciever    uuid.UUID
	Body        string
	CreatedAt   time.Time
	CreatedAt_2 time.Time
}

func (q *Queries) NumOfInboxMailWithFilter(ctx context.Context, arg NumOfInboxMailWithFilterParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, numOfInboxMailWithFilter,
		arg.Reciever,
		arg.Body,
		arg.CreatedAt,
		arg.CreatedAt_2,
	)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const numOfSentBoxMailWithFilter = `-- name: NumOfSentBoxMailWithFilter :one
select count(*) from email
where sender = $1 and body like $2 and created_at between $3 and $4
`

type NumOfSentBoxMailWithFilterParams struct {
	Sender      uuid.UUID
	Body        string
	CreatedAt   time.Time
	CreatedAt_2 time.Time
}

func (q *Queries) NumOfSentBoxMailWithFilter(ctx context.Context, arg NumOfSentBoxMailWithFilterParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, numOfSentBoxMailWithFilter,
		arg.Sender,
		arg.Body,
		arg.CreatedAt,
		arg.CreatedAt_2,
	)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const sentBoxMailWithFilter = `-- name: SentBoxMailWithFilter :many
select id, created_at, updated_at, subject, body, sender, reciever from email
where sender = $1 and body like $2 and created_at between $3 and $4
order by updated_at desc
limit $5
offset $6
`

type SentBoxMailWithFilterParams struct {
	Sender      uuid.UUID
	Body        string
	CreatedAt   time.Time
	CreatedAt_2 time.Time
	Limit       int32
	Offset      int32
}

func (q *Queries) SentBoxMailWithFilter(ctx context.Context, arg SentBoxMailWithFilterParams) ([]Email, error) {
	rows, err := q.db.QueryContext(ctx, sentBoxMailWithFilter,
		arg.Sender,
		arg.Body,
		arg.CreatedAt,
		arg.CreatedAt_2,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Email
	for rows.Next() {
		var i Email
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Subject,
			&i.Body,
			&i.Sender,
			&i.Reciever,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}