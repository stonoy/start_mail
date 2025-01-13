-- name: CreateMail :one
insert into email(id, created_at, updated_at, subject, body, sender, reciever)
values ($1, $2, $3, $4, $5, $6, $7)
returning *;

-- name: GetMailById :one
select * from email where id = $1 and (sender = $2 or reciever = $3);

-- name: InboxMailWithFilter :many
select * from email
where reciever = $1 and body like $2 and created_at between $3 and $4
order by updated_at desc
limit $5
offset $6;

-- name: NumOfInboxMailWithFilter :one
select count(*) from email
where reciever = $1 and body like $2 and created_at between $3 and $4;

-- name: SentBoxMailWithFilter :many
select * from email
where sender = $1 and body like $2 and created_at between $3 and $4
order by updated_at desc
limit $5
offset $6;

-- name: NumOfSentBoxMailWithFilter :one
select count(*) from email
where sender = $1 and body like $2 and created_at between $3 and $4;