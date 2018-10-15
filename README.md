# gofood

Basic command line app written in Go to print out the current menu for the dining halls at Princeton University.

## Installation

```
go get github.com/gronnesby/gofood

go install gofood
```


## Usage

Running `gofood` without arguments will print out todays menus for all dining halls and all meals.
Command line arguments can be added to limit the search.

Get a specific dining hall can be done by adding the command line argument `-loc=` followed by a search query.

```
$ gofood -loc=grad

-- Graduate College --
Brunch
-- Entrees --	
Bacon
Omelet Bar with Pork Options
Pork Sausage Links
Potato Puffs
Scrambled Eggs

-- Breakfast Bars --	
Blueberry Pancakes
Hot Cranberry Orange Sunrise Porridge
Pancake Toppings
Waffle Bar

-- Main Entree --	
Chicken Lemon

-- Vegetarian & Vegan Entree --	
Manicotti

-- On the Side --	
Honey Mashed Sweet Potatoes
Rustic Green Beans

```

To get the menu for the next day, add the `-t` argument.

```
$ gofood -t

Tomorrows Menu

-- Graduate College --
Dinner
-- Soups --	
Ecuadorian Vegetable Soup
Lentil & Ham Soup

-- Main Entree --	
Garlic Roasted Chicken Thigh
Oxtail Beef Stew

-- Vegetarian & Vegan Entree --	
Vegan Spinach & Quinoa Enchilada

-- On the Side --	
Creamed Spinach
Grilled Vegetables
Yellow Rice

-- From our Bakeshop --	
Caramel Apple Bar
Maple Glazed Sweet Potato Cake
Tiger Cupcake

```

Adding the `-tab` argument will print out each menu in a table.
Note that this may misalign for large menus.

```
-- Graduate College --
Brunch
-- Entrees --                   |-- Breakfast Bars --                     |-- Main Entree --    |-- Vegetarian & Vegan Entree --    |-- On the Side --              |
Bacon                           |Blueberry Pancakes                       |Chicken Lemon        |Manicotti                          |Honey Mashed Sweet Potatoes    |
Omelet Bar with Pork Options    |Hot Cranberry Orange Sunrise Porridge    |                     |                                   |Rustic Green Beans             |
Pork Sausage Links              |Pancake Toppings                         |                     |                                   |                               |
Potato Puffs                    |Waffle Bar                               |                     |                                   |                               |
Scrambled Eggs                  |                                         |                     |                                   |                               |
```