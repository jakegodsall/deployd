CREATE TABLE deployments (
    id TEXT NOT NULL PRIMARY KEY,
    application TEXT NOT NULL,
    environment TEXT NOT NULL,
    version INTEGER NOT NULL,
    status TEXT NOT NULL,
    started_at DATETIME NOT NULL,
    finished_at DATETIME NULL,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,

    UNIQUE (application, version)
);
