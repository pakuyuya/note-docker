CREATE SCHEMA sample;

CREATE TABLE sample.musers (
  id BIGINT PRIMARY KEY,
  loginid VARCHAR(60),
  loginpw VARCHAR(32),
  expiredt DATE,
  disabled BOOL,
  createat TIMESTAMP,
  updateat TIMESTAMP,
  upduser  VARCHAR(16)
);

INSERT INTO
  sample.musers (id, loginid, loginpw, expiredt, disabled, createat, updateat, upduser)
  VALUES
    (1, 'alpha', md5('alpha'), CURRENT_DATE + 365, false, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 'init-script'),
    (2, 'alpha', md5('alpha'), CURRENT_DATE, false, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 'init-script'),
    (3, 'alpha', md5('alpha'), CURRENT_DATE - 1, false, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 'init-script'),
    (4, 'alpha', md5('alpha'), CURRENT_DATE - 300, false, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 'init-script')
    ;
