CREATE TABLE IF NOT EXISTS `dogs` (
    `id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY ,
    `name` VARCHAR(100) NOT NULL ,
    `age` TINYINT NOT NULL ,
    `breed` VARCHAR(100) NOT NULL
) ENGINE = InnoDB;