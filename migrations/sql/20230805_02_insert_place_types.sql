-- +migrate Up
INSERT INTO `place_types` (`id`, `key`, `name`) VALUES
('7e43f169-0f9f-43c4-944f-0ead52ecdf4c', 'amusement_park','遊園地'),
('641ad05d-0c9c-4a3a-a98b-6dbf765bc74a', 'aquarium','水族館'),
('bad0bce3-e49f-4e09-a79a-4704073b52da', 'art_gallery','美術館'),
('1d0935bf-0dae-451f-ae27-ff1cc1721696', 'bakery','パン屋'),
('f304d338-ec93-439d-8ce8-ca02b8a3ebff', 'bar','バー'),
('a31de9fc-b937-4284-92b8-158c71834620', 'book_store','本屋'),
('e61c04d4-24c6-4b66-8bff-1d13fd414728', 'cafe','カフェ'),
('f02262e1-29f6-4677-afb0-cfd56aa06501', 'church','教会'),
('9bfc1bf6-87bc-4f31-b6b1-5cf66937bc18', 'convenience_store','コンビニ'),
('ed762ebd-763d-4227-81a0-a231ecea94af', 'library','図書館'),
('a7e16b0b-b8a6-4a53-80dc-26410ba3817d', 'movie_theater','映画館'),
('5ef8018d-40b6-4691-8e99-42af2d0d4a58', 'museum','博物館'),
('17e277d8-10ec-4c63-b678-735f17e33afa', 'park','公園'),
('74c24504-1bcf-4d90-8d33-1d25cfa9295b', 'pet_store','ペットショップ'),
('c5fe1cb6-abec-4166-aded-f46d2c5ef859', 'restaurant','レストラン'),
('3655f820-9479-444b-bc69-0a10a8d6af10', 'shopping_mall','ショッピングモール'),
('c64453af-2432-4317-8fe9-76c7cb01ad1b', 'zoo','動物園');


-- +migrate Down