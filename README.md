# Go Importer
First attempt at a basic go project.

### Prerequisite
Run `docker-compose build && docker-compose up -d` to create a DB instance for this project. After run the below query to create a `products` table:
```sql
DROP TABLE IF EXISTS `products`;
CREATE TABLE `products` (
    `id` int(6) unsigned NOT NULL AUTO_INCREMENT,
    `sku` int(6) unsigned NOT NULL,
    `plu` char(3) NOT NULL,
    `name` varchar(255) NOT NULL,
    `size` varchar(255) NOT NULL,
    `size_sort` varchar(255) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
```