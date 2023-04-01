# GoSqlRaw

Deleting the DB tables in every fresh run to avoid the duplication input of data in the database


### Task-3: calculate the total cost of each department. The cost will be calculated by teachers&#39;
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
