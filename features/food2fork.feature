Feature: Food2Fork Search 
    In order to get a recipe
    As a user
    I would like to perform a search recipe

@search
Scenario: verify recipe is returned
Given the food to fork service is running
When I request a recipe about "cookies monster cupcakes"
Then the recipe is returned

@search
Scenario: verify recipe is returned
Given the food to fork service is running
When I request a recipe about "shkajdfhskjdhfsjhdfkjshdfkls"
Then the recipe is returned