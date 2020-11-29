	-- Create tables.
	DROP TABLE IF EXISTS "usersList";
	DROP TABLE IF EXISTS "forums";
	DROP TABLE IF EXISTS "interestList";
	DROP TABLE IF EXISTS "users";
	CREATE TABLE "users"
	(
		"id"   SERIAL PRIMARY KEY,
		"name" VARCHAR(50) NOT NULL UNIQUE,
		"password" VARCHAR(50) NOT NULL
	);

	CREATE TABLE "interestList"
	(
		"id"   SERIAL PRIMARY KEY,
		"interest" VARCHAR(50) NULL,
		"userID" INT NOT NULL,
		FOREIGN KEY ("userID")
		REFERENCES users(id)
			ON DELETE CASCADE
			ON UPDATE CASCADE
	);

	CREATE TABLE "forums"
	(
		"id"   SERIAL PRIMARY KEY,
		"name" VARCHAR(50) NOT NULL UNIQUE,
		"topicKeyword" VARCHAR(50) NOT NULL UNIQUE
	);

	CREATE TABLE "usersList"
	(
		"id"   SERIAL PRIMARY KEY,
		"forumsID" INT NOT NULL,
		FOREIGN KEY ("forumsID")
		REFERENCES forums(id)
			ON DELETE CASCADE
			ON UPDATE CASCADE,
		"userID" INT NOT NULL,
		FOREIGN KEY ("userID")
		REFERENCES users(id)
			ON DELETE CASCADE
			ON UPDATE CASCADE
	);

	-- Insert demo data.
	INSERT INTO "users" (name, password) VALUES ('Dania', 'Password');
	INSERT INTO "users" (name, password) VALUES ('Dania0', 'Password');
	INSERT INTO "users" (name, password) VALUES ('Dania1', 'Password');
	INSERT INTO "interestList" ("interest", "userID") VALUES ('Games', 1);
	INSERT INTO "interestList" ("interest", "userID") VALUES ('Books', 1);
	INSERT INTO "interestList" ("interest", "userID") VALUES ('Games', 2);
	INSERT INTO "forums" (name, "topicKeyword") VALUES ('Game fan', 'Games');
	INSERT INTO "forums" (name, "topicKeyword") VALUES ('Book enjoyer', 'Books');
	INSERT INTO "usersList" ("forumsID", "userID") VALUES (1, 1);
	INSERT INTO "usersList" ("forumsID", "userID") VALUES (2, 1);
	INSERT INTO "usersList" ("forumsID", "userID") VALUES (1, 2);

    
-- select
-- 	forums.name, users.name
-- from
-- 	forums
-- left join
-- 	"usersList"
-- on
-- 	"usersList"."forumsID" = "forums".id
-- left join
-- 	users
-- on
-- 	users.id = "usersList"."userID"
-- where
-- 	forums.id = 1
-- GROUP BY
-- 	forums.name, users.name