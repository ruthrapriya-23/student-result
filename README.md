Student Result Processing CLI

Project Description

This project is a modular Command Line Interface (CLI) application developed using Go. It accepts student details and subject marks, calculates the total and average, assigns a grade, and determines the PASS or FAIL status.

Objectives

Learn basic Go programming.

Use variables, data types, slices, loops, and conditional statements.

Create and use functions.

Organize code using multiple packages.

Use Go Modules.

Practice Git and GitHub version control.


Technologies Used

Go 1.27.1

Git

GitHub

Visual Studio Code


Project Structure

student-result-cli/
│
├── go.mod
├── main.go
├── README.md
│
├── model/
│   └── student.go
│
├── result/
│   └── calculator.go
│
└── display/
└── display.go

Modules

Model Package

The "model" package contains the "Student" structure.

It stores:

Student name

Register number

Subject marks


Result Package

The "result" package performs:

Total calculation

Average calculation

Grade calculation

PASS/FAIL checking


Display Package

The "display" package formats and displays the student's result.

Grade System

Average| Grade
90 and above| A+
80–89| A
70–79| B
60–69| C
50–59| D
Below 50| F

A student is marked PASS only when every subject mark is at least 40.

How to Run

Clone or download the project and open the project folder in a terminal.

Initialize dependencies:

go mod tidy

Run the application:

go run .

Build the application:

go build

Sample Input

Enter Student Name: Nadhiya
Enter Register Number: 71
Enter Number of Subjects: 5
Enter marks for Subject 1: 80
Enter marks for Subject 2: 90
Enter marks for Subject 3: 95
Enter marks for Subject 4: 85
Enter marks for Subject 5: 98

Sample Output

====================================
STUDENT RESULT

Name        : Nadhiya
Register No : 71

Subject 1    : 80
Subject 2    : 90
Subject 3    : 95
Subject 4    : 85
Subject 5    : 98

Total       : 448
Average     : 89.60
Grade       : A
Status      : PASS

Git Workflow

The project uses Git for version control.

Main branch:

main

Development branch:

development

The development branch was used for making and committing changes before merging them into the main branch.

Example commands:

git init
git add .
git commit -m "Initial student result application"
git branch development
git switch development
git commit -m "Document grade calculation logic"
git switch main
git merge development
git push -u origin main
git push -u origin development

Conclusion

The project demonstrates a modular Go CLI application for processing student results. It applies Go programming concepts such as structures, slices, functions, loops, conditional statements, packages, and Go Modules. Git and GitHub are used to maintain the project using branches, commits, and merging.