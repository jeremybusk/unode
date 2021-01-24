# unode-server
Unode Server Software

# Install on Ubuntu 20.04
```
git clone https://github.com/jeremybusk/unode-server.git
cd unode-server
./install install-all myadminusername myadminuserpass
```

# One liner
```
cd ~/; rm -rf unode-server; git clone https://github.com/jeremybusk/unode-server.git; cd unode-server; chmod +x install; ./install install-all myadminusername myadminuserpass
```

You should now be able to go to these pages using webclient if installed correctly. Replace $host with your ipaddr or hostname
```
https://$host/health
https://$host/pgadmin4/
https://$host/get/client-ipaddr
curl -k https://$host/sign
curl -k https://$host/verify
export TOKEN="MYTOKEN"
curl http://$host:3000/todos
curl http://$host:3000/todos -X POST -H "Authorization: Bearer $TOKEN" -H "Content-Type: application/json" -d '{"task": "New task"}'
```

Command-line
```
psql -h $host -p 5432 -d unode -U myadminusername 
psql -h $host -p 44441 -d unode -U myadminuserpass 
```

More urls
```
https://$host/fw
```
