Initialize
-go mod init

compileDeamon
-go get github.com/githubnemo/CompileDaemon
-go install github.com/githubnemo/CompileDaemon

joho
-go get github.com/joho/godotenv

Gin web framework
-go get -u github.com/gin-gonic/gin

GORM library
-go get -u gorm.io/gorm
-go get -u gorm.io/driver/postgres

COMPILE
-$ CompileDaemon -command="./MyProgram -my-options"

Creating a postgres db

docker run -d   --name my-postgres-db   -e POSTGRES_USER=myuser   -e POSTGRES_PASSWORD=gocrud   -e POSTGRES_DB=gocrud   -p 5432:5432   postgres

password gocrudpass

CONNECTION_STRING="host=localhost user=myuser password=gocrudpass dbname=gocrud port=5432 sslmode=disable timezone=Europe/Dublin"
psql -h localhost -U myuser -d gocrud

- go run migrate/migrate.go

connectin to TablePlus
-postgresql://myuser@localhost/gocrud


-docker-compose logs


-docker run -it goapi /bin/sh

docker-compose down
docker-compose up -d

------

docker run --name my-postgres-db -e POSTGRES_USER=myuser -e POSTGRES_PASSWORD=gocrudpass -e POSTGRES_DB=gocrud -p 5432:5432 -v api_postgres_data:/var/lib/postgresql/data -d postgres

psql -U myuser -d gocrud 

list tables
-\dt

explore tables
-\d+ categories
-\d+ products


-curl http://localhost:13000

RUN app inside container
-docker exec -it crud-in-go-app-1 /bin/sh
-CompileDaemon -command=./crud-in-go


docker run --name go-postgres-db -e POSTGRES_USER=gopostgresuser -e POSTGRES_PASSWORD=gopostgrespass -e POSTGRES_DB=gopostgresdatabase -p 5432:5432 -d postgres

psql -U gopostgresuser -d gopostgresdatabase


-- Corrected INSERT statements for Men
INSERT INTO products (name, brand, category_id, price, stock_quantity, description, release_date, image_url) VALUES
('Mens Jacket', 'BrandA', 1, 99.99, 50, 'Warm and stylish jacket for men.', '2024-01-15', 'https://picsum.photos/300?random=1'),
('Mens Shirt', 'BrandB', 1, 29.99, 100, 'Casual shirt for men.', '2024-02-20', 'https://picsum.photos/300?random=2'),
('Mens Jeans', 'BrandC', 1, 49.99, 75, 'Comfortable jeans for men.', '2024-03-10', 'https://picsum.photos/300?random=3'),
('Mens Shoes', 'BrandD', 1, 79.99, 60, 'Durable shoes for men.', '2024-04-05', 'https://picsum.photos/300?random=4'),
('Mens Hat', 'BrandE', 1, 19.99, 120, 'Fashionable hat for men.', '2024-05-25', 'https://picsum.photos/300?random=5'),
('Mens Socks', 'BrandF', 1, 9.99, 200, 'Pack of 3 socks for men.', '2024-06-15', 'https://picsum.photos/300?random=6'),
('Mens Belt', 'BrandG', 1, 39.99, 80, 'Leather belt for men.', '2024-07-10', 'https://picsum.photos/300?random=7'),
('Mens Suit', 'BrandH', 1, 199.99, 40, 'Formal suit for men.', '2024-08-01', 'https://picsum.photos/300?random=8'),
('Mens Watch', 'BrandI', 1, 149.99, 30, 'Elegant watch for men.', '2024-09-10', 'https://picsum.photos/300?random=9'),
('Mens Sunglasses', 'BrandJ', 1, 89.99, 90, 'Stylish sunglasses for men.', '2024-10-20', 'https://picsum.photos/300?random=10');



