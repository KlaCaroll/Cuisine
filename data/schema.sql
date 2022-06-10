CREATE TABLE IF NOT EXISTS meal (
    id         INTEGER            NOT NULL,
    name       TINYTEXT           NOT NULL,
    planned_at DATE               NOT NULL,
    form       TEXT,
    sub_type   TEXT,
    pers       INTEGER,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS meal_recipe (
    meals_id        INTEGER,
    recipe_id      INTEGER,
    quantity_m       FLOAT
);

CREATE TABLE IF NOT EXISTS recipe (
    id       INTEGER  NOT NULL,
    recipe_name     TEXT                NOT NULL
--    min_pers [UNSIGNED INTERGER] NOT NULL,
--    max_pers [UNSIGNED INTEGER]  NOT NULL
);

CREATE TABLE IF NOT EXISTS recipe_food (
    recipe_id       INTEGER,
    food_id       INTEGER,
    quantity_r      FLOAT
);

CREATE TABLE IF NOT EXISTS food (
    id           INTEGER NOT NULL,
    food_name         TEXT               NOT NULL,
--    type         TEXT               NOT NULL,
--    sub_type     TEXT,
--    brand        TEXT,
--    min_quantity [UNSIGNED INTEGER],
--    min_weight   FLOAT,
--    Liter        INTEGER,
    PRIMARY KEY (id)
);
