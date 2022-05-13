-- Encodage texte utilisé : UTF-8

-- Table : meals
CREATE TABLE meals (
    id         [UNSIGNED INTEGER] NOT NULL,
    planned_at DATE               NOT NULL,
    type       TEXT               NOT NULL,
    sub_type   TEXT               NOT NULL,
    pers       INTEGER,
    name       TINYTEXT           NOT NULL,
    PRIMARY KEY (
        id
    )
);

-- Table : meals_recipes
CREATE TABLE meals_recipes (
    meals_id [UNSIGNED INTEGER],
    recipes_id [UNSIGNED INTEGER]
);

-- Table : recipes
CREATE TABLE recipes (
    id       [UNSIGNED INTEGER]  NOT NULL,
    name     TEXT                NOT NULL,
    min_pers [UNSIGNED INTERGER] NOT NULL,
    max_pers [UNSIGNED INTEGER]  NOT NULL
);


-- Table : recipes_food
CREATE TABLE recipes_food (
    dish_id       [UNSIGNED INTEGER],
    food_id [UNSIGNED INTEGER],
    quantity      [UNSIGNED INTEGER]
);

-- Table : food
CREATE TABLE food (
    id           [UNSIGNED INTEGER] NOT NULL,
    type         TEXT               NOT NULL,
    sub_type     TEXT,
    name         TEXT               NOT NULL,
    brand        TEXT,
--    min_quantity [UNSIGNED INTEGER],
--    min_weight   FLOAT,
--    Liter        INTEGER,
    PRIMARY KEY (
        id
    )
);