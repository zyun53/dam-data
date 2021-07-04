CREATE SEQUENCE nl_dam_institutions INCREMENT 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1;

CREATE TABLE IF NOT EXISTS "nl_dam_institutions" (
    "id" bigint DEFAULT nextval('dams_id_seq') NOT NULL,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    "code" int not NULL,
    "value" text NOT NULL
);

CREATE SEQUENCE nl_dam_precisions INCREMENT 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1;

CREATE TABLE IF NOT EXISTS "nl_dam_precisions" (
    "id" bigint DEFAULT nextval('dams_id_seq') NOT NULL,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    "code" int not NULL,
    "value" text NOT NULL
);

CREATE SEQUENCE nl_dam_purposes INCREMENT 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1;

CREATE TABLE IF NOT EXISTS "nl_dam_purposes" (
    "id" bigint DEFAULT nextval('dams_id_seq') NOT NULL,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    "code" int not NULL,
    "value" text NOT NULL
);

CREATE SEQUENCE nl_dam_types INCREMENT 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1;

CREATE TABLE IF NOT EXISTS "nl_dam_types" (
    "id" bigint DEFAULT nextval('dams_id_seq') NOT NULL,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    "code" int not NULL,
    "value" text NOT NULL
);

CREATE SEQUENCE nl_dams_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1;

CREATE TABLE IF NOT EXISTS "nl_dams" (
    "id" bigint DEFAULT nextval('dams_id_seq') NOT NULL,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    "position_lat" float8,
    "position_lng" float8,
    "name" text,
    "water_system_name" text,
    "river_name" text,
    "code" int primary key,
    "type" int references nl_dam_types(id),
    "purpose" int references nl_dam_purposes(id),
    "dam_scale_bank_height" real,
    "dam_scale_bank_span" real,
    "bank_volume" real,
    "total_pondage" real,
    "institution_in_charge" int references nl_dam_institutions(id),
    "year_of_completion" int,
    "address" text,
    "positional_information_precision" int references nl_dam_precisions(id)
);

CREATE TABLE IF NOT EXISTS "nl_dam_nl_purposes" (
    "nl_dams_id" int NOT NULL references nl_dams(id),
    "nl_dam_purposes_id" int NOT NULL references nl_dam_purposes(id)
);

INSERT INTO nl_dam_types VALUES (1, 'アーチダム');
INSERT INTO nl_dam_types VALUES (2, 'バットレスダム');
INSERT INTO nl_dam_types VALUES (3, 'アースダム');
INSERT INTO nl_dam_types VALUES (4, 'アスファルトフェイシングダム');
INSERT INTO nl_dam_types VALUES (5, 'アスファルトコアダム');
INSERT INTO nl_dam_types VALUES (6, 'フローティングゲートダム（可動堰）');
INSERT INTO nl_dam_types VALUES (7, '重力式コンクリートダム');
INSERT INTO nl_dam_types VALUES (8, '重力式アーチダム');
INSERT INTO nl_dam_types VALUES (9, '重力式コンクリートダム・フィルダム複合ダム');
INSERT INTO nl_dam_types VALUES (10, '中空重力式コンクリートダム');
INSERT INTO nl_dam_types VALUES (11, 'マルティプルアーチダム');
INSERT INTO nl_dam_types VALUES (12, 'ロックフィルダム');
INSERT INTO nl_dam_types VALUES (13, '台形 CSG ダム');

INSERT INTO nl_dam_purposes VALUES (1, '洪水調節，農地防災');
INSERT INTO nl_dam_purposes VALUES (2, '不特定用水，河川維持用水');
INSERT INTO nl_dam_purposes VALUES (3, '灌漑，特定(新規)灌漑用水');
INSERT INTO nl_dam_purposes VALUES (4, '上水道用水');
INSERT INTO nl_dam_purposes VALUES (5, '工業用水道用水');
INSERT INTO nl_dam_purposes VALUES (6, '発電');
INSERT INTO nl_dam_purposes VALUES (7, '消流雪用水');
INSERT INTO nl_dam_purposes VALUES (8, 'レクリエーション');

INSERT INTO nl_dam_precisions VALUES (1, '最高');
INSERT INTO nl_dam_precisions VALUES (2, '高');
INSERT INTO nl_dam_precisions VALUES (3, '中');
INSERT INTO nl_dam_precisions VALUES (4, '低');
INSERT INTO nl_dam_precisions VALUES (5, '低かつ河川名称不明');

INSERT INTO nl_dam_institutions VALUES (1, '国土交通省（各地方整備局，北海道開発局含む）');
INSERT INTO nl_dam_institutions VALUES (2, '沖縄振興局（旧沖縄開発庁）');
INSERT INTO nl_dam_institutions VALUES (3, '農林水産省（各地方農政局含む）');
INSERT INTO nl_dam_institutions VALUES (4, '都道府県');
INSERT INTO nl_dam_institutions VALUES (5, '市区町村');
INSERT INTO nl_dam_institutions VALUES (6, '水資源機構（旧水資源開発公団）');
INSERT INTO nl_dam_institutions VALUES (7, 'その他の公共企業体');
INSERT INTO nl_dam_institutions VALUES (8, '土地改良区');
INSERT INTO nl_dam_institutions VALUES (9, '利水組合・用水組合');
INSERT INTO nl_dam_institutions VALUES (10, '電力会社・電源開発株式会社');
INSERT INTO nl_dam_institutions VALUES (11, 'その他の企業');
INSERT INTO nl_dam_institutions VALUES (12, '個人');
INSERT INTO nl_dam_institutions VALUES (13, 'その他');

INSERT INTO nl_dams VALUES (
    1102,
    36.56647116,
    137.66208791,
    '黒部',
    '黒部川',
    '黒部川',
    1,
    6,
    186.0,
    492.0,
    1582,
    199285,
    10,
    1963,
    '富山県中新川郡立山町大字芦峅寺',
    1
)
