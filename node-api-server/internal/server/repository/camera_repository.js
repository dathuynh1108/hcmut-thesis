

const CameraModel = require("../model/camera_model");
const { newCamera } = require("../entity/camera");
const { newFromDatabaseConverter } = require("../data_converter/from_database_converter");
const { newToDatabaseConverter } = require("../data_converter/to_database_converter");

const { newInternalServerError } = require("../entity/error/internal_server_error");


class CameraRepository {
    constructor() {
        this.fromDatabaseConverter = newFromDatabaseConverter();
        this.toDatabaseConverter = newToDatabaseConverter();
    }



    async getAll() {
        let cameraDocs;
        try {
            cameraDocs = await CameraModel.find({});
        }
        catch(err) {
            throw newInternalServerError("Database error", err);
        }
        return cameraDocs.map(cameraDoc => {
            return this.fromDatabaseConverter.visit(newCamera(), cameraDoc);
        })
    }
    
    async create(cameraEntity) {
        const cameraDoc = this.toDatabaseConverter.visit(cameraEntity);
        let newCameraDoc;
        try {
            newCameraDoc = await CameraModel.create(cameraDoc);
        }
        catch(err) {
            throw newInternalServerError("Database error", err);
        }
        return this.fromDatabaseConverter.visit(newCamera(), newCameraDoc);
    }
    
    async findById(cameraId) {
        let cameraDoc;
        cameraId = cameraId.getValue();
        try {
            cameraDoc = await CameraModel.findById(cameraId).exec();
        }
        catch(err) {
            throw newInternalServerError("Database error", err);
        }
        
        return this.fromDatabaseConverter.visit(newCamera(), cameraDoc);
    }
    
    
    async findByName(cameraName) {
        let cameraDoc;
    
        try {
            cameraDoc = await CameraModel.findOne({
                camera_name: cameraName
            }).exec();
        }
        catch(err) {
            throw newInternalServerError("Database error", err);
        }
        
        return this.fromDatabaseConverter.visit(newCamera(), cameraDoc);
    }
    
    
    async findByIdAndUpdate(cameraId, cameraEntity) {
        const cameraDoc = this.toDatabaseConverter.visit(cameraEntity);
        const filter = {
            _id: cameraId.getValue()
        }
        let newCameraDoc;
        try {
            newCameraDoc = await CameraModel.findOneAndUpdate(filter, cameraDoc, { new: true }); // set new to true to return new document after update
        }
        catch(err) {
            throw newInternalServerError("Database error", err);
        }
        return this.fromDatabaseConverter.visit(newCamera(), newCameraDoc);
    }

    async findByIdAndDelete(cameraId) {
        const filter = {
            _id: cameraId.getValue()
        }
        let deleteCameraDoc;
        try {
            deleteCameraDoc = await CameraModel.findOneAndDelete(filter); // set new to true to return new document after update
        }
        catch(err) {
            throw newInternalServerError("Database error", err);
        }
        return this.fromDatabaseConverter.visit(newCamera(), deleteCameraDoc);
    }
}



function newCameraRepository() {
    return new CameraRepository();
}

module.exports.newCameraRepository = newCameraRepository;