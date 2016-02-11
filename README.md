# Sugoku
A small application to solve sudoku puzzles via a RESTful API 

## What It Does
The application has the following end points:
- POST to /puzzle returns the id of the posted puzzle
- GET /puzzle/{id} gets the persisted puzzle
- GET /solved/{id} gets the solved persisted puzzle

## Puzzle format
- puzzles are formatted as a single string where spaces are represented as _
- e.g.: 97_1_2685____85___8____6_3____43__5_71_6_8_93_2__97____3_8____6___75____5492_3_78

## Setup
- Sugoku is based on the Cloud Foundry sample application located at: https://github.com/cloudfoundry-samples/pong_matcher_go. Setup for Sugoku is nearly identical.
- The database name is sudokusolver_development and the DB user/password are both sudoku

### This isn't even my final form!
Sugoku is a learning exercise so it does woefully little validation or error handing. 







