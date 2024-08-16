package db

const createUserQuery = `INSERT INTO users (id, username, password) 
	VALUES (gen_random_uuid (), '%s', '%s')
	RETURNING id, username, password;`

const findUserQuery = `SELECT * FROM users 
	WHERE username='%s';`

const updateUserQuery = `UPDATE users 
	SET password='%s' 
	WHERE username='%s'
	RETURNING id, username, password;`