/**
 * CyberSource Merged Spec
 * All CyberSource API specs merged together. These are available at https://developer.cybersource.com/api/reference/api-reference.html
 *
 * OpenAPI spec version: 0.0.1
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 *
 * Swagger Codegen version: 2.3.0
 *
 * Do not edit the class manually.
 *
 */

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD. Register as an anonymous module.
    define(['ApiClient', 'model/PtsV2PaymentsPost201ResponseClientReferenceInformation', 'model/PtsV2PaymentsReversalsPost201ResponseLinks', 'model/PtsV2PayoutsPost201ResponseErrorInformation', 'model/PtsV2PayoutsPost201ResponseMerchantInformation', 'model/PtsV2PayoutsPost201ResponseOrderInformation', 'model/PtsV2PayoutsPost201ResponseProcessorInformation', 'model/PtsV2PayoutsPost201ResponseRecipientInformation'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'), require('./PtsV2PaymentsPost201ResponseClientReferenceInformation'), require('./PtsV2PaymentsReversalsPost201ResponseLinks'), require('./PtsV2PayoutsPost201ResponseErrorInformation'), require('./PtsV2PayoutsPost201ResponseMerchantInformation'), require('./PtsV2PayoutsPost201ResponseOrderInformation'), require('./PtsV2PayoutsPost201ResponseProcessorInformation'), require('./PtsV2PayoutsPost201ResponseRecipientInformation'));
  } else {
    // Browser globals (root is window)
    if (!root.CyberSource) {
      root.CyberSource = {};
    }
    root.CyberSource.PtsV2PayoutsPost201Response = factory(root.CyberSource.ApiClient, root.CyberSource.PtsV2PaymentsPost201ResponseClientReferenceInformation, root.CyberSource.PtsV2PaymentsReversalsPost201ResponseLinks, root.CyberSource.PtsV2PayoutsPost201ResponseErrorInformation, root.CyberSource.PtsV2PayoutsPost201ResponseMerchantInformation, root.CyberSource.PtsV2PayoutsPost201ResponseOrderInformation, root.CyberSource.PtsV2PayoutsPost201ResponseProcessorInformation, root.CyberSource.PtsV2PayoutsPost201ResponseRecipientInformation);
  }
}(this, function(ApiClient, PtsV2PaymentsPost201ResponseClientReferenceInformation, PtsV2PaymentsReversalsPost201ResponseLinks, PtsV2PayoutsPost201ResponseErrorInformation, PtsV2PayoutsPost201ResponseMerchantInformation, PtsV2PayoutsPost201ResponseOrderInformation, PtsV2PayoutsPost201ResponseProcessorInformation, PtsV2PayoutsPost201ResponseRecipientInformation) {
  'use strict';




  /**
   * The PtsV2PayoutsPost201Response model module.
   * @module model/PtsV2PayoutsPost201Response
   * @version 0.0.1
   */

  /**
   * Constructs a new <code>PtsV2PayoutsPost201Response</code>.
   * @alias module:model/PtsV2PayoutsPost201Response
   * @class
   */
  var exports = function() {
    var _this = this;












  };

  /**
   * Constructs a <code>PtsV2PayoutsPost201Response</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/PtsV2PayoutsPost201Response} obj Optional instance to populate.
   * @return {module:model/PtsV2PayoutsPost201Response} The populated <code>PtsV2PayoutsPost201Response</code> instance.
   */
  exports.constructFromObject = function(data, obj) {
    if (data) {
      obj = obj || new exports();

      if (data.hasOwnProperty('_links')) {
        obj['_links'] = PtsV2PaymentsReversalsPost201ResponseLinks.constructFromObject(data['_links']);
      }
      if (data.hasOwnProperty('id')) {
        obj['id'] = ApiClient.convertToType(data['id'], 'String');
      }
      if (data.hasOwnProperty('submitTimeUtc')) {
        obj['submitTimeUtc'] = ApiClient.convertToType(data['submitTimeUtc'], 'String');
      }
      if (data.hasOwnProperty('status')) {
        obj['status'] = ApiClient.convertToType(data['status'], 'String');
      }
      if (data.hasOwnProperty('reconciliationId')) {
        obj['reconciliationId'] = ApiClient.convertToType(data['reconciliationId'], 'String');
      }
      if (data.hasOwnProperty('errorInformation')) {
        obj['errorInformation'] = PtsV2PayoutsPost201ResponseErrorInformation.constructFromObject(data['errorInformation']);
      }
      if (data.hasOwnProperty('clientReferenceInformation')) {
        obj['clientReferenceInformation'] = PtsV2PaymentsPost201ResponseClientReferenceInformation.constructFromObject(data['clientReferenceInformation']);
      }
      if (data.hasOwnProperty('merchantInformation')) {
        obj['merchantInformation'] = PtsV2PayoutsPost201ResponseMerchantInformation.constructFromObject(data['merchantInformation']);
      }
      if (data.hasOwnProperty('orderInformation')) {
        obj['orderInformation'] = PtsV2PayoutsPost201ResponseOrderInformation.constructFromObject(data['orderInformation']);
      }
      if (data.hasOwnProperty('processorInformation')) {
        obj['processorInformation'] = PtsV2PayoutsPost201ResponseProcessorInformation.constructFromObject(data['processorInformation']);
      }
      if (data.hasOwnProperty('recipientInformation')) {
        obj['recipientInformation'] = PtsV2PayoutsPost201ResponseRecipientInformation.constructFromObject(data['recipientInformation']);
      }
    }
    return obj;
  }

  /**
   * @member {module:model/PtsV2PaymentsReversalsPost201ResponseLinks} _links
   */
  exports.prototype['_links'] = undefined;
  /**
   * An unique identification number assigned by CyberSource to identify the submitted request. It is also appended to the endpoint of the resource.
   * @member {String} id
   */
  exports.prototype['id'] = undefined;
  /**
   * Time of request in UTC. `Format: YYYY-MM-DDThh:mm:ssZ`  Example 2016-08-11T22:47:57Z equals August 11, 2016, at 22:47:57 (10:47:57 p.m.). The T separates the date and the time. The Z indicates UTC. 
   * @member {String} submitTimeUtc
   */
  exports.prototype['submitTimeUtc'] = undefined;
  /**
   * The status of the submitted transaction.  Possible values:  - ACCEPTED  - DECLINED  - INVALID_REQUEST 
   * @member {module:model/PtsV2PayoutsPost201Response.StatusEnum} status
   */
  exports.prototype['status'] = undefined;
  /**
   * Cybersource or merchant generated transaction reference number. This is sent to the processor and is echoed back in the response to the merchant. This is This value is used for reconciliation purposes. 
   * @member {String} reconciliationId
   */
  exports.prototype['reconciliationId'] = undefined;
  /**
   * @member {module:model/PtsV2PayoutsPost201ResponseErrorInformation} errorInformation
   */
  exports.prototype['errorInformation'] = undefined;
  /**
   * @member {module:model/PtsV2PaymentsPost201ResponseClientReferenceInformation} clientReferenceInformation
   */
  exports.prototype['clientReferenceInformation'] = undefined;
  /**
   * @member {module:model/PtsV2PayoutsPost201ResponseMerchantInformation} merchantInformation
   */
  exports.prototype['merchantInformation'] = undefined;
  /**
   * @member {module:model/PtsV2PayoutsPost201ResponseOrderInformation} orderInformation
   */
  exports.prototype['orderInformation'] = undefined;
  /**
   * @member {module:model/PtsV2PayoutsPost201ResponseProcessorInformation} processorInformation
   */
  exports.prototype['processorInformation'] = undefined;
  /**
   * @member {module:model/PtsV2PayoutsPost201ResponseRecipientInformation} recipientInformation
   */
  exports.prototype['recipientInformation'] = undefined;


  /**
   * Allowed values for the <code>status</code> property.
   * @enum {String}
   * @readonly
   */
  exports.StatusEnum = {
    /**
     * value: "ACCEPTED"
     * @const
     */
    "ACCEPTED": "ACCEPTED",
    /**
     * value: "DECLINED"
     * @const
     */
    "DECLINED": "DECLINED",
    /**
     * value: "INVALID_REQUEST"
     * @const
     */
    "INVALID_REQUEST": "INVALID_REQUEST"  };


  return exports;
}));


