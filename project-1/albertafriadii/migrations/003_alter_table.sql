alter table users drop deleted_at;
alter table users add deleted_at timestamp;