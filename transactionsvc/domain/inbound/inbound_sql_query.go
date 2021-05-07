package inbound

const(
	CreateInbound = `INSERT INTO inbound(po_number, created_at) VALUES(?, ?)`;
	CreateInboundDetail = `INSERT INTO inbound_detail(inbound_id, item_id, quantity) VALUES(?,?,?)`;
	ReadInbounds = `SELECT id, po_number, created_at FROM inbound`
	ReadInboundDetails = `SELECT id, inbound_id, item_id, quantity FROM inbound_detail WHERE inbound_id = ?`
)
