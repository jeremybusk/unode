# unode-server
Unode Server Software

# Install on Ubuntu 20.04
```
git clone https://github.com/jeremybusk/unode-server.git
cd unode-server
install
./install install-all myadminusername myadminuserpass
```

# One liner
```
cd ~/; rm -rf unode-server; git clone https://github.com/jeremybusk/unode-server.git; cd unode-server; chmod +x install; ./install install-all myadminusername myadminuserpass
```

You should now be able to go to these pages using webclient if installed correctly. Replace host with your ipaddr or hostname
```
https://host/pgadmin4/
https://host/health
https://host/get/client-ipaddr
```
