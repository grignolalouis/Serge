-- Données de test pour le MCP SERGE

-- Insérer un ordre
INSERT INTO orders (order_number, customer_name, status, total_amount)
VALUES ('ORD-001', 'Acme Corp', 'pending', 1500.00);

-- Insérer une ligne de commande
INSERT INTO order_lines (order_id, sku, quantity, unit_price)
VALUES (
  (SELECT id FROM orders WHERE order_number = 'ORD-001'),
  'WIDGET-100',
  10,
  150.00
);

-- Insérer un inventaire
INSERT INTO inventories (sku, warehouse_location, quantity_on_hand)
VALUES ('WIDGET-100', 'Warehouse A', 500);

-- Insérer un fournisseur
INSERT INTO suppliers (supplier_name, category, contact_email)
VALUES ('Tech Supplies Inc', 'electronics', 'sales@techsupplies.com');

-- Insérer une commande d'achat
INSERT INTO purchase_orders (sku, supplier_id, quantity, order_date)
VALUES (
  'WIDGET-100',
  (SELECT id FROM suppliers WHERE supplier_name = 'Tech Supplies Inc'),
  100,
  NOW()
);

-- Insérer un envoi
INSERT INTO shipments (order_id, tracking_number, status, estimated_delivery)
VALUES (
  (SELECT id FROM orders WHERE order_number = 'ORD-001'),
  'TRK-2024-123456',
  'in_transit',
  NOW() + INTERVAL '3 days'
);
