# triviageek-hand-s-on

## 1 - Install angular project
Angular project is located under *gui* folder.

Install grunt dependencies with

    npm install
    
And javascript dependencies with

    bower install

You can then run grunt server with

    grunt serve

## 2 - Run Golang project

To run project use

    make run
    
To execute test files use

    make test

## The game

During this hand's on you will be asked to implement the following scenario

When Bob joins the game
 
    {"pseudo":"Bob"}

The server answer its name and a startTime so that client can compute remaining seconds before starting

    {"name":"QaUshs","startTime":"2016-03-07T16:44:24.592546399-05:00","step":0}

The server then pushes every 20 seconds a new question to all players. 

    {"step":1, "smell": {"name":"","description":"..."]}

A game is made of 12 smells

    {"step":12, "smell": {"name":"","description":"..."]}

When client wants to submit an answer, he pushes the following json structure to server

    {"step":5,"value":"Combinatorial Explosion"}

At the end of the game, each player receives all results

    {"players":[{"pseudo":"Ben","score":0}]}}]}

