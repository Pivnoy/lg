insert into "user"(uuid, email, password) values
    ('47e94c22-933d-4029-9752-dad3cd53a85b','bebra228@mail.ru', '$2a$14$GIVQqHS8j5BC8UUONZlB2.u5BLRXsneGZpssZ7tpU/Q7mfusIdiu.'),
    ('88ec4f5a-1989-4016-8684-8a9dd89d1d73', 'jopa@gmail.com', '$2a$14$RLs4FenQxxyjtIw3inpSkupQvlxJco8DUpIvScrZsyQc4AiCcOh.C'),
    ('6e07758f-bb73-49ed-a9d7-0ac4e32064f5','admin@admin.com', '$2a$14$zSXemWq1Eeqed5KnMKhXFur77DnjdWtMbKq4wTItvw33JGMqXLbu.'),
    ('753a2846-a6ef-4970-9693-41d66a55210f', 'roma@grek.chechnya', '$2a$14$PiedNAExYhO39Rm6dfZA.uUrlgijnuYNOFFkzApJZ7TJ4lPKfl3JG');

/*
    'bebra228@mail.ru', 'qwerty1234'
    'jopa@gmail.com', 'ASdf123121'
    'admin@admin.com', 'admin1_DFfsdf'
    'roma@grek.chechnya', 'ya_ne_chachAAA228
 */

insert into administrator(user_uuid, email, firstname, lastname, patronymic) values
    ('6e07758f-bb73-49ed-a9d7-0ac4e32064f5', 'admin@admin.com', 'Evgeniy', 'Compiler', 'GPUvich');

insert into country(uuid, "name", code) values
    ('e2680202-79ed-4aee-9c07-d51eda5cd08c', 'Российская Федерация', 'ru'),
    ('20872fbf-0b78-450e-85c4-10983fb9de98', 'Азербайджан', 'az'),
    ('d99c9ba2-5c51-41ff-8095-ccdc8d24ee71', 'Армения', 'am'),
    ('bbb0daf9-0a21-455c-a6e3-25a1d6d69044', 'Белоруссия', 'by'),
    ('e35880de-4459-48de-847c-64f22d9c5a75', 'Казахстан', 'kz'),
    ('970a4fa3-9b8b-4b7b-96eb-9efd24fb4f70', 'Киргизия', 'kg'),
    ('65a994db-4ff5-4647-8581-27fa3f89beeb', 'Молдавия', 'md'),
    ('5aae69e8-4705-40ff-888d-aec1f7fd9f0e', 'Таджикистан', 'tj'),
    ('d231da9e-0396-4eb8-b9ca-48178534d111', 'Узбекистан', 'uz');

insert into city(uuid, "name", country_uuid) values
    ('dc1d3489-6e4e-4852-af5b-97a716ec1c3e', 'Москва', 'e2680202-79ed-4aee-9c07-d51eda5cd08c'),
    ('441c4ff7-de3a-471a-bfce-748fc4551993', 'Санкт-Петербург', 'e2680202-79ed-4aee-9c07-d51eda5cd08c'),
    ('2d9576ed-8293-47ab-9fb2-43c1dad82880', 'Архангельск', 'e2680202-79ed-4aee-9c07-d51eda5cd08c'),
    ('ba3bddcb-4d76-45e7-9699-93d523b5736d', 'Котлас', 'e2680202-79ed-4aee-9c07-d51eda5cd08c'),
    ('861e7374-e11a-4d05-9dcf-d1652bebc393','Крым', 'e2680202-79ed-4aee-9c07-d51eda5cd08c'),
    ('c9951e29-0bc0-45a5-8daf-e48070370f54', 'Херсон', 'e2680202-79ed-4aee-9c07-d51eda5cd08c'),
    ('4c7323f6-0857-4222-91da-691bc232972b', 'Баку', '20872fbf-0b78-450e-85c4-10983fb9de98'),
    ('c735ef55-d74d-4257-ab9c-fc32a3d974dd', 'Ереван', 'd99c9ba2-5c51-41ff-8095-ccdc8d24ee71'),
    ('af3b6cf5-550d-44ef-983c-459bd714ad72', 'Минск', 'bbb0daf9-0a21-455c-a6e3-25a1d6d69044'),
    ('22bc3512-60cd-45b1-b2f5-f32647c9c0ee', 'Гомель', 'bbb0daf9-0a21-455c-a6e3-25a1d6d69044'),
    ('5771fe6f-8870-4c66-95eb-5a95535013fc', 'Астана', 'e35880de-4459-48de-847c-64f22d9c5a75'),
    ('d191550c-38ca-49a2-bb76-8fb4cef77ded', 'Бишкек', '970a4fa3-9b8b-4b7b-96eb-9efd24fb4f70'),
    ('f017fd23-b3c1-4ba2-8131-e9f64d92a1c1', 'Кишинев', '65a994db-4ff5-4647-8581-27fa3f89beeb'),
    ('9e9f2546-8ea0-4548-915b-09cfb6d1bd03', 'Душанбе', '5aae69e8-4705-40ff-888d-aec1f7fd9f0e'),
    ('02a65625-f5f8-4ac2-a9a6-3f33d7a5ae7c', 'Ташкент', 'd231da9e-0396-4eb8-b9ca-48178534d111');

