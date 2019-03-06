# ASU-drug-monitoring-web-server
This is a Smart India Hackathon 2019 problem statement which basically is to implement a web portal with an efficient server to montior the overall lifecycle of a newly released drug to help provide better information about it in the community as well as make it conveyable to national PV (Pharmacovigilance) resource centres such as WHO and the NCC.

File strucutre overview :

Monster :
 This file contains my gRPC client-server implementation where my client is just another service creating and requesting data from my server service which access data from the mongoDB database. 
 
Mnstrest :
 This file contains my microservice with rest api implementation which parses the incoming JSON data sent though individual web applications by HTTP and using this data to process specified requests and then storing resultant data in MongoDB in bson format.
URI's are predefined and based on the user's demand, respective URI's are accessed and the rest api then executes the respective operations and processes data.

What I am trying to do is build a basic web application which goes through my rest gateway and directly communicates with the database with the help of my services and to communicate between other services im using gRPC for faster data streaming.
