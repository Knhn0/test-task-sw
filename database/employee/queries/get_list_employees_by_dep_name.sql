select e.id, e.employee_name, e.surname, e.phone, e.company_id, p.passport_type, p.passport_number, d.department_name, d.department_phone
from employees as e
         left join passports as p on p.id = e.passport_id
         left join departments as d on d.id = e.department_id
where d.department_name = $1