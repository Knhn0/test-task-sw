update "passports"
set passport_type = $2,
    passport_number = $3
where id = $1