use management_inventaris ;

/*
b. Write SQL scripts to:
     - Memasukkan data sampel ke dalam tabel-tabel.
     - Melakukan query untuk produk, inventaris, dan detail pesanan.
     - Melakukan agregasi seperti total pesanan untuk suatu produk atau tingkat stok di lokasi tertentu.
*/

INSERT INTO produks (id, nama, deskripsi, harga, kategori, created_at, updated_at, deleted_at) VALUES
(1, 'Dr. Martens 1460', 'Iconic 8-eye leather boots with AirWair sole', 2500000, 'Footwear', NOW(), NOW(), NULL),
(2, 'Dr. Martens 1461', 'Classic 3-eye shoe with smooth leather finish', 2200000, 'Footwear', NOW(), NOW(), NULL),
(3, 'Dr. Martens Jadon', 'Platform boots with bold, chunky sole', 3000000, 'Footwear', NOW(), NOW(), NULL),
(4, 'Dr. Martens 2976 Chelsea', 'Slip-on boots with elastic gussets and smooth leather', 2700000, 'Footwear', NOW(), NOW(), NULL),
(5, 'Dr. Martens Sinclair', 'Milled Nappa leather boots with removable jungle zip', 3200000, 'Footwear', NOW(), NOW(), NULL),
(6, 'Dr. Martens Pascal', '8-eye boots with soft Virginia leather for added comfort', 2600000, 'Footwear', NOW(), NOW(), NULL);

INSERT INTO inventaris (id_produk, jumlah, lokasi, created_at, updated_at, deleted_at) VALUES
(1, 150, 'Jakarta WH - Indonesia', NOW(), NOW(), NULL),
(2, 200, 'Shanghai WH - China', NOW(), NOW(), NULL),
(3, 120, 'Bangkok WH - Thailang', NOW(), NOW(), NULL),
(4, 180, 'London WH - UK', NOW(), NOW(), NULL),
(5, 160, 'Los Angeles WH - US', NOW(), NOW(), NULL),
(6, 140, 'Berlin WH - Germany', NOW(), NOW(), NULL);

INSERT INTO pesanans (id_pesanan, id_produk, jumlah, tanggal, created_at, updated_at, deleted_at) VALUES
(1001, 1, 3, '2025-02-08', NOW(), NOW(), NULL),
(1002, 2, 1, '2025-02-07', NOW(), NOW(), NULL),
(1003, 3, 2, '2025-02-06', NOW(), NOW(), NULL),
(1004, 4, 4, '2025-02-05', NOW(), NOW(), NULL),
(1005, 5, 2, '2025-02-04', NOW(), NOW(), NULL),
(1006, 6, 5, '2025-02-03', NOW(), NOW(), NULL);
