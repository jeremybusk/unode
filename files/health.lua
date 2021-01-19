local config = require("/etc/nginx/lua/config")
local driver = require "luasql.postgres"
env = assert (driver.postgres())
local sqlstr = "dbname=" .. config.test_pgname .. " host=" .. config.pghost .. " port=" .. config.pgport .. " user=" .. config.pguser ..  " password=" .. config.pgpass
print(sqlstr)
con = assert (env:connect(sqlstr))
res = con:execute"DROP TABLE people"
res = assert (con:execute[[
  CREATE TABLE people(
    name  varchar(50),
    email varchar(50)
  )
]])
list = {
  { name="Jose das Couves", email="jose@couves.com", },
  { name="Manoel Joaquim", email="manoel.joaquim@cafundo.com", },
  { name="Maria das Dores", email="maria@dores.com", },
}
for i, p in pairs (list) do
  res = assert (con:execute(string.format([[
    INSERT INTO people
    VALUES ('%s', '%s')]], p.name, p.email)
  ))
end
cur = assert (con:execute"SELECT name, email from people")
row = cur:fetch ({}, "a")
while row do
  print(string.format("Name: %s, E-mail: %s", row.name, row.email))
  row = cur:fetch (row, "a")
end
cur:close() -- already closed because all the result set was consumed
con:close()
env:close()


require("socket")
local https = require("ssl.https")
local url = "https://example.org/"
local body, code, headers, status = https.request(url)
-- local body, code, headers, status = https.request{url = "https://www.google.com", protocol = "tlsv1_3"}
if code ~= 200 then
  ngx.say("FAIL: HTTP connection to " .. url .. " failed!" )
else
  ngx.say("PASS: HTTP connection to " .. url .. " succeeded!" )
end
-- ngx.say(body)
-- ngx.say(headers)
ngx.say(status)
ngx.status = ngx.HTTP_OK
ngx.say("PASS: All tests passed!")
return ngx.exit(ngx.HTTP_OK)
