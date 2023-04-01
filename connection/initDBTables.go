package connection

func (d *DB) createDepartmentTable() error {
	err := d.Exec(`CREATE TABLE IF NOT EXISTS DEPARTMENTS (
		ID SERIAL PRIMARY KEY, 
		DEPARTMENT_NAME VARCHAR(50) NOT NULL,
		DEPARTMENT_CODE VARCHAR(50) NOT NULL)`).Error
	return err
}

func (d *DB) createTeachersTable() error {
	err := d.Exec(`CREATE TABLE IF NOT EXISTS TEACHERS (
					ID SERIAL PRIMARY KEY, 
		 			TEACHER_NAME VARCHAR(50) NOT NULL,
					SALARY INT NOT NULL,
					DEPARTMENT_ID INT NOT NULL REFERENCES DEPARTMENTS(ID))`).Error
	return err
}

func (d *DB) createStudentsTable() error {
	err := d.Exec(`CREATE TABLE IF NOT EXISTS STUDENTS (
		ID SERIAL PRIMARY KEY, 
		STUDENT_NAME VARCHAR(50) NOT NULL,
		TEACHER_ID INT NOT NULL REFERENCES TEACHERS(ID))`).Error
	return err
}

func (d *DB) initialInsertIntoDepartmentsTable() error {
	err := d.Exec(`INSERT INTO DEPARTMENTS (DEPARTMENT_NAME, DEPARTMENT_CODE) VALUES 
					('English', 'ELA_01'),
					('Mathematics', 'ELA_02'),
					('Bangla', 'ELA_03')`).Error
	return err
}

func (d *DB) initialInsertIntoTeachersTable() error {
	err := d.Exec(`INSERT INTO TEACHERS (TEACHER_NAME, SALARY, DEPARTMENT_ID) VALUES
					('William Shakespeare', 20000, 1),
					('Christopher Marlowe', 15000, 1),
					('John Milton', 12000, 2),
					('John Dryden', 10000, 2),
					('William Wordsworth', 17000, 3),
					('S.T. Coleridge', 11000, 2) `).Error
	return err
}

func (d *DB) initialInsertIntoStudentsTable() error {
	err := d.Exec(`INSERT INTO STUDENTS (STUDENT_NAME, TEACHER_ID) VALUES
					('Stuart Mil', 1),
					('Lord Alfred', 1),
					('Thomas Hardy', 3),
					('Emily Bronte', 3),
					('Leo Tolstoy', 1),
					('Karl Marx', 5)`).Error
	return err
}