insert into citizenship(uuid, "name", country_uuid) values
    ('8e0dff70-ae44-49e4-9b29-107fd08c9fe6', 'Российская Федерация', 'e2680202-79ed-4aee-9c07-d51eda5cd08c'),
    ('cb615029-0978-44a9-afb7-814b28b5c4e6', 'Азербайджан', '20872fbf-0b78-450e-85c4-10983fb9de98'),
    ('a8916bb6-14f7-4bbc-a718-315aaae499f7', 'Армения', 'd99c9ba2-5c51-41ff-8095-ccdc8d24ee71'),
    ('66fe4254-a426-45ed-b42b-a18965137868', 'Белоруссия', 'bbb0daf9-0a21-455c-a6e3-25a1d6d69044'),
    ('2319d698-558c-487e-8659-3d3fb4eac073', 'Казахстан', 'e35880de-4459-48de-847c-64f22d9c5a75'),
    ('705ff0e4-2994-42b5-94fe-30901ca81e5d', 'Киргизия', '970a4fa3-9b8b-4b7b-96eb-9efd24fb4f70'),
    ('b3f8c2ee-e68d-440f-95cc-0bc921fbbee0', 'Молдавия', '65a994db-4ff5-4647-8581-27fa3f89beeb'),
    ('c0c47d83-fe4b-4eaf-8e14-8a04b38c22d5', 'Таджикистан', '5aae69e8-4705-40ff-888d-aec1f7fd9f0e'),
    ('0fca4caa-39d1-44d8-bdd3-7837d928fa42', 'Узбекистан', 'd231da9e-0396-4eb8-b9ca-48178534d111');

insert into company(uuid, "name", inn) values
    ('e474bd33-e7df-4450-ae17-35b3fa71e1a0', 'ООО НАСТОЯЩАЯ РАБОТА', '7720440174'),
    ('01ee4013-6ada-409d-be4f-b7179cdc090a', 'ООО ЧКО', '7729577470'),
    ('6b97fb4d-2e16-425b-8a3e-50ab4bb666f5', 'ООО РОГА И КОПЫТА', '7736637017'),
    ('34ae926c-2ca4-4b6a-88cd-495066984876', 'ООО МОЯ ОБОРОНА', '0123456789'),
    ('a23dc5e9-c49f-4335-833e-de4129a3d437', 'ООО ЭС КАК ДОЛЛАР', '4632259118'),
    ('da567074-74fe-412b-b5d2-8131b9b78bf2', 'ООО ААА', '6165036747');

insert into skill_category(uuid, "name", "value") values
    ('1d346133-c774-4c86-9a20-449278ab71e4', 'IT-технологии', 'it'),
    ('f3c2eded-42fd-4c99-ab06-b3057b0540fb', 'Экономика', 'economy'),
    ('52941dcd-6c9a-4db7-9be9-d32aedf134b4', 'Менеджмент', 'management');