-- Corrected INSERT statements for Women
INSERT INTO products (name, brand, category_id, price, stock_quantity, description, release_date, image_url) VALUES
('Womens Dress', 'BrandK', 2, 89.99, 55, 'Elegant dress for women.', '2024-01-15', 'https://picsum.photos/300?random=11'),
('Womens Skirt', 'BrandL', 2, 39.99, 65, 'Stylish skirt for women.', '2024-02-20', 'https://picsum.photos/300?random=12'),
('Womens Blouse', 'BrandM', 2, 29.99, 80, 'Chic blouse for women.', '2024-03-10', 'https://picsum.photos/300?random=13'),
('Womens Trousers', 'BrandN', 2, 49.99, 70, 'Comfortable trousers for women.', '2024-04-05', 'https://picsum.photos/300?random=14'),
('Womens Shoes', 'BrandO', 2, 79.99, 60, 'Fashionable shoes for women.', '2024-05-25', 'https://picsum.photos/300?random=15'),
('Womens Scarf', 'BrandP', 2, 19.99, 100, 'Warm scarf for women.', '2024-06-15', 'https://picsum.photos/300?random=16'),
('Womens Bag', 'BrandQ', 2, 89.99, 45, 'Stylish bag for women.', '2024-07-10', 'https://picsum.photos/300?random=17'),
('Womens Coat', 'BrandR', 2, 129.99, 35, 'Warm coat for women.', '2024-08-01', 'https://picsum.photos/300?random=18'),
('Womens Hat', 'BrandS', 2, 19.99, 120, 'Fashionable hat for women.', '2024-09-10', 'https://picsum.photos/300?random=19'),
('Womens Sunglasses', 'BrandT', 2, 89.99, 90, 'Stylish sunglasses for women.', '2024-10-20', 'https://picsum.photos/300?random=20');

-- Corrected INSERT statements for Kids
INSERT INTO products (name, brand, category_id, price, stock_quantity, description, release_date, image_url) VALUES
('Kids T-shirt', 'BrandU', 3, 14.99, 150, 'Comfortable T-shirt for kids.', '2024-01-15', 'https://picsum.photos/300?random=21'),
('Kids Pants', 'BrandV', 3, 19.99, 100, 'Stylish pants for kids.', '2024-02-20', 'https://picsum.photos/300?random=22'),
('Kids Shoes', 'BrandW', 3, 29.99, 75, 'Durable shoes for kids.', '2024-03-10', 'https://picsum.photos/300?random=23'),
('Kids Hat', 'BrandX', 3, 9.99, 200, 'Cool hat for kids.', '2024-04-05', 'https://picsum.photos/300?random=24'),
('Kids Socks', 'BrandY', 3, 4.99, 300, 'Pack of 5 socks for kids.', '2024-05-25', 'https://picsum.photos/300?random=25'),
('Kids Jacket', 'BrandZ', 3, 39.99, 50, 'Warm jacket for kids.', '2024-06-15', 'https://picsum.photos/300?random=26'),
('Kids Scarf', 'BrandAA', 3, 14.99, 100, 'Cute scarf for kids.', '2024-07-10', 'https://picsum.photos/300?random=27'),
('Kids Gloves', 'BrandBB', 3, 9.99, 120, 'Winter gloves for kids.', '2024-08-01', 'https://picsum.photos/300?random=28'),
('Kids Watch', 'BrandCC', 3, 24.99, 60, 'Stylish watch for kids.', '2024-09-10', 'https://picsum.photos/300?random=29'),
('Kids Backpack', 'BrandDD', 3, 29.99, 70, 'Cool backpack for kids.', '2024-10-20', 'https://picsum.photos/300?random=30');

-- Corrected INSERT statements for Accessories
INSERT INTO products (name, brand, category_id, price, stock_quantity, description, release_date, image_url) VALUES
('Watch', 'BrandEE', 4, 199.99, 50, 'Luxury watch.', '2024-01-15', 'https://picsum.photos/300?random=31'),
('Bracelet', 'BrandFF', 4, 29.99, 100, 'Stylish bracelet.', '2024-02-20', 'https://picsum.photos/300?random=32'),
('Necklace', 'BrandGG', 4, 49.99, 75, 'Elegant necklace.', '2024-03-10', 'https://picsum.photos/300?random=33'),
('Ring', 'BrandHH', 4, 99.99, 60, 'Beautiful ring.', '2024-04-05', 'https://picsum.photos/300?random=34'),
('Earrings', 'BrandII', 4, 19.99, 120, 'Chic earrings.', '2024-05-25', 'https://picsum.photos/300?random=35'),
('Belt', 'BrandJJ', 4, 39.99, 80, 'Leather belt.', '2024-06-15', 'https://picsum.photos/300?random=36'),
('Scarf', 'BrandKK', 4, 14.99, 100, 'Warm scarf.', '2024-07-10', 'https://picsum.photos/300?random=37'),
('Hat', 'BrandLL', 4, 24.99, 120, 'Stylish hat.', '2024-08-01', 'https://picsum.photos/300?random=38'),
('Sunglasses', 'BrandMM', 4, 79.99, 40, 'Cool sunglasses.', '2024-09-10', 'https://picsum.photos/300?random=39'),
('Wallet', 'BrandNN', 4, 29.99, 90, 'Leather wallet.', '2024-10-20', 'https://picsum.photos/300?random=40');
