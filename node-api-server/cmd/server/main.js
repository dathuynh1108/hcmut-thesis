const { newGrpcServer } = require("../../internal/server/interface/grpc/server");
const { newConsumers } = require("../../internal/server/interface/rabbitmq/consumers");
const { newMongoDB } = require("../../internal/server/database/mongodb");


async function main() {
    const db = newMongoDB();
    const grpcServer = newGrpcServer();
    const consumers = newConsumers();
    db.start();
    grpcServer.start();
    consumers.start();



    // const { newCameraStreamInfoHandler } = require("../../internal/server/grpc_client/handler/camera_stream_info_handler");
    // const handler = newCameraStreamInfoHandler();
    // handler.createCameraStream({})
}

main();



// const { newReportRepository } = require("../../internal/server/repository/report_repository");
// async function main1() {
//     const db = newMongoDB();
//     await db.start();

//     let repo = newReportRepository();
//     repo.findNumberOfIotEventWithAreaAndTime();
// }

// main1();










// // const app = require("../../internal/server/interface/http/app");
// const config = require("../../pkg/config/config");

// // const httpApp = app.NewHTTPApp();


// const { GRPCServer } = require("../../internal/server/interface/grpc/server");

// const grpcServer = new GRPCServer();
// grpcServer.start();




// const { MongoDb } = require("../../internal/server/database/mongodb");
// const mongo = new MongoDb();
// mongo.start();

















// var PROTO_PATH = __dirname + '/../../protos/example/example.proto';
// var grpc = require('@grpc/grpc-js');
// var protoLoader = require('@grpc/proto-loader');
// var packageDefinition = protoLoader.loadSync(
//     PROTO_PATH,
//     {keepCase: true,
//      longs: String,
//      enums: String,
//      defaults: true,
//      oneofs: true
//     });
// var protoDescriptor = grpc.loadPackageDefinition(packageDefinition);
// const grpcServer = new grpc.Server();
// let message = {
//   name: 'legiahuy',
//   num: 2
// }


// grpcServer.addService(protoDescriptor.Example.service, {
//   get: (_, callback) => {
//     console.log(999999999999999)
//     callback(null, message);
//   },
// });


// grpcServer.bindAsync(
//   "localhost:3000",
//   grpc.ServerCredentials.createInsecure(),
//   (error, port) => {
//     console.log(`Server is running at http://localhost:3000`);
//     grpcServer.start();
//   }
// );








// const grpc = require("@grpc/grpc-js");
// const protoLoader = require("@grpc/proto-loader");
// const PROTO_PATH = __dirname + "/password.proto";


// const loaderOptions = {
//     keepCase: true,
//     longs: String,
//     enums: String,
//     defaults: true,
//     oneofs: true,
// };

// var packageDef = protoLoader.loadSync(PROTO_PATH, loaderOptions);

// const grpcObj = grpc.loadPackageDefinition(packageDef);
// const ourServer = new grpc.Server();

// let dummyRecords = {
//     "passwords": [
//         { id: "153642", password: "default1", hashValue: "default", saltValue:       "default" },
//         { id: "234654", password: "default2", hashValue: "default", saltValue: "default" }]
// };

// ourServer.addService(grpcObj.PasswordService.service, {
//     /*our protobuf message(passwordMessage) for the RetrievePasswords was Empty. */
//     retrievePasswords: (passwordMessage, callback) => {
//         console.log(9999999999999999999999)
//         callback(null, dummyRecords);
//     },
//     addNewDetails: (passwordMessage, callback) => {
//         const passwordDetails = { ...passwordMessage.request };
//         dummyRecords.passwords.push(passwordDetails);
//         callback(null, passwordDetails);
//     },
//     updatePasswordDetails: (passwordMessage, callback) => {
//         const detailsID = passwordMessage.request.id;
//         const targetDetails = dummyRecords.passwords.find(({ id }) => detailsID == id);
//         targetDetails.password = passwordMessage.request.password;
//         targetDetails.hashValue = passwordMessage.request.hashValue;
//         targetDetails.saltValue = passwordMessage.request.saltValue;
//         callback(null, targetDetails);
//     },
// });



// ourServer.bindAsync(
//     "0.0.0.0:50051",
//     grpc.ServerCredentials.createInsecure(),
//     (error, port) => {
//         console.log("Server running at http://127.0.0.1:50051");
//         ourServer.start();
//     }
// );









// const http = require('http');
// const server = http.createServer(httpApp);
// const { Server } = require("socket.io");
// const io = new Server(server, { cors: { origin: "*" } });

// io.on('connection', (socket) => {
//   console.log('a user connected');
// });



// const { MQTTSubscriber } = require("../../internal/server/interface/mqtt/subcriber");
// const subscriber = new MQTTSubscriber(io);
// subscriber.start();


// server.listen(config.config.server.port, () =>
//   console.log(
//     `Server is running on http://127.0.0.1:${config.config.server.port}`
//   )
// );
