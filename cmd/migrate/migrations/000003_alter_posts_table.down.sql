-- Remove the 'status' column
ALTER TABLE posts
DROP COLUMN IF EXISTS status;

-- Drop the trigger and function
DROP TRIGGER IF EXISTS trigger_update_posts_updated_at ON posts;

DROP FUNCTION IF EXISTS update_updated_at_column();
