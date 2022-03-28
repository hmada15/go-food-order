CREATE TABLE `product_categories` (
  `id` INT PRIMARY KEY AUTO_INCREMENT,
  `name` VARCHAR(255) NOT NULL,
  `slug` VARCHAR(255) NOT NULL,
  `description` VARCHAR(255),
  `is_publish` BOOLEAN NOT NULL,
  `created_at` TIMESTAMP DEFAULT now(),
  `updated_at` TIMESTAMP DEFAULT now(),
  UNIQUE (`slug`)
);

