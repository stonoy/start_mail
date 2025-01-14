-- name: CreateMail :one
insert into email(id, created_at, updated_at, subject, body, sender, reciever)
values ($1, $2, $3, $4, $5, $6, $7)
returning *;

-- name: GetMailById :one
select e.*, u1.email, u2.email from email e
inner join users u1 on e.sender = u1.id
inner join users u2 on e.reciever = u2.id
where e.id = $1 and (e.sender = $2 or e.reciever = $3);

-- name: InboxMailWithFilter :many
select email.*, users.email from email
inner join users on email.sender = users.id
where email.reciever = $1 and email.body like $2 and email.created_at between $3 and $4
order by email.updated_at desc
limit $5
offset $6;

-- name: NumOfInboxMailWithFilter :one
select count(*) from email
where reciever = $1 and body like $2 and created_at between $3 and $4;

-- name: SentBoxMailWithFilter :many
select email.*, users.email from email
inner join users on email.reciever = users.id
where email.sender = $1 and email.body like $2 and email.created_at between $3 and $4
order by email.updated_at desc
limit $5
offset $6;

-- name: NumOfSentBoxMailWithFilter :one
select count(*) from email
where sender = $1 and body like $2 and created_at between $3 and $4;

-- name: GetInboxCount :one
select count(*) from email where reciever = $1;

-- name: GetSentBoxCount :one
select count(*) from email where sender = $1;