CREATE TABLE user_account (
  id SERIAL NOT NULL,
  created_at TIMESTAMP NOT NULL,
  created_by INTEGER NOT NULL,
  updated_at TIMESTAMP,
  updated_by INTEGER,
  name TEXT,
  email TEXT,
  password TEXT,
  PRIMARY KEY(id)
);

ALTER TABLE user_account ADD CONSTRAINT user_account_uniq1 UNIQUE (uid) ;

COMMENT ON TABLE user_account IS 'ユーザーアカウント';
COMMENT ON COLUMN user_account.id IS 'アカウントID';
COMMENT ON COLUMN user_account.created_at IS '登録日時';
COMMENT ON COLUMN user_account.created_by IS '登録者';
COMMENT ON COLUMN user_account.updated_at IS '更新日時';
COMMENT ON COLUMN user_account.updated_by IS '更新者';
COMMENT ON COLUMN user_account.name IS 'アカウント名';
COMMENT ON COLUMN user_account.email IS 'メールアドレス';
COMMENT ON COLUMN user_account.password IS 'パスワード';