insert into skill(uuid, "name", "value", skill_category_uuid) values
    ('6286f604-231b-40c2-91f2-4cfff2534f61', 'Docker', 'docker', '1d346133-c774-4c86-9a20-449278ab71e4'),
    ('fed7184a-f1c6-463b-92c7-20fe14d9388b', 'Java', 'java', '1d346133-c774-4c86-9a20-449278ab71e4'),
    ('f3e2177c-1802-40ad-b710-1b25222f0cd2', 'ООП', 'oop', '1d346133-c774-4c86-9a20-449278ab71e4'),
    ('5b96bff6-a860-453a-a973-ef2535ba5cbc', 'SQL', 'sql', '1d346133-c774-4c86-9a20-449278ab71e4'),
    ('fa489014-e2de-493d-86ef-a3c26e3c2a69', 'Трейдинг', 'trading', 'f3c2eded-42fd-4c99-ab06-b3057b0540fb'),
    ('d2026134-ae4c-4a8e-a3bf-7331adacedac', 'Аналитика', 'analytics', 'f3c2eded-42fd-4c99-ab06-b3057b0540fb'),
    ('6357134e-4365-4a0b-a6ac-b2f24f7cdd2c', 'Тренинги', 'training', '52941dcd-6c9a-4db7-9be9-d32aedf134b4'),
    ('9d9969bc-a33f-4ad7-8d3b-8ef819fe7140', 'Тим билдинг', 'team_building', '52941dcd-6c9a-4db7-9be9-d32aedf134b4');

insert into specialization(uuid, "name", "value") values
    ('348fb263-966c-4aad-822c-004013db095f', 'Программирование', 'programming'),
    ('6287ff96-62c0-47a9-859b-c1ddd538433b', 'Маркетинг', 'marketing'),
    ('254543a3-c17e-4cf6-b033-702a12c8c188', 'Продажи', 'sales');

insert into "role"(uuid, "name", specialization_uuid) values
    ('b45bf4e7-8a3f-4afc-9ee9-be47d97de1e0', 'Бэкенд', '348fb263-966c-4aad-822c-004013db095f'),
    ('2bf83093-55fd-470d-bae4-ba85785f2dbb', 'Фронтенд', '348fb263-966c-4aad-822c-004013db095f'),
    ('8f780e31-152c-4e03-bf6d-5be7127ff223', 'Архитектор системы', '348fb263-966c-4aad-822c-004013db095f'),
    ('ffb2f566-ba66-43c4-a2de-01773c7c3ef3', 'Тим лид', '348fb263-966c-4aad-822c-004013db095f'),
    ('45ffa061-5a32-4ebc-acbd-ff1665657439', 'Проджект менеджер', '348fb263-966c-4aad-822c-004013db095f'),
    ('3ac81060-b2d7-4745-81c6-696605ed4dec', 'СММщик', '6287ff96-62c0-47a9-859b-c1ddd538433b'),
    ('e0e33f90-cff4-4b33-889e-2c08bcec1296', 'ТикТокер', '6287ff96-62c0-47a9-859b-c1ddd538433b'),
    ('271ab91c-5adc-4829-9bb5-7e5f42934be5', 'Креативщик', '6287ff96-62c0-47a9-859b-c1ddd538433b'),
    ('36835b74-9135-4ca1-ad6b-5719e2be503e', 'Продавец-консультант', '254543a3-c17e-4cf6-b033-702a12c8c188'),
    ('29730be3-162e-4b67-8436-61f339806487', 'Оценщик', '254543a3-c17e-4cf6-b033-702a12c8c188');

insert into team(uuid, "name", "value") values
    ('48c27232-f192-46cb-a76d-a7087c5ff084', 'Команда Акакдем', 'team_academ'),
    ('27401b51-05df-404d-952a-4d36940d6812', 'Фунфырики', 'funfyiriks');

