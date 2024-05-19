update "employees"
set employee_name = $2,
    surname       = $3,
    phone         = $4,
    company_id    = $5
where id = $1

returning passport_id, department_id;