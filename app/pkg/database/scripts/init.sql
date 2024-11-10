CREATE TABLE IF NOT EXISTS tbl_client (
    id TEXT NOT NULL,
    secret TEXT NOT NULL,
    name TEXT NOT NULL,
    confidential INTEGER NOT NULL DEFAULT FALSE,
    PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS tbl_redirect_uri (
    client_id TEXT NOT NULL,
    uri TEXT NOT NULL,
    FOREIGN KEY(client_id) REFERENCES tbl_client(id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS tbl_authorization_code (
    client_id TEXT NOT NULL,
    code TEXT NOT NULL,
    redirect_uri TEXT NOT NULL,
    redirect_uri_forced INTEGER NOT NULL DEFAULT FALSE,
    PRIMARY KEY(client_id, code),
    FOREIGN KEY(client_id) REFERENCES tbl_client(id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS tbl_access_token (
    client_id TEXT NOT NULL,
    token TEXT NOT NULL,
    expiration_date NOT NULL,
    PRIMARY KEY(client_id, token),
    FOREIGN KEY(client_id) REFERENCES tbl_client(id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS tbl_jwts (
    id TEXT NOT NULL,
    client_id TEXT NOT NULL,
    public_key BLOB NOT NULL,
    private_key BLOB NOT NULL,
    PRIMARY KEY(id),
    FOREIGN KEY(client_id) REFERENCES tbl_client(id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS tbl_api (
    id TEXT PRIMARY KEY,
    uri TEXT NOT NULL UNIQUE,
    name TEXT NOT NULL,
    signing_algorithm TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS tbl_user (
    id TEXT NOT NULL,
    email TEXT NOT NULL,
    password TEXT NOT NULL,
    PRIMARY KEY(id)
    UNIQUE(email)

);