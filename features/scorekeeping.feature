Feature: Keep players scores
    In order to keep track of player scores
    As a user of the scorekeeper microservice
    I need to be able to add points and substract points from players scores

    Background:
      Given that the scorekeeper service is running

    Scenario: Scorekeeper single user operations
      Given I add 5 points to user "Bob"
      And I add 5 points to user "Bob"
      And I add 5 points to user "Bob"
      Then "Bob" has now 15 points
      And I subs 5 points to user "Bob"
      Then "Bob" has now 10 points

    Scenario: Scorekeeper multiple users operations
      Given I add 5 points to user "Bob"
      And I add 10 points to user "Alice"
      And I subs 2 points to user "Bob"
      Then "Alice" has now 10 points
      And "Bob" has now 3 points

    Scenario: Scorekeeper prevents negative points operations
      Given I add 10 points to user "Alice"
      And I subs 15 points to user "Alice"
      Then "Alice" has now 0 points