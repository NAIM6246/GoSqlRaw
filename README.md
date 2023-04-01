# GoSqlRaw

### Task-1: create three tables by SQL script.
Inside connection package's initDBTables.go file there are three methods createDepartmentTable(), createTeachersTable() and createStudentsTable() which creates required tables using raw SQL script.

### Task-2: save those data in the corresponding table by go-program.
In the initDBTables.go of connection package there are three methods initialInsertIntoDepartmentsTable(), initialInsertIntoTeachersTable(), initialInsertIntoStudentsTable() which inserts the initial data(BULK data) of the tables.

And in the db.go file there is three methods AddStudents(), AddTeacher(), AddDepartment() which insert single data to the specific table.

### Task-3: calculate the total cost of each department. The cost will be calculated by teachers's salaries.

salaries.
Endpoint: localhost:8005/api/calculate-department-costs
Method: GET
Response:

```
[
  {
    "id": 1,
    "departmentName": "English",
    "totalCost": 35000
  },
  {
    "id": 2,
    "departmentName": "Mathematics",
    "totalCost": 33000
  },
  {
    "id": 3,
    "departmentName": "Bangla",
    "totalCost": 17000
  }
]
```

### Task-4: find two teachers whose salary is higher than the other teachers.

Endpoint: localhost:8005/api/highest-salary-teachers
Method: GET
Response:

```
[
  {
    "id": 1,
    "teacherName": "William Shakespeare",
    "salary": 20000,
    "departmentName": "English"
  },
  {
    "id": 5,
    "teacherName": "William Wordsworth",
    "salary": 17000,
    "departmentName": "Bangla"
  }
]
```

### Task-5: find the total student of each department.

Endpoints: localhost:8005/api/departments-total-student
Method: GET
Response:

```
[
  {
    "id": 1,
    "departmentName": "",
    "totalStudent": 3
  },
  {
    "id": 2,
    "departmentName": "",
    "totalStudent": 2
  },
  {
    "id": 3,
    "departmentName": "",
    "totalStudent": 1
  }
]
```

### Task-7: write a program in which your expected output result will be written in a csv file.

Inside utils package's helper.go file there is a InitFileWriter() which creates a csv file and new csv writer.
GetCSVFileWriter() returns the CsvWriter's instance which was initialized before.
Write() method of helper.go writes in the csv file previously initialized

### To add student:

Endpoint: localhost:8005/api/add-student
Method: POST
Request Body:

```
{
    "studentName" : "st1",
    "teacherID" : 1
}
```

Response: Status Code: 201

### To add teacher

Endpoint: localhost:8005/api/add-teacher
Method: POST
Request Body:

```
{
  "teacherName" : "Mr. ABC",
  "salary" : 19000,
  "departmentID" : 2
}
```

Response: Status Code: 201

### To add department

Endpoint: localhost:8005/api/add-department
Method: POST
Request Body:

```
{
  "departmentName" : "Physics",
  "departmentCode" : "PHY_01"
}
```

Response: Status Code: 201

### Notes

Deleting the DB tables in every fresh run to avoid the duplication input of data in the database
