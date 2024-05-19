update "departments"
set department_name = $2,
    department_phone = $3
where id = $1