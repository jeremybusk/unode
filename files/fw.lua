local function osExecute(cmd)
  local fileHandle     = assert(io.popen(cmd, 'r'))
  local commandOutput  = assert(fileHandle:read('*a'))
  local returnTable    = {fileHandle:close()}
  return commandOutput,returnTable[3]            -- rc[3] contains returnCode
end

local ipaddr = ngx.var.remote_addr
local cmd = "sudo /app/unode/bin/fw-update fw-allow-ip " .. ipaddr

r = osExecute(cmd)
-- r = os.execute(cmd)

ngx.status = ngx.HTTP_OK
ngx.say("I: ipaddr " .. ipaddr .. " allowed.")
ngx.say(r)
return ngx.exit(ngx.HTTP_OK)
