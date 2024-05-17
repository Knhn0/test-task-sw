select e.id, e.employee_name, e.surname, e.phone, e.company_id, e.passport_id, e.department_id
from employees as e
where e.id = $1;