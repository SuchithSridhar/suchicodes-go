CREATE TABLE ContactMessage (
	id UUID NOT NULL, 
	date DATETIME,
    email VARCHAR,
	subject VARCHAR, 
	message VARCHAR, 
	ip_address VARCHAR, 
	PRIMARY KEY (id)
);

CREATE TABLE IpLog (
	id UUID NOT NULL, 
	timestamp DATETIME, 
	url VARCHAR, 
	ip_address VARCHAR, 
	platform VARCHAR, 
	reference VARCHAR, 
	user_agent VARCHAR, 
	PRIMARY KEY (id)
);

CREATE TABLE URLRedirection (
    id UUID NOT NULL,
    alias VARCHAR,
    dst_url VARCHAR,
    PRIMARY KEY(id)
);

CREATE TABLE Server (
    id UUID NOT NULL,
    name VARCHAR,
    location VARCHAR,
    ip_address VARCHAR,
    status VARCHAR,
    status_detail JSON,
    PRIMARY KEY(id)
)

CREATE TABLE ServerLog (
    id UUID NOT NULL,
    server UUID NOT NULL,
    timestamp DATETIME,
    detail JSON,
    PRIMARY KEY(id)
    FOREIGN KEY(server) REFERENCES Server(id)
);


CREATE TABLE Project (
	id UUID NOT NULL, 
	title VARCHAR, 
	html TEXT, 
	markdown TEXT, 
	url VARCHAR, 
	brief TEXT, 
	icon_img VARCHAR, 
	PRIMARY KEY (id)
);

CREATE TABLE Category (
    id UUID NOT NULL,
    parent UUID,
    order INTEGER NOT NULL DEFAULT 0,
    name VARCHAR,
    PRIMARY KEY(id)
    FOREIGN KEY(parent) REFERENCES Category (id)
);

CREATE TABLE User (
    id UUID NOT NULL,
    username VARCHAR NOT NULL,
    password BINARY(60) NOT NULL,
    created DATETIME,
    PRIMARY KEY (id),
    UNIQUE (username)
);

CREATE TABLE Blog (
	id UUID NOT NULL, 
	category UUID NOT NULL, 
	created DATETIME, 
	updated DATETIME, 
    author UUID NOT NULL,
	title VARCHAR, 
	html TEXT,
	markdown TEXT,
	brief TEXT,
	date_updated DATETIME, 
	PRIMARY KEY (id), 
	FOREIGN KEY(category) REFERENCES Category (id)
)

-- Use this table to keep track of uploaded images for blogs
CREATE TABLE ImageMap (
    id UUID NOT NULL,
    blog UUID NOT NULL,
    original_name VARCHAR,
    generated_name VARCHAR,
    image_hash BINARY(32)
    PRIMARY KEY(id),
    FOREIGN KEY(blog) REFERENCES Blog(id)
);
