CREATE TABLE IF NOT EXISTS Task(
   Id serial PRIMARY KEY,
   Name VARCHAR (200) NOT NULL,
   EventID VARCHAR(200) NOT NULL,
   Details Text,
   Created_At timestamp NOT NULL DEFAULT NOW(),
   Updated_At timestamp NOT NULL,
   Deleted_At timestamp NOT NULL
); 