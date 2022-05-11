--
-- Fichier généré par SQLiteStudio v3.3.3 sur mer. mai 11 15:20:19 2022
--
-- Encodage texte utilisé : UTF-8
--
PRAGMA foreign_keys = off;
BEGIN TRANSACTION;

-- Table : dishes
CREATE TABLE dishes (
    id   [UNSIGNED INTEGER] NOT NULL,
    name TEXT               NOT NULL
);


-- Table : dishes_ingredients
CREATE TABLE dishes_ingredients (
    dish_id       [UNSIGNED INTEGER],
    ingredient_id [UNSIGNED INTEGER],
    quantity      [UNSIGNED INTEGER]
);


-- Table : food
CREATE TABLE food (
    id           [UNSIGNED INTEGER] NOT NULL,
    type         TEXT               NOT NULL,
    sub_type     TEXT,
    name         TEXT               NOT NULL,
    brand        TEXT,
    min_quantity [UNSIGNED INTEGER],
    min_weight   FLOAT,
    Liter        INTEGER,
    PRIMARY KEY (
        id
    )
);



-- Table : ingredients
CREATE TABLE ingredients (
    id   [UNSIGNED INTEGER] NOT NULL,
    name TEXT               NOT NULL
);


-- Table : meal
CREATE TABLE meal (
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

-- Table : recipes
CREATE TABLE recipes (
    id       [UNSIGNED INTEGER]  NOT NULL,
    name     TEXT                NOT NULL,
    min_pers [UNSIGNED INTERGER] NOT NULL,
    max_pers [UNSIGNED INTEGER]  NOT NULL
);


COMMIT TRANSACTION;
PRAGMA foreign_keys = on;
