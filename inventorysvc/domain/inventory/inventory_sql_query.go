package inventory

const(
	CreateInventory = `INSERT INTO inventory(inbound_detail_id, item_id, quantity, created_at) VALUES(?, ?, ?, ?)`
	SumInventoryQuantityByItem = `SELECT SUM(quantity) from inventory where item_id = ?`
	GetInventoryByInboundDetail =  `SELECT id, inbound_detail_id, item_id, quantity, created_at FROM inventory where inbound_detail_id = ?`
)
