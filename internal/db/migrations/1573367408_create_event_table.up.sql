CREATE TABLE IF NOT EXISTS Event(
   Id serial PRIMARY KEY,
   Name VARCHAR (200) NOT NULL,
   Type VARCHAR (50) NOT NULL,
   Status VARCHAR (300) NOT NULL
   Details []byte VARCHAR,
   Created_At date NOT NULL,
   Updated_At date NOT NULL,
   Deleted_At date NOT NULL
);