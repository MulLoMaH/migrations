
ALTER TABLE products ADD COLUMN quality BOOLEAN;
ALTER TABLE products ADD CONSTRAINT title_unique UNIQUE(title);