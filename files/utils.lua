local utils = {}


function utils.pgcon()
    local pgmoon = require("pgmoon")
    local pg = pgmoon.new({
        -- host = "pyroportal.pyrofex.io",
        host = "10.64.41.6",
        port = "5432",
        database = "pyroportal",
        password = "pyroportal1in",
        ssl = true,
        user = "pyroportal"
    })
    assert(pg:connect())
    return pg
end


function utils.rediscon()
    local redis = require "nginx.redis"
    local rdb = redis:new()
    rdb:set_timeout(1000) -- 1 sec
    local ok, err = rdb:connect("127.0.0.1", 6379)
    if not ok then
        ngx.say("failed to connect: ", err)
        return
    end
    local res, err = rdb:auth("redis1in")
    if not res then
        ngx.say("failed to authenticate: ", err)
        return
    end
    return rdb
end


function utils.get_header_val(var_name)
    -- gets value from http headers in format of my_var: somevalue
    retval = nil
    local h = ngx.req.get_headers()
    for k, v in pairs(h) do
        if k == var_name then
            retval = v
        end
    end
    return retval
end


return utils
