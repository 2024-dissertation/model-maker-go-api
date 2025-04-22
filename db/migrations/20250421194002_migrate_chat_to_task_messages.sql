-- +goose Up
-- +goose StatementBegin

-- Drop foreign key constraint from chat_messages to chats
ALTER TABLE chats
DROP CONSTRAINT IF EXISTS fk_chat_task;

ALTER TABLE chat_messages
DROP CONSTRAINT IF EXISTS fk_message_chat;

-- Drop the chats table
DROP TABLE IF EXISTS chats;

-- Add task_id to chat_messages
ALTER TABLE chat_messages
ADD COLUMN task_id INTEGER NOT NULL;

-- Add foreign key constraint from chat_messages to tasks
ALTER TABLE chat_messages
ADD CONSTRAINT fk_chat_messages_task FOREIGN KEY (task_id) REFERENCES tasks(id) ON DELETE CASCADE;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- Drop the foreign key constraint and column
ALTER TABLE chat_messages
DROP CONSTRAINT IF EXISTS fk_chat_messages_task;

ALTER TABLE chat_messages
DROP COLUMN IF EXISTS task_id;

-- Recreate the chats table
CREATE TABLE chats (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    task_id INTEGER NOT NULL,
    CONSTRAINT fk_chat_task FOREIGN KEY (task_id) REFERENCES tasks(id) ON DELETE CASCADE
);

-- Restore the old foreign key if needed
ALTER TABLE chat_messages
ADD COLUMN chat_id INTEGER;

ALTER TABLE chat_messages
ADD CONSTRAINT chat_messages_chat_id_fkey FOREIGN KEY (chat_id) REFERENCES chats(id) ON DELETE CASCADE;

-- +goose StatementEnd
