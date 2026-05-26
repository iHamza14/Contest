CREATE TABLE contests (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,

    contest_id TEXT UNIQUE NOT NULL,
    title TEXT NOT NULL,
    start_time TIMESTAMP,
    end_time TIMESTAMP
);

CREATE TABLE problems (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,

    problem_id TEXT UNIQUE NOT NULL,
    problem_title TEXT NOT NULL,
    problem_statement TEXT,
    time_limit_ms INT NOT NULL,
    memory_limit_mb INT NOT NULL,
    contest_ref BIGINT REFERENCES contests(id) ON DELETE CASCADE
);

CREATE TABLE test_cases (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,

    input_file TEXT NOT NULL,
    output_file TEXT NOT NULL,
    is_hidden BOOLEAN DEFAULT FALSE,
    problem_ref BIGINT REFERENCES problems(id) ON DELETE CASCADE
);