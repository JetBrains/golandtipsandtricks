CREATE TABLE IF NOT EXISTS visitors(
	id serial constraint visitors_pk primary key,
	user_agent varchar(250),
	datetime timestamp
);

CREATE INDEX IF NOT EXISTS visitors_user_agent_index ON visitors (user_agent);