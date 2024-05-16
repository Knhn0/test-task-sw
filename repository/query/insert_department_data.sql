insert into "departments" (department_name, department_phone)
values ($1, $2)
RETURNING id;