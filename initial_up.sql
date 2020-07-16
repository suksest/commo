USE commo_db;

CREATE TABLE commo_db.user(
    id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    phone VARCHAR(15) UNIQUE NOT NULL,
    role ENUM("admin", "user") NOT NULL,
    password VARCHAR(50) NOT NULL
);

INSERT INTO commo_db.user (name, phone, role, password) VALUES('Jonas Kahnwald', '081321456876', 'user', MD5('1BwP'));
INSERT INTO commo_db.user (name, phone, role, password) VALUES('Martha Nielsen', '081111111111', 'user', MD5('2CrF'));
INSERT INTO commo_db.user (name, phone, role, password) VALUES('Hanno Tauber', '084554805480', 'admin', MD5('5FbN'));
INSERT INTO commo_db.user (name, phone, role, password) VALUES('Hannah Kruger', '082442804280', 'admin', MD5('W4rL'));
