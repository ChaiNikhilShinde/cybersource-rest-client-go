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
    define(['ApiClient', 'model/Ptsv2paymentsClientReferenceInformationPartner'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'), require('./Ptsv2paymentsClientReferenceInformationPartner'));
  } else {
    // Browser globals (root is window)
    if (!root.CyberSource) {
      root.CyberSource = {};
    }
    root.CyberSource.TssV2TransactionsPost201ResponseEmbeddedPointOfSaleInformation = factory(root.CyberSource.ApiClient, root.CyberSource.Ptsv2paymentsClientReferenceInformationPartner);
  }
}(this, function(ApiClient, Ptsv2paymentsClientReferenceInformationPartner) {
  'use strict';




  /**
   * The TssV2TransactionsPost201ResponseEmbeddedPointOfSaleInformation model module.
   * @module model/TssV2TransactionsPost201ResponseEmbeddedPointOfSaleInformation
   * @version 0.0.1
   */

  /**
   * Constructs a new <code>TssV2TransactionsPost201ResponseEmbeddedPointOfSaleInformation</code>.
   * @alias module:model/TssV2TransactionsPost201ResponseEmbeddedPointOfSaleInformation
   * @class
   */
  var exports = function() {
    var _this = this;





  };

  /**
   * Constructs a <code>TssV2TransactionsPost201ResponseEmbeddedPointOfSaleInformation</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/TssV2TransactionsPost201ResponseEmbeddedPointOfSaleInformation} obj Optional instance to populate.
   * @return {module:model/TssV2TransactionsPost201ResponseEmbeddedPointOfSaleInformation} The populated <code>TssV2TransactionsPost201ResponseEmbeddedPointOfSaleInformation</code> instance.
   */
  exports.constructFromObject = function(data, obj) {
    if (data) {
      obj = obj || new exports();

      if (data.hasOwnProperty('terminalId')) {
        obj['terminalId'] = ApiClient.convertToType(data['terminalId'], 'String');
      }
      if (data.hasOwnProperty('terminalSerialNumber')) {
        obj['terminalSerialNumber'] = ApiClient.convertToType(data['terminalSerialNumber'], 'String');
      }
      if (data.hasOwnProperty('deviceId')) {
        obj['deviceId'] = ApiClient.convertToType(data['deviceId'], 'String');
      }
      if (data.hasOwnProperty('partner')) {
        obj['partner'] = Ptsv2paymentsClientReferenceInformationPartner.constructFromObject(data['partner']);
      }
    }
    return obj;
  }

  /**
   * Identifier for the terminal at your retail location. You can define this value yourself, but consult the processor for requirements.  #### FDC Nashville Global To have your account configured to support this field, contact CyberSource Customer Support. This value must be a value that FDC Nashville Global issued to you.  For details, see the `terminal_id` field description in [Card-Present Processing Using the SCMP API.](https://apps.cybersource.com/library/documentation/dev_guides/Retail_SCMP_API/html/wwhelp/wwhimpl/js/html/wwhelp.htm)  **For Payouts**: This field is applicable for CtV. 
   * @member {String} terminalId
   */
  exports.prototype['terminalId'] = undefined;
  /**
   * Terminal serial number assigned by the hardware manufacturer. This value is provided by the client software that is installed on the POS terminal.  CyberSource does not forward this value to the processor. Instead, the value is forwarded to the CyberSource reporting functionality.  This field is supported only on American Express Direct, FDC Nashville Global, and SIX.  For details, see the `terminal_serial_number` field description in [Card-Present Processing Using the SCMP API.](https://apps.cybersource.com/library/documentation/dev_guides/Retail_SCMP_API/html/wwhelp/wwhimpl/js/html/wwhelp.htm) 
   * @member {String} terminalSerialNumber
   */
  exports.prototype['terminalSerialNumber'] = undefined;
  /**
   * The description for this field is not available.
   * @member {String} deviceId
   */
  exports.prototype['deviceId'] = undefined;
  /**
   * @member {module:model/Ptsv2paymentsClientReferenceInformationPartner} partner
   */
  exports.prototype['partner'] = undefined;



  return exports;
}));


