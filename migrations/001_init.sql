CREATE TABLE voters (
    id SERIAL PRIMARY KEY,
    nim VARCHAR(50) UNIQUE NOT NULL,
    qrimage TEXT
);

CREATE TABLE candidates (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);

CREATE TABLE votes (
    id SERIAL PRIMARY KEY,
    voter_nim VARCHAR(50) NOT NULL,
    candidate_id INT NOT NULL,
    hash TEXT NOT NULL
);
