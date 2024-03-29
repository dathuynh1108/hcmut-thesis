const grpc = require("@grpc/grpc-js");
const { newProtoLoader } = require("../../../../pkg/grpc/proto/proto_loader");

const { socketIO } = require("../../socket_io/socket_io");
const { config } = require("../../../../pkg/config/config");
const { logger } = require("../../../../pkg/logger/logger");

const { newAreaHandler } = require("./handler/area_handler");
const { newCameraMapHandler } = require("./handler/camera_map_handler");
const { newCameraTypeHandler } = require("./handler/camera_type_handler");
const { newCameraHandler } = require("./handler/camera_handler");
const { newEventHandler } = require("./handler/event_handler");
const { newEventTypeHandler } = require("./handler/event_type_handler");
const { newIotDeviceMapHandler } = require("./handler/iot_device_map_handler");
const { newIotDeviceTypeHandler } = require("./handler/iot_device_type_handler");
const { newIotDeviceHandler } = require("./handler/iot_device_handler");
const { newSystemUtilityHandler } = require("./handler/system_utility_handler");

const { newReportHandler } = require("./handler/report_handler");
const { newUserHandler } = require("./handler/user_handler");
const { newCameraStreamInfoHandler } = require("./handler/camera_stream_info_handler");

const { newStreamConnectionHandler } = require("./handler/stream_connection_handler");

class GrpcServer {
    constructor() {
        this.server = new grpc.Server();
        this.protoLoader = newProtoLoader();
        this.conf = config.server.grpc;
        this.initService();
    }

}

GrpcServer.prototype.initService = function() {
    this.server.addService(this.protoLoader.getService('area.proto', 'AreaService'), newAreaHandler());
    this.server.addService(this.protoLoader.getService('camera_map.proto', 'CameraMapService'), newCameraMapHandler());
    this.server.addService(this.protoLoader.getService('camera_type.proto', 'CameraTypeService'), newCameraTypeHandler());
    this.server.addService(this.protoLoader.getService('camera.proto', 'CameraService'), newCameraHandler());
    this.server.addService(this.protoLoader.getService('event.proto', 'EventService'), newEventHandler());
    this.server.addService(this.protoLoader.getService('event_type.proto', 'EventTypeService'), newEventTypeHandler());
    this.server.addService(this.protoLoader.getService('iot_device_map.proto', 'IotDeviceMapService'), newIotDeviceMapHandler());
    this.server.addService(this.protoLoader.getService('iot_device_type.proto', 'IotDeviceTypeService'), newIotDeviceTypeHandler());
    this.server.addService(this.protoLoader.getService('iot_device.proto', 'IotDeviceService'), newIotDeviceHandler());
    this.server.addService(this.protoLoader.getService('system_utility.proto', 'SystemUtilityService'), newSystemUtilityHandler());

    this.server.addService(this.protoLoader.getService('user.proto', 'UserService'), newUserHandler());
    this.server.addService(this.protoLoader.getService('report.proto', 'ReportService'), newReportHandler());
    this.server.addService(this.protoLoader.getService('camera_stream_info.proto', 'CameraStreamInfoService'), newCameraStreamInfoHandler());

    this.server.addService(this.protoLoader.getService('stream_connection.proto', 'StreamConnectionService'), newStreamConnectionHandler());
}

GrpcServer.prototype.start = function(host=this.conf.host, port=this.conf.port) {
    this.server.bindAsync(
        `${this.conf.binding_ip_address}:${port}`,
        grpc.ServerCredentials.createInsecure(),
        (error, port) => {
            this.server.start();
            logger.info(`Grpc server is running at ${host}:${port}`);
        }
    );
}

function newGrpcServer() {
    return new GrpcServer();
}

module.exports.newGrpcServer = newGrpcServer;

