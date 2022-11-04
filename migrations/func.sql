/*
 на PLpgSQL
приходят внешние ключи из таблицы profile
надо в этой функции проверить, что эти сущности существуют
if все ключи есть, возвращаю строку ok
если чего-то нет, то вовзращаю название сущности, которой нет

 */

create or replace function check_fk_profile(user_uuid_check uuid,
                                            country_uuid_check uuid,
                                            city_uuid_check uuid,
                                            citizenship_uuid_check uuid,
                                            university_uuid_check uuid,
                                            eduspeciality_uuid_check uuid,
                                            employment_uuid_check uuid,
                                            achievement_uuid_check uuid,
                                            team_uuid_check uuid,
                                            specialization_uuid_check uuid,
                                            company_uuid_check uuid)
    returns uuid as $$
begin
    if not exists (select * from "user" where uuid = user_uuid_check)
        then
        return user_uuid_check;
end if;
    if not exists (select * from country where uuid = country_uuid_check)
        then
        return country_uuid_check;
end if;
    if not exists (select * from city where uuid = city_uuid_check)
        then
        return city_uuid_check;
end if;
    if not exists(select * from citizenship where uuid = citizenship_uuid_check)
        then
        return citizenship_uuid_check;
end if;
    if not exists(select * from university where uuid = university_uuid_check)
        then
        return university_uuid_check;
end if;
    if not exists(select * from eduspeciality where uuid = eduspeciality_uuid_check)
        then
        return eduspeciality_uuid_check;
end if;
    if not exists(select * from employment where uuid = employment_uuid_check)
        then
        return employment_uuid_check;
end if;
    if not exists(select * from achievement where uuid = achievement_uuid_check)
        then
        return achievement_uuid_check;
end if;
    if not exists(select * from team where uuid = team_uuid_check)
        then
        return team_uuid_check;
end if;
    if not exists(select * from specialization where uuid = specialization_uuid_check)
        then
        return specialization_uuid_check;
end if;
    if not exists(select * from company where uuid = company_uuid_check)
        then
        return company_uuid_check;
end if;
return uuid_nil();
end;
    $$ language 'plpgsql';

create or replace function check_fk_lineup(team_uuid_check uuid,
                                            role_uuid_check uuid,
                                            profile_uuid_check uuid,
                                            project_uuid_check uuid)
    returns uuid as $$
begin
    if not exists(select * from team where uuid = team_uuid_check)
        then
        return team_uuid_check;
end if;
    if not exists(select * from "role" where uuid = role_uuid_check)
        then
        return role_uuid_check;
end if;
    if not exists(select * from profile where uuid = profile_uuid_check)
        then
        return profile_uuid_check;
end if;
    if not exists(select * from project where uuid = project_uuid_check)
        then
        return project_uuid_check;
end if;
return uuid_nil();
end;
$$ language 'plpgsql';