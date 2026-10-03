CREATE TABLE deployments (
    application TEXT NOT NULL,
    environment TEXT NOT NULL,
    version INTEGER NOT NULL,
    status TEXT NOT NULL,
    started_at DATETIME NOT NULL,
    finished_at DATETIME NULL,

    PRIMARY KEY (application, version)
);
