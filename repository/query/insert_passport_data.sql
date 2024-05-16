insert into "passports" (passport_type, passport_number_hash)
values ($1, $2)
RETURNING id;