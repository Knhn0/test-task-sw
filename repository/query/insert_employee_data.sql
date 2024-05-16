insert into "users" (user_name, surname, phone, company_id, passport_id, department_id)
values ($1, $2, $3, $4, $5, $6)
RETURNING id;