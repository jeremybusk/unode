-- sudo -u postgres psql -d unode < api.postgrest.sql
create role authenticator noinherit login password 'mysecretpassword';
create role web_anon nologin;
grant web_anon to authenticator;

create schema api;
create table api.todos (
  id serial primary key,
  done boolean not null default false,
  task text not null,
  due timestamptz
);

insert into api.todos (task) values
  ('finish tutorial 0'), ('pat self on back');

grant usage on schema api to web_anon;
grant select on api.todos to web_anon;

create role todo_user nologin;
grant todo_user to authenticator;

grant usage on schema api to todo_user;
grant all on api.todos to todo_user;
grant usage, select on sequence api.todos_id_seq to todo_user;
