SET foreign_key_checks = 0;
TRUNCATE TABLE `articles`;
TRUNCATE TABLE `tags`;
TRUNCATE TABLE `tag_relation`;

SET NAMES utf8mb4;

INSERT INTO `articles` (`title`, `body`, `date`) VALUES
('Article 1',	'This is a very long body.',	'2021-02-02 06:52:44'),
('Article 2',	'This is a very long body.',	'2021-02-02 06:53:00'),
('Article 3',	'Today is different day',	'2021-02-02 21:28:26'),
('Article 4',	'Today is another different day',	'2021-02-02 21:28:51');


INSERT INTO `tags` (`name`) VALUES
('science'),
('english'),
('biology'),
('opt'),
('computing'),
('psychology');


INSERT INTO `tag_relation` (`id`, `article_id`, `tag_id`, `date`) VALUES
(1,	1,	1,	'2021-02-02 06:52:44'),
(2,	1,	2,	'2021-02-02 06:52:44'),
(3,	2,	3,	'2021-02-02 06:53:00'),
(4,	2,	4,	'2021-02-02 06:53:00'),
(5,	3,	1,	'2021-02-02 21:28:26'),
(6,	3,	5,	'2021-02-02 21:28:26'),
(7,	4,	1,	'2021-02-02 21:28:51'),
(8,	4,	6,	'2021-02-02 21:28:51');
