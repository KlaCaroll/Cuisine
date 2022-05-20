CREATE TABLE IF NOT EXISTS meal (
    id         [UNSIGNED INTEGER] NOT NULL,
    name       TINYTEXT           NOT NULL,
    planned_at DATE               NOT NULL,
    type       TEXT               NOT NULL,
    sub_type   TEXT,
    pers       INTEGER,
    PRIMARY KEY (
        id
    )
);

CREATE TABLE IF NOT EXISTS meal_recipe (
    meals_id        [UNSIGNED INTEGER],
    recipe_id      [UNSIGNED INTEGER],
    quantity        [UNSIGNED INTEGER]
);

CREATE TABLE IF NOT EXISTS recipe (
    id       [UNSIGNED INTEGER]  NOT NULL,
    name     TEXT                NOT NULL,
--    min_pers [UNSIGNED INTERGER] NOT NULL,
--    max_pers [UNSIGNED INTEGER]  NOT NULL
);

CREATE TABLE IF NOT EXISTS recipe_food (
    dish_id       [UNSIGNED INTEGER],
    food_id       [UNSIGNED INTEGER],
    quantity      [UNSIGNED INTEGER]
);

CREATE TABLE IF NOT EXISTS food (
    id           [UNSIGNED INTEGER] NOT NULL,
    name         TEXT               NOT NULL,
--    type         TEXT               NOT NULL,
--    sub_type     TEXT,
--    brand        TEXT,
--    min_quantity [UNSIGNED INTEGER],
--    min_weight   FLOAT,
--    Liter        INTEGER,
    PRIMARY KEY (
        id
    )
);

