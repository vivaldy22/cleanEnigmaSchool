package queries

const (
	SELECT_ALL_TEACHER = `SELECT * FROM teacher`
	SELECT_TEACHER_ID  = `SELECT * FROM teacher WHERE id = ?`
	INSERT_TEACHER     = `INSERT INTO teacher VALUES (?, ?, ?, ?)`
	DELETE_TEACHER_ID  = `DELETE FROM teacher WHERE id = ?`
	UPDATE_TEACHER     = `UPDATE teacher
						  SET first_name = ?, last_name = ?, email = ?
						  WHERE id = ?`

	SELECT_ALL_STUDENT = `SELECT * FROM student`
	SELECT_STUDENT_ID  = `SELECT * FROM student WHERE id = ?`
	INSERT_STUDENT     = `INSERT INTO student VALUES (?, ?, ?, ?)`
	DELETE_STUDENT_ID  = `DELETE FROM student WHERE id = ?`
	UPDATE_STUDENT     = `UPDATE student
						  SET first_name = ?, last_name = ?, email = ?
						  WHERE id = ?`

	SELECT_ALL_SUBJECT = "SELECT * FROM subject"
	SELECT_SUBJECT_ID  = `SELECT * FROM subject WHERE id = ?`
	INSERT_SUBJECT     = `INSERT INTO subject VALUES (?, ?)`
	DELETE_SUBJECT_ID  = `DELETE FROM subject WHERE id = ?`
	UPDATE_SUBJECT     = `UPDATE subject
						  SET subject_name = ?
						  WHERE id = ?`
)
