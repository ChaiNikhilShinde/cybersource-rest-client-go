/**
 * CyberSource Flex API
 * Simple PAN tokenization service
 *
 * OpenAPI spec version: 0.0.1
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 *
 * Swagger Codegen version: 2.2.3
 *
 * Do not edit the class manually.
 *
 */

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD. Register as an anonymous module.
    define(['ApiClient', 'model/ReportingV3NotificationofChangesGet400Response', 'model/ReportingV3ReportsGet200Response', 'model/ReportingV3ReportsIdGet200Response', 'model/RequestBody1'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'), require('../model/ReportingV3NotificationofChangesGet400Response'), require('../model/ReportingV3ReportsGet200Response'), require('../model/ReportingV3ReportsIdGet200Response'), require('../model/RequestBody1'));
  } else {
    // Browser globals (root is window)
    if (!root.CyberSource) {
      root.CyberSource = {};
    }
    root.CyberSource.ReportsApi = factory(root.CyberSource.ApiClient, root.CyberSource.ReportingV3NotificationofChangesGet400Response, root.CyberSource.ReportingV3ReportsGet200Response, root.CyberSource.ReportingV3ReportsIdGet200Response, root.CyberSource.RequestBody1);
  }
}(this, function(ApiClient, ReportingV3NotificationofChangesGet400Response, ReportingV3ReportsGet200Response, ReportingV3ReportsIdGet200Response, RequestBody1) {
  'use strict';

  /**
   * Reports service.
   * @module api/ReportsApi
   * @version 0.0.1
   */

  /**
   * Constructs a new ReportsApi. 
   * @alias module:api/ReportsApi
   * @class
   * @param {module:ApiClient} apiClient Optional API client implementation to use,
   * default to {@link module:ApiClient#instance} if unspecified.
   */
  var exports = function(configObject, apiClient = undefined) {
    this.apiClient = apiClient || ApiClient.instance;

    this.apiClient.setConfiguration(configObject);


    /**
     * Callback function to receive the result of the createReport operation.
     * @callback module:api/ReportsApi~createReportCallback
     * @param {String} error Error message, if any.
     * @param data This operation does not return a value.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Create Adhoc Report
     * Create one time report
     * @param {module:model/RequestBody1} requestBody Report subscription request payload
     * @param {module:api/ReportsApi~createReportCallback} callback The callback function, accepting three arguments: error, data, response
     */
    this.createReport = function(requestBody, callback) {
      var postBody = requestBody;

      // verify the required parameter 'requestBody' is set
      if (requestBody === undefined || requestBody === null) {
        throw new Error("Missing the required parameter 'requestBody' when calling createReport");
      }


      var pathParams = {
      };
      var queryParams = {
      };
      var headerParams = {
      };
      var formParams = {
      };

      var authNames = [];
      var contentTypes = ['application/json'];
      var accepts = ['application/hal+json'];
      var returnType = null;

      return this.apiClient.callApi(
        '/reporting/v3/reports', 'POST',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the getReportByReportId operation.
     * @callback module:api/ReportsApi~getReportByReportIdCallback
     * @param {String} error Error message, if any.
     * @param {module:model/ReportingV3ReportsIdGet200Response} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Get Report based on reportId
     * ReportId is mandatory input
     * @param {String} reportId Valid Report Id
     * @param {Object} opts Optional parameters
     * @param {String} opts.organizationId Valid Cybersource Organization Id
     * @param {module:api/ReportsApi~getReportByReportIdCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/ReportingV3ReportsIdGet200Response}
     */
    this.getReportByReportId = function(reportId, opts, callback) {
      opts = opts || {};
      var postBody = null;

      // verify the required parameter 'reportId' is set
      if (reportId === undefined || reportId === null) {
        throw new Error("Missing the required parameter 'reportId' when calling getReportByReportId");
      }


      var pathParams = {
        'reportId': reportId
      };
      var queryParams = {
        'organizationId': opts['organizationId']
      };
      var headerParams = {
      };
      var formParams = {
      };

      var authNames = [];
      var contentTypes = ['application/json;charset=utf-8'];
      var accepts = ['application/hal+json', 'application/xml'];
      var returnType = ReportingV3ReportsIdGet200Response;

      return this.apiClient.callApi(
        '/reporting/v3/reports/{reportId}', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the searchReports operation.
     * @callback module:api/ReportsApi~searchReportsCallback
     * @param {String} error Error message, if any.
     * @param {module:model/ReportingV3ReportsGet200Response} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Retrieve available reports
     * Retrieve list of available reports
     * @param {Date} startTime Valid report Start Time in **ISO 8601 format** Please refer the following link to know more about ISO 8601 format. - https://xml2rfc.tools.ietf.org/public/rfc/html/rfc3339.html#anchor14   **Example date format:**   - yyyy-MM-dd&#39;T&#39;HH:mm:ssXXX 
     * @param {Date} endTime Valid report End Time in **ISO 8601 format** Please refer the following link to know more about ISO 8601 format. - https://xml2rfc.tools.ietf.org/public/rfc/html/rfc3339.html#anchor14   **Example date format:**   - yyyy-MM-dd&#39;T&#39;HH:mm:ssXXX 
     * @param {module:model/String} timeQueryType Specify time you woud like to search
     * @param {Object} opts Optional parameters
     * @param {String} opts.organizationId Valid Cybersource Organization Id
     * @param {module:model/String} opts.reportMimeType Valid Report Format
     * @param {module:model/String} opts.reportFrequency Valid Report Frequency
     * @param {String} opts.reportName Valid Report Name
     * @param {Number} opts.reportDefinitionId Valid Report Definition Id
     * @param {module:model/String} opts.reportStatus Valid Report Status
     * @param {module:api/ReportsApi~searchReportsCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/ReportingV3ReportsGet200Response}
     */
    this.searchReports = function(startTime, endTime, timeQueryType, opts, callback) {
      opts = opts || {};
      var postBody = null;

      // verify the required parameter 'startTime' is set
      if (startTime === undefined || startTime === null) {
        throw new Error("Missing the required parameter 'startTime' when calling searchReports");
      }

      // verify the required parameter 'endTime' is set
      if (endTime === undefined || endTime === null) {
        throw new Error("Missing the required parameter 'endTime' when calling searchReports");
      }

      // verify the required parameter 'timeQueryType' is set
      if (timeQueryType === undefined || timeQueryType === null) {
        throw new Error("Missing the required parameter 'timeQueryType' when calling searchReports");
      }


      var pathParams = {
      };
      var queryParams = {
        'organizationId': opts['organizationId'],
        'startTime': startTime,
        'endTime': endTime,
        'timeQueryType': timeQueryType,
        'reportMimeType': opts['reportMimeType'],
        'reportFrequency': opts['reportFrequency'],
        'reportName': opts['reportName'],
        'reportDefinitionId': opts['reportDefinitionId'],
        'reportStatus': opts['reportStatus']
      };
      var headerParams = {
      };
      var formParams = {
      };

      var authNames = [];
      var contentTypes = ['application/json;charset=utf-8'];
      var accepts = ['application/hal+json'];
      var returnType = ReportingV3ReportsGet200Response;

      return this.apiClient.callApi(
        '/reporting/v3/reports', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }
  };

  return exports;
}));
