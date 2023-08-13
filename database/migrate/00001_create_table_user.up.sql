CREATE TABLE accounts
(
    user_id     bigserial PRIMARY KEY,
    username    VARCHAR(255) UNIQUE                              NOT NULL,
    password    VARCHAR(255)                                     NOT NULL,
    email       VARCHAR(255) UNIQUE                              NOT NULL,
    last_login  TIMESTAMP default now(),
    is_active   BIT          default B'1'::"bit"                 not null,
    created_by  VARCHAR(255) default 'SYSTEM'::character varying not null,
    created_at  TIMESTAMP    default now()                       not null,
    modified_by VARCHAR(255) default 'SYSTEM'::character varying not null,
    modified_at TIMESTAMP    default now()                       not null,
    deleted_by  varchar(255),
    deleted_at  TIMESTAMP
);

CREATE INDEX idx_username ON accounts(username);
CREATE INDEX idx_email ON accounts(email);