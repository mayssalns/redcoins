CREATE DATABASE IF NOT EXISTS oauth;
  
USE oauth;

CREATE TABLE users (
	id INTEGER NOT NULL AUTO_INCREMENT,
    email VARCHAR(25) not null,
    password text not null,
	PRIMARY KEY (id)
);

SET character_set_client = utf8;
SET character_set_connection = utf8;
SET character_set_results = utf8;
SET collation_connection = utf8_general_ci;

