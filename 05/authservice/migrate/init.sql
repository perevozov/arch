CREATE TABLE IF NOT EXISTS user (
    id int not null auto_increment,
    username varchar(256) not null,
    first_name varchar(256) not null,
	last_name varchar(256) not null,
	email varchar(256),
	phone varchar(256),
    password_hash varchar(256),

    PRIMARY KEY (id),
    UNIQUE INDEX ux_user_01(username)
);

