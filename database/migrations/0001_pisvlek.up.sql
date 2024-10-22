CREATE TABLE users (
	id	INTEGER NOT NULL UNIQUE,
	email TEXT NOT NULL,
	full_name TEXT NOT NULL,
	password TEXT NOT NULL,
	PRIMARY KEY(id)
);

-- WW 'pisvlek', puur voor het voorbeeld, niet in prod zo doen
INSERT INTO users (email, full_name, password) VALUES ('pisvlek@jemoeder.com', 'Pisvlek Deluxe', '$2a$10$qBousQaWleTqqRMMglKdYOrIHOBVgYcJrGU8gWRa1Rb7YCCUnY3Ty');