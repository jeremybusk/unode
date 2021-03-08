# Universal Node

“The intellect of man is forced to choose
Perfection of the life, or of the work
And if it take the second must refuse
A heavenly mansion, raging in the dark.”

― William Butler Yeats

If you want to enjoy yourself have a BBQ or grab a brew. Perhaps get the boat and hit the lake or binge watch another show.
We are trying to do some really hard modeling of our universe here. This will be a long process. Please join the cause.
The only way to get out of this rats maze is to understand it. One of the best ways to understand something is to model it.

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
# or lets use nginx passthrough
curl -k https://$host/api/todos
curl -k https://$host/api/todos -X POST -H "Authorization: Bearer $TOKEN" -H "Content-Type: application/json" -d '{"task": "New task"}'
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
