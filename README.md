# HW3
If you made changes to the front folder make sure to run 
```
cd front/banana-notes
npm run build
```
Then just do 
```
docker-compose build
docker-compose up --scale server=3
```
The `--scale server=3` runs 3 servers to achieve load balancing.