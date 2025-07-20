-- Update invoice total
UPDATE invoices i
JOIN (
    SELECT s.invoice_id, SUM(p.price * s.quantity) AS total
    FROM sales s
    JOIN products p ON p.id = s.product_id
    GROUP BY s.invoice_id
) t ON i.id = t.invoice_id
SET i.total = t.total;