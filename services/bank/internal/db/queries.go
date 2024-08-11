package db

const createBankQuery = `INSERT INTO bank (id, value) 
	VALUES (%s, %d)
	RETURNING id, value;`

const findBankQuery = `SELECT * FROM bank 
	WHERE id='%s';`

const updateBankQuery = `UPDATE bank 
	SET value='%d' 
	WHERE id='%s'
	RETURNING id, value;`