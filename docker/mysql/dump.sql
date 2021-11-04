-- Use getground;

CREATE TABLE `guest_table` (
  `table_number` INT NOT NULL,
  `space` int,
  PRIMARY KEY (`table_number`)
);

CREATE TABLE `guest_info` (
  `id` int auto_increment,
  `table_number` int,
  `name` varchar(64),
  `accompanying_guests` int,
  `time_arrived` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
);

insert into guest_table(table_number, space) values(1, 10);
insert into guest_table(table_number, space) values(2, 10);
insert into guest_table(table_number, space) values(3, 10);
insert into guest_table(table_number, space) values(4, 10);

insert into guest_info(table_number, name,accompanying_guests) values(1, "boris", 2);
insert into guest_info(table_number, name,accompanying_guests) values(1, "nina", 3);

