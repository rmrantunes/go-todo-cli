-- +goose Up
-- +goose StatementBegin
CREATE TABLE todo (
  id INTEGER PRIMARY KEY,
  description VARCHAR(256) NOT NULL,
  done BOOLEAN DEFAULT false NOT NULL,
  created_at TIMESTAMP NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE todo;
-- +goose StatementEnd
