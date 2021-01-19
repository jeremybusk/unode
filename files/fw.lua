local ipaddr = ngx.var.remote_addr
local cmd = "sudo /app/unode/bin/fw-update.sh fw-allow-ip " .. ipaddr
os.execute(cmd)
ngx.status = ngx.HTTP_OK
ngx.say("I: ipaddr " .. ipaddr .. " allowed.")
return ngx.exit(ngx.HTTP_OK)
