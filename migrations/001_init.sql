-- =====================================================
-- INIT DATABASE E-VOTING
-- File: 001_init.sql
-- DB  : PostgreSQL
-- =====================================================

-- ==============================
-- 1. TABLE: candidates
-- ==============================
CREATE TABLE IF NOT EXISTS candidates (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- ==============================
-- 2. TABLE: voters
-- ==============================
CREATE TABLE IF NOT EXISTS voters (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    token_hash VARCHAR(255) UNIQUE NOT NULL,
    has_voted BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- ==============================
-- 3. TABLE: votes
-- ==============================
CREATE TABLE IF NOT EXISTS votes (
    id SERIAL PRIMARY KEY,
    voter_id INT NOT NULL,
    candidate_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_voter
        FOREIGN KEY (voter_id)
        REFERENCES voters(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_candidate
        FOREIGN KEY (candidate_id)
        REFERENCES candidates(id)
        ON DELETE CASCADE
);

-- ==============================
-- 4. INDEXES (Performance)
-- ==============================
CREATE INDEX IF NOT EXISTS idx_voters_token_hash ON voters(token_hash);
CREATE INDEX IF NOT EXISTS idx_votes_voter_id ON votes(voter_id);
CREATE INDEX IF NOT EXISTS idx_votes_candidate_id ON votes(candidate_id);

-- ==============================
-- 5. UNIQUE CONSTRAINT (ANTI DOUBLE VOTE)
-- ==============================
CREATE UNIQUE INDEX IF NOT EXISTS unique_voter_vote
ON votes(voter_id);

-- =====================
-- 6. ADMIN DEFAULT
-- =====================
INSERT INTO admins (username, password, created_at, updated_at)
VALUES (
  'admin',
  'adminpassword',
  NOW(),
  NOW()
);
-- =====================================================
-- END INIT
-- =====================================================