insert into achievement(uuid, "text") values
                                          ('acd177f5-0a74-42bf-9f2f-41ea6f6fb90f', '[Раунд 1: Oxxxymiron]
Я здесь чисто по фану, поглумиться над слабым
Ты же вылез из мамы под мой дисс на Бабана
Обличительный пафос — это пшик против папы
Эти рифмы писал мне пьяный Крипл под спайсом
Ты смешной, слишком длинный, откровенно нескладный
У тебя телосложение, как у беременной цапли
Непропорционально, как твой хайп и твой вклад в рэп —
Не облако в штанах. Ты лишь мода, как клауд-рэп
На пару сезонов.'),
                                          ('be1db4bd-4511-49fd-b8a7-2def8d5afdc2', 'Реализовал RestAPI сервис на Go.
• Настроил репликацию и автоматические горячие бэкапы Postgres.
• Законтейнеризировал приложение с помощью Docker и написал скрипт для развертывания docker-compose.'),
                                          ('5a4861f8-3434-4e7b-b67f-975f96dd9dca', '• Спроектировал систему команд, разработал язык программирования, похожий по синтаксису на Assembler.
• Написал транслятор для языка программирования в свою систему команд.
• Разработал эмулятор процессора для запуска кода.');

insert into employment(uuid, "name", "value") values
    ('348eeb1d-e0a5-48bb-bbd0-008b42863416', 'По найму', 'employed'),
    ('3736959d-e352-4e95-8b5e-72dfa16fbc08', 'На себя', 'self_employed'),
    ('62e033e6-24e7-403f-b954-f03bed1a0fd9', 'Без работы', 'unemployed');

insert into eduspeciality(uuid, "name", code) values
    ('94f0a874-92f5-4ebc-9b77-2dbcd2f6b5e9', 'Информатика и вычислительная техника', '09.03.01'),
    ('decb179e-0e22-448c-9d71-fb32d05fc3ed', 'Прикладная математика и информатика', '01.03.02'),
    ('7526d255-8211-4fdc-8965-35f071018ca9', 'Экология и природопользование', '05.03.06'),
    ('0eb9a1a9-caad-4a45-9cde-d17f22e7b0dd', 'Информационные системы и технологии', '09.04.02'),
    ('f6027b24-82fd-4805-951a-61314edef9f5', 'Приборостроение', '12.03.01');

insert into university(uuid, "name", city_uuid) values
    ('e4b276d9-9d0a-460a-8121-ae5710e995a1', 'МГУ', 'dc1d3489-6e4e-4852-af5b-97a716ec1c3e'),
    ('5ac03f10-f7ba-4046-b6c0-16ce6b583628', 'Бауманка', 'dc1d3489-6e4e-4852-af5b-97a716ec1c3e'),
    ('b35963d2-4fd4-49da-b959-6afa31497c4a', 'ИТМО', '441c4ff7-de3a-471a-bfce-748fc4551993'),
    ('59c739de-d2e3-4a10-a5b6-3252847ea917', 'СПбГУ', '441c4ff7-de3a-471a-bfce-748fc4551993'),
    ('c1095894-ffe2-4441-a8d8-d8cd158731c5', 'СПбПМУ', '441c4ff7-de3a-471a-bfce-748fc4551993'),
    ('aad32909-3263-4d29-94a9-9d2724a1d0a1', 'СГМУ', '2d9576ed-8293-47ab-9fb2-43c1dad82880');

insert into profile(user_uuid, firstname, lastname, patronymic, country_uuid, city_uuid,
                    citizenship_uuid, gender, phone, email, university_uuid, eduspeciality_uuid,
                    graduation_year, employment_uuid, experience, achievement_uuid, team_uuid, specialization_uuid, company_uuid, creation_date) VALUES
    ('47e94c22-933d-4029-9752-dad3cd53a85b', 'Алексей', 'Куценко', 'Викторович', 'e2680202-79ed-4aee-9c07-d51eda5cd08c', '441c4ff7-de3a-471a-bfce-748fc4551993', '8e0dff70-ae44-49e4-9b29-107fd08c9fe6', 'male', '+79876543210', 'pochta-dlya-svyazi@mail.ru', 'b35963d2-4fd4-49da-b959-6afa31497c4a', '94f0a874-92f5-4ebc-9b77-2dbcd2f6b5e9', 2023, '348eeb1d-e0a5-48bb-bbd0-008b42863416', 2, 'acd177f5-0a74-42bf-9f2f-41ea6f6fb90f', '48c27232-f192-46cb-a76d-a7087c5ff084', '348fb263-966c-4aad-822c-004013db095f', '01ee4013-6ada-409d-be4f-b7179cdc090a', '2022-11-02 01:05:30'),
    ('88ec4f5a-1989-4016-8684-8a9dd89d1d73', 'Григорий', 'Валевин', 'Александрович', 'e2680202-79ed-4aee-9c07-d51eda5cd08c', 'ba3bddcb-4d76-45e7-9699-93d523b5736d', '8e0dff70-ae44-49e4-9b29-107fd08c9fe6', 'other', '+79995552288', 'jopa@gmail.com', 'b35963d2-4fd4-49da-b959-6afa31497c4a', '7526d255-8211-4fdc-8965-35f071018ca9', 2024, '62e033e6-24e7-403f-b954-f03bed1a0fd9', 0, 'be1db4bd-4511-49fd-b8a7-2def8d5afdc2', '27401b51-05df-404d-952a-4d36940d6812', '254543a3-c17e-4cf6-b033-702a12c8c188', '6b97fb4d-2e16-425b-8a3e-50ab4bb666f5', '2022-09-28 12:22:30'),
    ('753a2846-a6ef-4970-9693-41d66a55210f', 'Роман', 'Логинов', null, 'd99c9ba2-5c51-41ff-8095-ccdc8d24ee71', 'c9951e29-0bc0-45a5-8daf-e48070370f54', '8e0dff70-ae44-49e4-9b29-107fd08c9fe6', 'male', '+79212281488', 'roma@grek.chechnya', 'b35963d2-4fd4-49da-b959-6afa31497c4a', '94f0a874-92f5-4ebc-9b77-2dbcd2f6b5e9', 2023, '3736959d-e352-4e95-8b5e-72dfa16fbc08', 1, '5a4861f8-3434-4e7b-b67f-975f96dd9dca', '48c27232-f192-46cb-a76d-a7087c5ff084', '348fb263-966c-4aad-822c-004013db095f', '34ae926c-2ca4-4b6a-88cd-495066984876', '2020-08-13 20:37:15');

insert into profile_skill(profile_uuid, skill_uuid) values
    ('47e94c22-933d-4029-9752-dad3cd53a85b', '6286f604-231b-40c2-91f2-4cfff2534f61'),
    ('47e94c22-933d-4029-9752-dad3cd53a85b', 'f3e2177c-1802-40ad-b710-1b25222f0cd2'),
    ('47e94c22-933d-4029-9752-dad3cd53a85b', '5b96bff6-a860-453a-a973-ef2535ba5cbc'),
    ('47e94c22-933d-4029-9752-dad3cd53a85b', '9d9969bc-a33f-4ad7-8d3b-8ef819fe7140'),
    ('88ec4f5a-1989-4016-8684-8a9dd89d1d73', '9d9969bc-a33f-4ad7-8d3b-8ef819fe7140'),
    ('753a2846-a6ef-4970-9693-41d66a55210f', '6286f604-231b-40c2-91f2-4cfff2534f61'),
    ('753a2846-a6ef-4970-9693-41d66a55210f', 'fed7184a-f1c6-463b-92c7-20fe14d9388b'),
    ('753a2846-a6ef-4970-9693-41d66a55210f', 'f3e2177c-1802-40ad-b710-1b25222f0cd2'),
    ('753a2846-a6ef-4970-9693-41d66a55210f', 'fa489014-e2de-493d-86ef-a3c26e3c2a69'),
    ('753a2846-a6ef-4970-9693-41d66a55210f', '6357134e-4365-4a0b-a6ac-b2f24f7cdd2c'),
    ('753a2846-a6ef-4970-9693-41d66a55210f', '9d9969bc-a33f-4ad7-8d3b-8ef819fe7140');

insert into category(uuid, "name") values
    ('13dadff0-98f4-438b-a803-65b3c7e4d833', 'Инновационные IT-решения'),
    ('45e07768-da80-4681-a01a-80654e40059b', 'Маркетинг'),
    ('5ac10c61-7b1b-4347-9b08-872c9a765cc6', 'Робототехника'),
    ('9d3c9663-dae0-4b11-a79a-0825807599be', 'Маркетплейсы');

insert into project(uuid, "name", description, category_uuid, project_link, presentation_link, creator_uuid, is_visible, creation_date) VALUES
    ('e153b3a7-ca0c-4fc7-b500-efaef1bd73cd', 'Закрытие академов', 'Планируем жесточайше ботать чтобы закрыть все академы и долги', '13dadff0-98f4-438b-a803-65b3c7e4d833', 'https://docs.google.com/document/d/1yTMA8CKSxGP3-H_PPTIxDdm_TnUx87uTLiyKCIYPtik/edit', 'https://miro.com/app/board/uXjVPahvELM=/', '47e94c22-933d-4029-9752-dad3cd53a85b', 'visible', '2022-11-03 04:05:30'),
    ('fa044ee2-cebd-484d-9697-9f8b25646188', 'Попить пива', 'Предлагаю всем смачно попить пыва на выходных на ветеранах', '45e07768-da80-4681-a01a-80654e40059b', 'https://ru.wikipedia.org/wiki/%D0%9F%D0%B8%D0%B2%D0%BE', 'https://pptcloud.ru/obzh/mify-o-pive', '88ec4f5a-1989-4016-8684-8a9dd89d1d73', 'visible', '2022-10-10 10:50:15');

insert into lineup(team_uuid, role_uuid, profile_uuid, project_uuid) VALUES
    ('48c27232-f192-46cb-a76d-a7087c5ff084', '2bf83093-55fd-470d-bae4-ba85785f2dbb', '47e94c22-933d-4029-9752-dad3cd53a85b', 'e153b3a7-ca0c-4fc7-b500-efaef1bd73cd'),
    ('48c27232-f192-46cb-a76d-a7087c5ff084', 'b45bf4e7-8a3f-4afc-9ee9-be47d97de1e0', '753a2846-a6ef-4970-9693-41d66a55210f', 'e153b3a7-ca0c-4fc7-b500-efaef1bd73cd'),
    ('27401b51-05df-404d-952a-4d36940d6812', 'ffb2f566-ba66-43c4-a2de-01773c7c3ef3', '88ec4f5a-1989-4016-8684-8a9dd89d1d73', 'fa044ee2-cebd-484d-9697-9f8b25646188'),
    ('27401b51-05df-404d-952a-4d36940d6812', 'e0e33f90-cff4-4b33-889e-2c08bcec1296', '47e94c22-933d-4029-9752-dad3cd53a85b', 'fa044ee2-cebd-484d-9697-9f8b25646188');

insert into message(uuid, author_uuid, "content", created_at_time) values
    ('9f7ea035-69c6-46ad-8680-dc1fb2f5a4a5', '47e94c22-933d-4029-9752-dad3cd53a85b', 'прив как дела', '2022-11-03 19:00:00'),
    ('1de3c811-1848-4898-aea1-c2a028868671', '753a2846-a6ef-4970-9693-41d66a55210f', 'какать хочется пиздец', '2022-11-03 19:01:00'),
    ('ef18493a-afd3-4625-b863-7dafc10f4e35', '47e94c22-933d-4029-9752-dad3cd53a85b', 'я люблю красить кнопки а еще фронтить', '2022-11-03 19:02:00'),
    ('f216b077-4e79-46a8-aee9-cd9d6c624603', '753a2846-a6ef-4970-9693-41d66a55210f', 'санечка, брат, есть туалетка? моя кончилась, 3 рулона за 2 пары пиздосссс', '2022-11-03 19:02:35'),
    ('614b131c-8946-4bc1-a563-a39ff760e27d', '88ec4f5a-1989-4016-8684-8a9dd89d1d73', 'я купил себе чарон и теперь сижу его парю', '2022-11-03 19:03:00'),
    ('93df9b6a-509c-4510-bdb0-aeabe537f25f', '753a2846-a6ef-4970-9693-41d66a55210f', 'женечек мой черный брат как дела снюсик будеш', '2022-11-03 20:10:00'),
    ('8a746954-04ec-4963-b3ba-81827b59e60e', '6e07758f-bb73-49ed-a9d7-0ac4e32064f5', 'брат сильно тебе бабаха шлепандухнула', '2022-11-03 20:15:00');

insert into chat(user_uuid, message_uuid) values
    ('{47e94c22-933d-4029-9752-dad3cd53a85b, 753a2846-a6ef-4970-9693-41d66a55210f, 88ec4f5a-1989-4016-8684-8a9dd89d1d73}', '{9f7ea035-69c6-46ad-8680-dc1fb2f5a4a5, 1de3c811-1848-4898-aea1-c2a028868671, ef18493a-afd3-4625-b863-7dafc10f4e35, f216b077-4e79-46a8-aee9-cd9d6c624603, 614b131c-8946-4bc1-a563-a39ff760e27d}');



select * from chat;

select message."content", chat.user_uuid from chat
inner join message on chat.message_uuid = message.uuid;

