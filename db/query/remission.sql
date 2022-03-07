-- name: FindAll :many
select * from reception;

-- name: FindByCustomer :many
select * from reception where customer = $1;

-- name: FindByEquipment :one
select * from reception where equipment = $1;

-- name: FindById :one
select * from reception where id = $1;

-- name: Insert :one
insert into reception (reason, observation, status, customer, equipment)
values ($1, $2, $3, $4, $5)
RETURNING *;

-- name: UpdateReasonAndObservation :one
update reception
set reason = $2 , observation = $3
where id = $1
RETURNING *;

-- name: UpdateDateDelivery :one
update reception
set date_delivery = $2, status = $3
where id = $1
RETURNING *;

-- name: UpdateStatus :one
update reception
set status = $2
where id = $1
RETURNING *;

-- name: Delete :exec
delete from reception where id = $1;
