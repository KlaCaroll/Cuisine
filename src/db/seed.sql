insert into food (id, name) values
(1, 'oeuf'),
(2, 'lait'),
(3, 'beurre'),
(4, 'chocolat');

insert into recipe (id, name) values
(1, 'gateau'),
(2, 'oeuf au plat');

insert into recipe_food (recipe_id, food_id, quantity) values
(1, 1, 1),
(1, 2, 200),
(1, 3, 200),
(1, 4, 200);

insert into meal (id, planned_at, guests) values
(1, '2022-07-01 21:00:00', 4);

insert into meal_recipe (meal_id, recipe_id) values
(1, 1);