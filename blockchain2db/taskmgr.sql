CREATE DATABASE IF NOT EXISTS octopus;

USE octopus;

CREATE TABLE IF NOT EXISTS `write_set` (
    `event_id` varchar(128) NOT NULL,
    `namespace` varchar(64) NOT NULL,
    `key` varchar(256) NOT NULL,
    `value` text NOT NULL,
    `create_time` datetime NOT NULL,
    `tx_id` varchar(64) NOT NULL,
    `is_delete` tinyint(1) NOT NULL,
    `creator` varchar(64) DEFAULT NULL,
    PRIMARY KEY (`event_id`,`key`,`namespace`),
    KEY `key` (`key`),
    KEY `namespace` (`namespace`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `taskmgr` (
    `key` varchar(256) NOT NULL,
    `task_name` varchar(256) NOT NULL,
    `creator` varchar(256) DEFAULT NULL,
    `require` varchar(1024) DEFAULT NULL,
    `approved` varchar(1024) DEFAULT NULL,
    `description` varchar(1024) DEFAULT NULL,
    `update_time` DATETIME,
    PRIMARY KEY (`key`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

DROP TRIGGER IF EXISTS octopus_insert_rewrite;

DELIMITER $$
CREATE TRIGGER octopus_insert_rewrite BEFORE INSERT ON `write_set`
FOR EACH ROW
BEGIN
    SET NEW.value = FROM_BASE64(NEW.value);
    INSERT INTO taskmgr 
		(`key`, `task_name`, `creator`, `require`, `approved`, `description`, `update_time`)
	VALUES (
		NEW.key,
		NEW.value->>'$.name',
        NEW.value->>'$.creator',
        NEW.value->>'$.requires',
        NEW.value->>'$.approved',
        NEW.value->>'$.description',
        NOW()
	) ON DUPLICATE KEY UPDATE
		`task_name`=NEW.value->>'$.name',
		`creator`=NEW.value->>'$.creator',
		`require`=NEW.value->>'$.requires',
		`approved`=NEW.value->>'$.approved',
		`description`=NEW.value->>'$.description',
        `update_time`=NOW();
END; $$
DELIMITER ;
