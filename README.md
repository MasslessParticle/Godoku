# goduku
A small application to solve sudoku puzzles via a RESTful API 

Godoku is a learning exercise so it does woefully little validation or error handing.

You'll need a mysql instance installed and running with the user sudoku and the database sudokusolver_development 

The application has the following end points:
POST to /puzzle returns the id of the posted puzzle
GET /puzzle/{id} gets the persisted puzzle
GET /solved/{id} gets the solved persisted puzzle

