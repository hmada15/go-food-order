CREATE TABLE `products` (
  `id` INT PRIMARY KEY AUTO_INCREMENT,
  `name` VARCHAR(255) NOT NULL,
  `slug` VARCHAR(255) NOT NULL,
  `description` VARCHAR(255),
  `regular_price` INT(11) NOT NULL,
  `sale_price` INT(11) NOT NULL,
  `in_stock` BOOLEAN NOT NULL,
  `is_publish` BOOLEAN NOT NULL,
  `category_id` INT NOT NULL,
  `created_at` TIMESTAMP DEFAULT now(),
  `updated_at` TIMESTAMP DEFAULT now(),
  UNIQUE (`slug`)
);

ALTER TABLE `products` ADD FOREIGN KEY (`category_id`) REFERENCES `product_categories` (`id`);
