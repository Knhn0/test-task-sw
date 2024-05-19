delete from employees as e
where e.id = $1
returning passport_id;