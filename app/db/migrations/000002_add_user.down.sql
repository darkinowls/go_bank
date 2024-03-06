-- Drop the foreign key constraint first to avoid errors
ALTER TABLE "accounts" DROP CONSTRAINT IF EXISTS "accounts_owner_fkey";

-- Drop the index
DROP INDEX IF EXISTS "unique_owner_currency_idx";

-- Drop the users table
DROP TABLE IF EXISTS "users";
