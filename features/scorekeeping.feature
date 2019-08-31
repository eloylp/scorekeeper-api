Feature: Keep players scores
    In order to keep track of player scores
    As a user of the scorekeeper microservice
    I need to be able to add points and subtract points from players scores

    Background:
        Given that the scorekeeper service is running

    Scenario: Scorekeeper service points operations
        Given I add 5 points to user "Bob"
        And I add 5 points to user "Bob"
        And I add 5 points to user "Bob"
        Then "Bob" has now 15 points
        And I subs 5 points to user "Bob"
        Then "Bob" has now 10 points

