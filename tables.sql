CREATE TABLE `links` (
  `id`  INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `url` VARCHAR(191)     NOT NULL,
  PRIMARY KEY (`id`)
)
  ENGINE = MYISAM
  AUTO_INCREMENT = 100000000
  DEFAULT CHARSET = utf8;

CREATE TABLE `link_clicks` (
  `link_id`    INT(10) UNSIGNED NOT NULL,
  `clicked`    DATETIME         NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `ip_address` VARCHAR(191)     NOT NULL,
  KEY `link_clicks_link_index` (`link_id`),
  CONSTRAINT `link_clicks_link_fk` FOREIGN KEY (`link_id`) REFERENCES `links` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE
)
  ENGINE = MYISAM
  DEFAULT CHARSET = utf8;
