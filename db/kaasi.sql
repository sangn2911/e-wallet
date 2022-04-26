-- docker exec -i mysql_container mysql -uroot -proot kaasi < kaasi.sql
-- Add sql file from local computer to docker mysql database
drop table if exists `user`;

drop table if exists `customer`;

drop table if exists `document`;

drop table if exists `transaction`;

drop table if exists `affiliate`;

CREATE TABLE `user` (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(128) NOT NULL UNIQUE,
  `email` varchar(128) NOT NULL,
  `password` varchar(255) NOT NULL,
  PRIMARY KEY (`id`)
);

CREATE TABLE `customer` (
  `id` int NOT NULL AUTO_INCREMENT,
  `firstName` varchar(128) NOT NULL,
  `lastName` varchar(128) NOT NULL,
  `dateOfBirth` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `nationality` varchar(255) NOT NULL,
  `address` varchar(255) NOT NULL,
  PRIMARY KEY (`id`)
);

CREATE TABLE `document` (
  `id` int NOT NULL AUTO_INCREMENT,
  `docType` varchar(128) NOT NULL,
  `docNumber` varchar(128) NOT NULL,
  `issuingAuthority` varchar(255) NOT NULL,
  `expiryDate` varchar(255) NOT NULL,
  `img` varchar(255) NOT NULL,
  PRIMARY KEY (`id`)
);

CREATE TABLE `transaction` (
  `id` int NOT NULL AUTO_INCREMENT,
  `senderName` varchar(128) NOT NULL,
  `receiverName` varchar(128) NOT NULL,
  `date` varchar(255) NOT NULL,
  `money` varchar(255) NOT NULL,
  `message` varchar(255) NOT NULL,
  PRIMARY KEY (`id`)
);

CREATE TABLE `affiliate` (
  `id` int NOT NULL AUTO_INCREMENT,
  `affiname` varchar(128) NOT NULL,
  `district` varchar(128) NOT NULL,
  `address` varchar(255) NOT NULL,
  `phoneNumber` varchar(255) NOT NULL,
  `fax` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  PRIMARY KEY (`id`)
);

INSERT INTO
  user (username, email, password)
VALUES
  ('user1', 'user1@gmail.com', 'pass1'),
  ('user2', 'user2@gmail.com', 'pass2'),
  ('user3', 'user3@gmail.com', 'pass3');

INSERT INTO
  affiliate (
    affiname,
    district,
    address,
    phoneNumber,
    fax,
    email
  )
VALUES
  (
    'Ho Chi Minh',
    'District 10',
    'addr1',
    '0123456789',
    'fax1',
    'email1'
  ),
  (
    'Ho Chi Minh',
    'District 11',
    'addr2',
    '0123456789',
    'fax2',
    'email2'
  ),
  (
    'Ho Chi Minh',
    'District 1',
    'addr3',
    '0123456789',
    'fax3',
    'email3'
  );