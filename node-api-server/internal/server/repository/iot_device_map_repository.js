const IotDeviceMapModel = require("../model/iot_device_map_model");
const { newIotDeviceMap } = require("../entity/iot_device_map")
const { newFromDatabaseConverter } = require("../data_converter/from_database_converter");
const { newToDatabaseConverter } = require("../data_converter/to_database_converter");

const { newInternalServerError } = require("../entity/error/internal_server_error");


class IotDeviceMapRepository {
    constructor() {
        this.fromDatabaseConverter = newFromDatabaseConverter();
        this.toDatabaseConverter = newToDatabaseConverter();
    }

    async getAll() {
        let iotDeviceMapDocs;
        try {
            iotDeviceMapDocs = await IotDeviceMapModel.find({});
        }
        catch(err) {
            throw newInternalServerError("Database error", err);
        }
        return iotDeviceMapDocs.map(iotDeviceMapDoc => {
            return this.fromDatabaseConverter.visit(newIotDeviceMap(), iotDeviceMapDoc);
        })
    }

    async create(iotDeviceMapEntity) {
        const iotDeviceMapDoc = this.toDatabaseConverter.visit(iotDeviceMapEntity);
        let newIotDeviceMapDoc;
        try {
            newIotDeviceMapDoc = await IotDeviceMapModel.create(iotDeviceMapDoc);
        }
        catch(err) {
            throw newInternalServerError("Database error", err);
        }
        return this.fromDatabaseConverter.visit(newIotDeviceMap(), newIotDeviceMapDoc);
    }

    async findById(iotDeviceMapId) {
        let iotDeviceMapDoc;
        iotDeviceMapId = iotDeviceMapId.getValue();
        try {
            iotDeviceMapDoc = await IotDeviceMapModel.findById(iotDeviceMapId).exec();
        }
        catch(err) {
            throw newInternalServerError("Database error", err);
        }
        
        return this.fromDatabaseConverter.visit(newIotDeviceMap(), iotDeviceMapDoc);
    }

    async findByIdAndUpdate(iotDeviceMapId, iotDeviceMapEntity) {
        const iotDeviceMapDoc = this.toDatabaseConverter.visit(iotDeviceMapEntity);
        const filter = {
            _id: iotDeviceMapId.getValue()
        }
        let newIotDeviceMapDoc;
        try {
            newIotDeviceMapDoc = await IotDeviceMapModel.findOneAndUpdate(filter, iotDeviceMapDoc, { new: true }); // set new to true to return new document after update
        }
        catch(err) {
            throw newInternalServerError("Database error", err);
        }
        return this.fromDatabaseConverter.visit(newIotDeviceMap(), newIotDeviceMapDoc);
    }


    async findByIdAndDelete(iotDeviceMapId) {
        const filter = {
            _id: iotDeviceMapId.getValue()
        }
        let deleteIotDeviceMapDoc;
        try {
            deleteIotDeviceMapDoc = await IotDeviceMapModel.findOneAndDelete(filter); // set new to true to return new document after update
        }
        catch(err) {
            throw newInternalServerError("Database error", err);
        }
        return this.fromDatabaseConverter.visit(newIotDeviceMap(), deleteIotDeviceMapDoc);
    }

    async findByConnectIot(connectIotId) {
        let iotDeviceMapDoc;
        const filter = {
            connect_iot: connectIotId.getValue()
        }
        try {
            iotDeviceMapDoc = await IotDeviceMapModel.findOne(filter).exec();
        }
        catch(err) {
            throw newInternalServerError("Database error", err);
        }
        
        return this.fromDatabaseConverter.visit(newIotDeviceMap(), iotDeviceMapDoc);
    }
}


function newIotDeviceMapRepository() {
    return new IotDeviceMapRepository();
}

module.exports.newIotDeviceMapRepository = newIotDeviceMapRepository;


