CREATE TABLE IF NOT EXISTS `users` (
  `id` INT PRIMARY KEY,
  `name` VARCHAR(255) NOT NULL
);

INSERT INTO `users` (`id`, `name`) VALUES (1, 'John'), (2, 'Hanako');