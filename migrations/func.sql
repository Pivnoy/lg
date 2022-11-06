create or replace function check_fk_profile(user_uuid_check uuid,
                                            country_uuid_check uuid,
                                            city_uuid_check uuid,
                                            citizenship_uuid_check uuid,
                                            university_uuid_check uuid,
                                            eduspeciality_uuid_check uuid,
                                            employment_uuid_check uuid,
                                            team_uuid_check uuid,
                                            specialization_uuid_check uuid)

    returns text as $$
begin
    if not exists (select * from "user" where uuid = user_uuid_check)
    then
        return 'user uuid';
    end if;
    if not exists (select * from country where uuid = country_uuid_check)
    then
        return 'country uuid';
    end if;
    if not exists (select * from city where uuid = city_uuid_check)
    then
        return 'city uuid';
    end if;
    if not exists(select * from citizenship where uuid = citizenship_uuid_check)
    then
        return 'citizenship uuid';
    end if;
    if university_uuid_check != uuid_nil() and not exists(select * from university where uuid = university_uuid_check)
    then
        return 'university uuid';   -- null
    end if;
    if eduspeciality_uuid_check != uuid_nil() and not exists(select * from eduspeciality where uuid = eduspeciality_uuid_check)
    then
        return 'eduspeciality uuid'; -- null
    end if;
    if not exists(select * from employment where uuid = employment_uuid_check)
    then
        return 'employment uuid';
    end if;
    if team_uuid_check != uuid_nil() and not exists(select * from team where uuid = team_uuid_check)
    then
        return 'team uuid';     -- null
    end if;
    if not exists(select * from specialization where uuid = specialization_uuid_check)
    then
        return 'specialization uuid';
    end if;
    return 'ok';
end
$$ language 'plpgsql';

-- create or replace function check_fk_lineup(team_uuid_check uuid,
--                                            role_uuid_check uuid,
--                                            profile_uuid_check uuid,
--                                            project_uuid_check uuid)
--     returns uuid as $$
-- begin
--     if not exists(select * from team where uuid = team_uuid_check)
--     then
--         return team_uuid_check;
--     end if;
--     if not exists(select * from "role" where uuid = role_uuid_check)
--     then
--         return role_uuid_check;
--     end if;
--     if not exists(select * from profile where uuid = profile_uuid_check)
--     then
--         return profile_uuid_check;
--     end if;
--     if not exists(select * from project where uuid = project_uuid_check)
--     then
--         return project_uuid_check;
--     end if;
--     return uuid_nil();
-- end;
-- $$ language 'plpgsql';
