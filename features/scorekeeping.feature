Feature: Keep players scores
    In order to keep track of player scores
    As a user of the scorekeeper microservice
    I need to be able to add points and subtract points from players scores

    Scenario: Scorekeeper service
        Given that the scorekeeper service is running
        Then I can add some points

    Scenario: Scorekeeper service
        Given that the scorekeeper service is running
        Then I can subs some points

    Scenario: Scorekeeper service
        Given that the scorekeeper service is running
        Then I cant multiply points