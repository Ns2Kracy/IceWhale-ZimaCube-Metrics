/**
 * This file was auto-generated by openapi-typescript.
 * Do not make direct changes to the file.
 */


export interface paths {
  "/": {
    /** Get the service metrics of ZimaCube */
    get: operations["getMetrics"];
    /** Add a ssh host to monitor */
    post: operations["postAddZimaCube"];
    /** Delete a ssh connection from monitoring */
    delete: operations["deleteZimaCube"];
  };
}

export type webhooks = Record<string, never>;

export interface components {
  schemas: {
    BaseResponse: {
      /** @description message returned by server side if there is any */
      message?: string;
    };
    ZimaCubeMetrics: {
      /** @description IP address of the server */
      ip?: string;
      /** @description Metrics of the server */
      metrics?: components["schemas"]["Metric"][];
    };
    Metric: {
      /**
       * @description Name of the service
       * @example zimaos
       */
      name?: string;
      /**
       * @description CPU usage of the server
       * @example 0.5%
       */
      cpu?: string;
      /**
       * @description Maximum CPU usage of the server in the last 1 hour
       * @example 5%
       */
      max_cpu?: string;
      /**
       * @description Average CPU usage of the server in the last 1 hour
       * @example 5%
       */
      avg_cpu?: string;
      /**
       * @description Memory usage of the server
       * @example 0.5%
       */
      mem?: string;
      /**
       * @description Maximum Memory usage of the server in the last 1 hour
       * @example 5%
       */
      max_mem?: string;
      /**
       * @description Average Memory usage of the server in the last 1 hour
       * @example 5%
       */
      avg_mem?: string;
      /**
       * @description Uptime of the service
       * @example 1 day 2 hours 3 minutes
       */
      uptime?: string;
    };
  };
  responses: {
    /** @description OK */
    ResponseOK: {
      content: {
        "application/json": components["schemas"]["BaseResponse"];
      };
    };
    /** @description Bad Request */
    ResponeBadRequest: {
      content: {
        "application/json": components["schemas"]["BaseResponse"];
      };
    };
    /** @description Internal Server Error */
    ResponseInternalServerError: {
      content: {
        "application/json": components["schemas"]["BaseResponse"];
      };
    };
    /** @description Metrics of ZimaCube */
    ResponseZimaCubeMetricsOK: {
      content: {
        "application/json": components["schemas"]["BaseResponse"] & {
          data?: components["schemas"]["Metric"][];
        };
      };
    };
  };
  parameters: never;
  requestBodies: {
    AddZimaCube?: {
      content: {
        "application/json": {
          /**
           * @description IP address of the server
           * @example 10.0.0.85
           */
          ip?: string;
          /**
           * @description Port of the server
           * @example 22
           */
          port?: number;
          /**
           * @description Username of the server
           * @example root
           */
          username?: string;
          /**
           * @description Password of the server
           * @example password
           */
          password?: string;
        };
      };
    };
  };
  headers: never;
  pathItems: never;
}

export type $defs = Record<string, never>;

export type external = Record<string, never>;

export interface operations {

  /** Get the service metrics of ZimaCube */
  getMetrics: {
    responses: {
      200: components["responses"]["ResponseZimaCubeMetricsOK"];
      400: components["responses"]["ResponeBadRequest"];
      500: components["responses"]["ResponseInternalServerError"];
    };
  };
  /** Add a ssh host to monitor */
  postAddZimaCube: {
    requestBody: components["requestBodies"]["AddZimaCube"];
    responses: {
      200: components["responses"]["ResponseOK"];
      400: components["responses"]["ResponeBadRequest"];
      500: components["responses"]["ResponseInternalServerError"];
    };
  };
  /** Delete a ssh connection from monitoring */
  deleteZimaCube: {
    parameters: {
      query: {
        /** @description IP address of the server */
        ip: string;
      };
    };
    responses: {
      200: components["responses"]["ResponseOK"];
      400: components["responses"]["ResponeBadRequest"];
      500: components["responses"]["ResponseInternalServerError"];
    };
  };
}
