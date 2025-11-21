-- Add a new column 'status' to posts
ALTER TABLE posts
ADD COLUMN status VARCHAR(50) NOT NULL DEFAULT 'draft';

-- Ensure updated_at is automatically updated on row modification
-- First, create a function to update updated_at
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Then, create a trigger on the posts table
DROP TRIGGER IF EXISTS trigger_update_posts_updated_at ON posts;

CREATE TRIGGER trigger_update_posts_updated_at
BEFORE UPDATE ON posts
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();
