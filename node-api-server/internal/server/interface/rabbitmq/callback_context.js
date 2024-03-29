const { newIotEventNewCallback } = require("./iot_event_new_callback");
const { newCameraEventNewCallback } = require("./camera_event_new_callback");
const { newEventVerifiedCallback } = require("./event_verified_callback");
const { newEventNotificationCallback } = require("./event_notification_callback");

class CallbackContext {

    getEventCallback(exchangeName, queueName) {
        if (exchangeName == "amq.topic" && queueName == "iot_event_new") {
            return newIotEventNewCallback();
        }
        if (exchangeName == "event_processing" && queueName == "event_verified") {
            return newEventVerifiedCallback();
        }
        if (exchangeName == "stream_processsing" && queueName == "camera_event_new") {
            return newCameraEventNewCallback();
        }
        if (exchangeName == "event_notification") {
            return newEventNotificationCallback();
        }
        return null;
    }
}

function newCallbackContext() {
    return new CallbackContext();
}

module.exports.newCallbackContext = newCallbackContext;
