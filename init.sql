CREATE TABLE IF NOT EXISTS original_links (
                                              id SERIAL PRIMARY KEY,
                                              original_link VARCHAR(250) NOT NULL,
                                              new_link VARCHAR(250) NOT NULL

);

CREATE TABLE IF NOT EXISTS new_links (
                                         id SERIAL PRIMARY KEY,
                                         new_link VARCHAR(250) NOT NULL,
                                         original_link VARCHAR(250) NOT NULL

);