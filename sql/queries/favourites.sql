-- name: AddToFavourites :one
insert into favourite(id, created_at, updated_at, userid, emailid)
values ($1, $2, $3, $4, $5)
returning *;

-- name: GetAllFavOfUser :many
select f.id, f.userid, f.created_at, e.* from favourite f
inner join email e
on f.emailid = e.id
where f.userid = $1
order by f.created_at desc
limit $2
offset $3;

-- name: NumOfAllFavUser :one
select count(*) from favourite f
inner join email e
on f.emailid = e.id
where f.userid = $1;

-- name: DeleteFav :one
delete from favourite where id = $1 and userid = $2
returning *;

-- name: GetNumFav :one
select count(*) from favourite where userid = $1;