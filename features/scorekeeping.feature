Feature: Keep players scores
    In order to keep track of player scores
    As a user of the scorekeeper microservice
    I need to be able to add points and subtract points from players scores

    Background:
        Given that the scorekeeper service is running

    Scenario: Scorekeeper service
        Then I can add some points

    Scenario: Scorekeeper service
        Then I can subs some points

    Scenario: Scorekeeper service
        Then I cant multiply points