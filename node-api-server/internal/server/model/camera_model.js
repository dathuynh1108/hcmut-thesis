const mongoose = require('mongoose');
const Schema = mongoose.Schema;
const { newStreamConnectionService } = require("../service/stream_connection_service");
const { newFromDatabaseConverter } = require("../data_converter/from_database_converter");
const { newToDatabaseConverter } = require("../data_converter/to_database_converter");
const { newCamera } = require("../entity/camera");
const { newErrorHandler } = require("../error/error_handler");

const streamConnectionService = newStreamConnectionService();
const fromDatabaseConverter = newFromDatabaseConverter();
const toDatabaseConverter = newToDatabaseConverter();
const errorHandler = newErrorHandler();


const cameraSchema = new Schema (
    {
        camera_name: {
            type: String,
            unique: true,
            required: true
        },
        status: {
            type: String,
            enum: ['used', 'free'],
            default: 'free'
        },

        rtsp_stream_url: String,
        sfu_rtsp_stream_url: String,
        username: String, // rtsp stream username
        password: String, // rtsp stream password

        offset_x_begin: Number,
        offset_x_end: Number,
        offset_y_begin: Number,
        offset_y_end: Number,
        is_set_line: {
            type: Boolean,
            default: false
        },

        camera_type: {
            type: Schema.Types.ObjectID,
            ref: 'CameraType'
        },

        event_type: {
            type: Schema.Types.ObjectID,
            ref: 'EventType'
        },
        connect_to_controller: {
            type: Boolean,
            default: false
        },
        connect_to_ai: {
            type: Boolean, 
            default: false
        },



        hostname: String, // camera server host name 
        port: Number,
        stream_resolution: {
            type: [Number],
            default: []
        }, // [width, height]
        iot_event_zone_coords: {
            type: [Number],
            default: []
        }, // from_x, from_y, to_x, to_y
        camera_event_zone_coords: {
            type: [Number],
            default: []
        },

        line_crossing_vector: {
            type: [Number],
            default: []
        },

        need_update_connection: {
            type: Boolean,
            default: false
        }
    }
)

async function deleteCameraRelation(schema) {
    const doc = await schema.model.findOne(schema.getFilter());
    if (doc) {
        await mongoose.model("CameraMap").findOneAndUpdate({ connect_camera: doc._id }, { connect_camera: null });
        await mongoose.model("Event").deleteMany({ camera: doc._id });
    }
}




async function createStreamConnection(schema) {
    let state = newCamera();
    try {
        const doc = await mongoose.model("Camera").findOne({ _id: schema._id }).populate("event_type");

        const camera = fromDatabaseConverter.visit(newCamera(), doc);
        await streamConnectionService.handleCreateStream(camera, state);
    }
    catch(err) {
        errorHandler.execute(err);
        // errorHandler.execute(new Error("Failed when creating stream connection to other servers"));
    }
    finally {
        state = toDatabaseConverter.visit(state);
        await mongoose.model('Camera').findOneAndUpdate({ _id: state._id }, state);
    }

}


async function deleteStreamConnection(doc) {
    try {
        if (doc) {
            const camera = fromDatabaseConverter.visit(newCamera(), doc);
            await streamConnectionService.handleDeleteStream(camera);
        }
    }
    catch(err) {
        errorHandler.execute(err);
        // errorHandler.execute(new Error("Failed when deleting stream connection to other servers"));
    }

}


async function updateStreamConnection(updateDoc, updatedDoc) {
    if (!updatedDoc || updateDoc.connect_to_controller !== undefined || updateDoc.connect_to_ai !== undefined) return;
    // if (updatedDoc.need_update_connection === false) return;
    let state = newCamera();
    try {
        updatedDoc = await mongoose.model("Camera").findOne({ _id: updatedDoc._id }).populate("event_type");
        const updatedCamera = fromDatabaseConverter.visit(newCamera(), updatedDoc);

        const needDeleteInController = updatedDoc.need_update_connection;
        const needCreateInController = updatedDoc.need_update_connection;

        await streamConnectionService.handleDeleteStream(updatedCamera, state, needDeleteInController);
        await streamConnectionService.handleCreateStream(updatedCamera, state, needCreateInController);
    }
    catch(err) {
        errorHandler.execute(err);
        // errorHandler.execute(new Error("Failed when updating stream connection to other servers"));
    }
    finally {
        state = toDatabaseConverter.visit(state);
        
        await mongoose.model('Camera').findOneAndUpdate({ _id: state._id }, state);
    }
}


// async function updateStreamConnection(schema) {
//     try {
//         const doc = await schema.model.findOne(schema.getFilter());
//         const updateDoc = schema.getUpdate();
//         if (updateDoc.connect_to_controller = undefined && updateDoc.connect_to_ai == undefined && doc) {
//             const oldCamera = fromDatabaseConverter.visit(newCamera(), doc);
//             const updateCamera = fromDatabaseConverter.visit(newCamera(), updateDoc);
//             let state = await streamConnectionService.handleUpdateStream(oldCamera, updateCamera);
//             state = toDatabaseConverter.visit(state);
//             await mongoose.model('Camera').findOneAndUpdate({ _id: state._id }, state);
//         }
//     }
//     catch(err) {
//         errorHandler.execute(err);
//         errorHandler.execute(new Error("Failed when updating stream connection to other servers"));
//     }
// }



cameraSchema.post('save', async function(doc, next) {
    createStreamConnection(this);
})

cameraSchema.pre('findOneAndUpdate', async function(next) {
    const doc = await this.model.findOne(this.getFilter());
    const updateDoc = this.getUpdate();
    if (doc) {

        if ((updateDoc.rtsp_stream_url !== undefined && doc.rtsp_stream_url != updateDoc.rtsp_stream_url) || (updateDoc.username !== undefined && doc.username != updateDoc.username) || (doc.password != updateDoc.password && updateDoc.password !== undefined) || (updateDoc.status !== undefined && doc.status != updateDoc.status)) {
            this._update.need_update_connection = true;
        }
        else {
            this._update.need_update_connection = false;
        }
    }
    next();
})

cameraSchema.post('findOneAndUpdate', async function(doc, next) {
    const updateDoc = this.getUpdate().$set;
    const updatedDoc = doc;
    updateStreamConnection(updateDoc, updatedDoc);
    // console.log(doc);
    // console.log(this.getFilter());
    // console.log(this.getUpdate().$set);
})


cameraSchema.pre('findOneAndDelete', async function(next) {
    await deleteCameraRelation(this);
    next();
})

cameraSchema.pre('deleteMany', async function(next) {
    await deleteCameraRelation(this);
    next();
})


cameraSchema.post('findOneAndDelete', async function(doc, next) {
    deleteStreamConnection(doc);
})

cameraSchema.post('deleteMany', async function(doc, next) {
    deleteStreamConnection(doc);
})

// cameraSchema.pre("remove", async function(next) {
//     const self = this;
//     await mongoose.model("Event").deleteMany({ camera: self._id });
//     await mongoose.model("CameraMap").updateMany({ connect_camera: self._id }, { connect_camera: null });
//     next();
// })


module.exports = mongoose.model('Camera', cameraSchema);
