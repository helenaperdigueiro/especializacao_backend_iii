CREATE SCHEMA `avaliacao-ii`;

CREATE TABLE `avaliacao-ii`.`dentist` (
	`id` INT NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(100) NOT NULL,
    `last_name` VARCHAR(100)  NOT NULL,
    `enrollment` VARCHAR(100)  NOT NULL,
    PRIMARY KEY (`id`)
);
    
CREATE TABLE `avaliacao-ii`.`patient` (
	`id` INT NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(100) NOT NULL,
    `last_name` VARCHAR(100)  NOT NULL,
    `rg` VARCHAR(100)  NOT NULL,
    `register_date` VARCHAR(100)  NOT NULL,
    PRIMARY KEY (`id`)
);

CREATE TABLE `avaliacao-ii`.`appointment` (
	`id` INT NOT NULL AUTO_INCREMENT,
    `patient_id` INT NOT NULL,
    `dentist_id` INT  NOT NULL,
    `date` VARCHAR(100)  NOT NULL,
    `description` VARCHAR(100)  NOT NULL,
    PRIMARY KEY (`id`),
		FOREIGN KEY (`patient_id`)
        REFERENCES `avaliacao-ii`.`patient` (`id`),
        FOREIGN KEY (`dentist_id`)
        REFERENCES `avaliacao-ii`.`dentist` (`id`)
);

INSERT INTO `avaliacao-ii`.`dentist` (`name`, `last_name`, `enrollment`)
VALUES ('Helena', 'Perdigueiro', '123');

INSERT INTO `avaliacao-ii`.`patient` (`name`, `last_name`, `rg`, `register_date`)
VALUES ('Carol', 'Haka', '123', '12/12/2022');

INSERT INTO `avaliacao-ii`.`appointment` (`patient_id`, `dentist_id`, `date`, `description`)
VALUES (1, 1, '13/12/2022', 'Avocado!appointmentappointment');