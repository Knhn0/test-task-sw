insert into "passports" (passport_type, passport_number)
values ($1, $2)
RETURNING id;