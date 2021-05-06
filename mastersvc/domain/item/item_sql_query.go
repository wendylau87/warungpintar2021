package item

const(
	CreateItem = `INSERT INTO item(name) VALUES(?)`;
	ReadItem = `SELECT id, name, total FROM item`

)
