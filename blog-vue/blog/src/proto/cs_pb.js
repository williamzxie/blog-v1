// source: cs.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!
/* eslint-disable */
// @ts-nocheck

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

goog.exportSymbol('proto.proto.CsId', null, global);
goog.exportSymbol('proto.proto.RequestPkg', null, global);
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.proto.RequestPkg = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.proto.RequestPkg, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.proto.RequestPkg.displayName = 'proto.proto.RequestPkg';
}



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.proto.RequestPkg.prototype.toObject = function(opt_includeInstance) {
  return proto.proto.RequestPkg.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.proto.RequestPkg} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.proto.RequestPkg.toObject = function(includeInstance, msg) {
  var f, obj = {
    cmdid: jspb.Message.getFieldWithDefault(msg, 1, 0),
    currentpage: jspb.Message.getFieldWithDefault(msg, 2, 0),
    articleid: jspb.Message.getFieldWithDefault(msg, 3, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.proto.RequestPkg}
 */
proto.proto.RequestPkg.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.proto.RequestPkg;
  return proto.proto.RequestPkg.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.proto.RequestPkg} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.proto.RequestPkg}
 */
proto.proto.RequestPkg.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {!proto.proto.CsId} */ (reader.readEnum());
      msg.setCmdid(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readUint32());
      msg.setCurrentpage(value);
      break;
    case 3:
      var value = /** @type {number} */ (reader.readUint32());
      msg.setArticleid(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.proto.RequestPkg.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.proto.RequestPkg.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.proto.RequestPkg} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.proto.RequestPkg.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getCmdid();
  if (f !== 0.0) {
    writer.writeEnum(
      1,
      f
    );
  }
  f = message.getCurrentpage();
  if (f !== 0) {
    writer.writeUint32(
      2,
      f
    );
  }
  f = message.getArticleid();
  if (f !== 0) {
    writer.writeUint32(
      3,
      f
    );
  }
};


/**
 * optional CsId cmdId = 1;
 * @return {!proto.proto.CsId}
 */
proto.proto.RequestPkg.prototype.getCmdid = function() {
  return /** @type {!proto.proto.CsId} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {!proto.proto.CsId} value
 * @return {!proto.proto.RequestPkg} returns this
 */
proto.proto.RequestPkg.prototype.setCmdid = function(value) {
  return jspb.Message.setProto3EnumField(this, 1, value);
};


/**
 * optional uint32 currentPage = 2;
 * @return {number}
 */
proto.proto.RequestPkg.prototype.getCurrentpage = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/**
 * @param {number} value
 * @return {!proto.proto.RequestPkg} returns this
 */
proto.proto.RequestPkg.prototype.setCurrentpage = function(value) {
  return jspb.Message.setProto3IntField(this, 2, value);
};


/**
 * optional uint32 articleId = 3;
 * @return {number}
 */
proto.proto.RequestPkg.prototype.getArticleid = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/**
 * @param {number} value
 * @return {!proto.proto.RequestPkg} returns this
 */
proto.proto.RequestPkg.prototype.setArticleid = function(value) {
  return jspb.Message.setProto3IntField(this, 3, value);
};


/**
 * @enum {number}
 */
proto.proto.CsId = {
  CSBEGININDEX: 0,
  CSGETARTICLES: 1,
  CSGETARTICLEBYID: 2,
  CSGETBLOGHOMEINFO: 3
};

goog.object.extend(exports, proto.proto);