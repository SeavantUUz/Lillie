CREATE TABLE message (
  msgid bigint PRIMARY KEY ,
  seq bigint NOT NULL ,
  sourceid bigint NOT NULL ,
  targetid bigint NOT NULL ,
  operation SMALLINT NOT NULL ,
  content bytea NOT NULL ,
  ts timestamp DEFAULT CURRENT_TIMESTAMP
)